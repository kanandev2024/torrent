package torrent

import (
	"math/rand"
	"sync"

	"github.com/anacrolix/missinggo/bitmap"
	"github.com/bradfitz/iter"

	pp "github.com/anacrolix/torrent/peer_protocol"
)

// Piece priority describes the importance of obtaining a particular piece.

type piecePriority byte

func (me *piecePriority) Raise(maybe piecePriority) {
	if maybe > *me {
		*me = maybe
	}
}

const (
	PiecePriorityNone      piecePriority = iota // Not wanted.
	PiecePriorityNormal                         // Wanted.
	PiecePriorityReadahead                      // May be required soon.
	PiecePriorityNext                           // Succeeds a piece where a read occurred.
	PiecePriorityNow                            // A read occurred in this piece.
)

type piece struct {
	// The completed piece SHA1 hash, from the metainfo "pieces" field.
	Hash pieceSum
	// Chunks we've written to since the last check. The chunk offset and
	// length can be determined by the request chunkSize in use.
	DirtyChunks      bitmap.Bitmap
	Hashing          bool
	QueuedForHash    bool
	EverHashed       bool
	PublicPieceState PieceState
	priority         piecePriority

	pendingWritesMutex sync.Mutex
	pendingWrites      int
	noPendingWrites    sync.Cond
}

func (p *piece) pendingChunkIndex(chunkIndex int) bool {
	return !p.DirtyChunks.Contains(chunkIndex)
}

func (p *piece) pendingChunk(cs chunkSpec, chunkSize pp.Integer) bool {
	return p.pendingChunkIndex(chunkIndex(cs, chunkSize))
}

func (p *piece) hasDirtyChunks() bool {
	return p.DirtyChunks.Len() != 0
}

func (p *piece) numDirtyChunks() (ret int) {
	return p.DirtyChunks.Len()
}

func (p *piece) unpendChunkIndex(i int) {
	p.DirtyChunks.Add(i)
}

func (p *piece) pendChunkIndex(i int) {
	p.DirtyChunks.Remove(i)
}

func chunkIndexSpec(index int, pieceLength, chunkSize pp.Integer) chunkSpec {
	ret := chunkSpec{pp.Integer(index) * chunkSize, chunkSize}
	if ret.Begin+ret.Length > pieceLength {
		ret.Length = pieceLength - ret.Begin
	}
	return ret
}

func (p *piece) shuffledPendingChunkSpecs(t *torrent, piece int) (css []chunkSpec) {
	// defer func() {
	// 	log.Println(piece, css)
	// }()
	numPending := t.pieceNumPendingChunks(piece)
	if numPending == 0 {
		return
	}
	css = make([]chunkSpec, 0, numPending)
	for ci := range iter.N(t.pieceNumChunks(piece)) {
		if !p.DirtyChunks.Contains(ci) {
			css = append(css, t.chunkIndexSpec(ci, piece))
		}
	}
	if len(css) <= 1 {
		return
	}
	for i := range css {
		j := rand.Intn(i + 1)
		css[i], css[j] = css[j], css[i]
	}
	return
}

func (p *piece) incrementPendingWrites() {
	p.pendingWritesMutex.Lock()
	p.pendingWrites++
	p.pendingWritesMutex.Unlock()
}

func (p *piece) decrementPendingWrites() {
	p.pendingWritesMutex.Lock()
	if p.pendingWrites == 0 {
		panic("assertion")
	}
	p.pendingWrites--
	if p.pendingWrites == 0 {
		p.noPendingWrites.Broadcast()
	}
	p.pendingWritesMutex.Unlock()
}

func (p *piece) waitNoPendingWrites() {
	p.pendingWritesMutex.Lock()
	for p.pendingWrites != 0 {
		p.noPendingWrites.Wait()
	}
	p.pendingWritesMutex.Unlock()
}
