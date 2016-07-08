package storage

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/anacrolix/missinggo"
	"github.com/anacrolix/missinggo/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anacrolix/torrent/metainfo"
)

func TestShortFile(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	s := NewFile(td)
	info := &metainfo.InfoEx{
		Info: metainfo.Info{
			Name:        "a",
			Length:      2,
			PieceLength: missinggo.MiB,
		},
	}
	ts, err := s.OpenTorrent(info)
	assert.NoError(t, err)
	f, err := os.Create(filepath.Join(td, "a"))
	err = f.Truncate(1)
	f.Close()
	var buf bytes.Buffer
	p := info.Piece(0)
	n, err := io.Copy(&buf, io.NewSectionReader(ts.Piece(p), 0, p.Length()))
	assert.EqualValues(t, 1, n)
	assert.Equal(t, io.ErrUnexpectedEOF, err)
}

// Two different torrents opened from the same storage. Closing one should not
// break the piece completion on the other.
func testIssue95(t *testing.T, c Client) {
	i1 := &metainfo.InfoEx{
		Bytes: []byte("a"),
		Info: metainfo.Info{
			Files:  []metainfo.FileInfo{{Path: []string{"a"}}},
			Pieces: make([]byte, 20),
		},
	}
	t1, err := c.OpenTorrent(i1)
	require.NoError(t, err)
	i2 := &metainfo.InfoEx{
		Bytes: []byte("b"),
		Info: metainfo.Info{
			Files:  []metainfo.FileInfo{{Path: []string{"a"}}},
			Pieces: make([]byte, 20),
		},
	}
	t2, err := c.OpenTorrent(i2)
	require.NoError(t, err)
	t2p := t2.Piece(i2.Piece(0))
	assert.NoError(t, t1.Close())
	assert.NotPanics(t, func() { t2p.GetIsComplete() })
}

func TestIssue95File(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	testIssue95(t, NewFile(td))
}

func TestIssue95MMap(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	testIssue95(t, NewMMap(td))
}

func TestIssue95ResourcePieces(t *testing.T) {
	td, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(td)
	testIssue95(t, NewResourcePieces(resource.OSFileProvider{}))
}
