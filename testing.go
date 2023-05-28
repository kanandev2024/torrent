package torrent

import (
	"testing"
	"time"

	"github.com/anacrolix/log"

	pp "github.com/anacrolix/torrent/peer_protocol"
)

func TestingConfig(t testing.TB) *ClientConfig {
	cfg := NewDefaultClientConfig()
	cfg.ListenHost = LoopbackListenHost
	cfg.NoDHT = true
	cfg.DataDir = t.TempDir()
	cfg.DisableTrackers = true
	cfg.NoDefaultPortForwarding = true
	cfg.DisableAcceptRateLimiting = true
	cfg.ListenPort = 0
	cfg.KeepAliveTimeout = time.Millisecond
	cfg.MinPeerExtensions.SetBit(pp.ExtensionBitFast, true)
	cfg.Logger = log.Default.WithContextText(t.Name())
	// 2 would suffice for the greeting test, but 5 is needed for a few other tests. This should be
	// something slightly higher than the usual chunk size, so it gets tickled in some tests.
	cfg.MaxAllocPeerRequestDataPerConn = 5
	//cfg.Debug = true
	//cfg.Logger = cfg.Logger.WithText(func(m log.Msg) string {
	//	t := m.Text()
	//	m.Values(func(i interface{}) bool {
	//		t += fmt.Sprintf("\n%[1]T: %[1]v", i)
	//		return true
	//	})
	//	return t
	//})
	return cfg
}
