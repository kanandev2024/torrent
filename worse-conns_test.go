package torrent

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func TestWorseConnLastHelpful(t *testing.T) {
	c := qt.New(t)
	c.Check((&worseConnInput{}).Less(&worseConnInput{LastHelpful: time.Now()}), qt.IsTrue)
	c.Check((&worseConnInput{}).Less(&worseConnInput{CompletedHandshake: time.Now()}), qt.IsTrue)
	c.Check((&worseConnInput{LastHelpful: time.Now()}).Less(&worseConnInput{CompletedHandshake: time.Now()}), qt.IsFalse)
	c.Check((&worseConnInput{
		LastHelpful: time.Now(),
	}).Less(&worseConnInput{
		LastHelpful:        time.Now(),
		CompletedHandshake: time.Now(),
	}), qt.IsTrue)
	now := time.Now()
	c.Check((&worseConnInput{
		LastHelpful: now,
	}).Less(&worseConnInput{
		LastHelpful:        now.Add(-time.Nanosecond),
		CompletedHandshake: now,
	}), qt.IsFalse)
	readyPeerPriority := func() (peerPriority, error) {
		return 42, nil
	}
	c.Check((&worseConnInput{
		GetPeerPriority: readyPeerPriority,
	}).Less(&worseConnInput{
		GetPeerPriority: readyPeerPriority,
		Pointer:         1,
	}), qt.IsTrue)
	c.Check((&worseConnInput{
		GetPeerPriority: readyPeerPriority,
		Pointer:         2,
	}).Less(&worseConnInput{
		GetPeerPriority: readyPeerPriority,
		Pointer:         1,
	}), qt.IsFalse)
}
