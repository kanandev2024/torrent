module github.com/anacrolix/torrent

require (
	bazil.org/fuse v0.0.0-20200407214033-5883e5a4b512
	crawshaw.io/sqlite v0.3.3-0.20210127221821-98b1f83c5508
	github.com/RoaringBitmap/roaring v0.9.4 // indirect
	github.com/alexflint/go-arg v1.3.0
	github.com/anacrolix/chansync v0.1.0
	github.com/anacrolix/confluence v1.8.0 // indirect
	github.com/anacrolix/dht/v2 v2.10.3
	github.com/anacrolix/envpprof v1.1.1
	github.com/anacrolix/go-libutp v1.0.4
	github.com/anacrolix/log v0.9.0
	github.com/anacrolix/missinggo v1.3.0
	github.com/anacrolix/missinggo/perf v1.0.0
	github.com/anacrolix/missinggo/v2 v2.5.2
	github.com/anacrolix/multiless v0.1.1-0.20210529082330-de2f6cf29619
	github.com/anacrolix/sync v0.4.0
	github.com/anacrolix/tagflag v1.3.0
	github.com/anacrolix/upnp v0.1.2-0.20200416075019-5e9378ed1425
	github.com/anacrolix/utp v0.1.0
	github.com/bradfitz/iter v0.0.0-20191230175014-e8f45d346db8
	github.com/davecgh/go-spew v1.1.1
	github.com/dustin/go-humanize v1.0.0
	github.com/edsrzf/mmap-go v1.0.0
	github.com/elliotchance/orderedmap v1.4.0
	github.com/frankban/quicktest v1.11.3
	github.com/fsnotify/fsnotify v1.4.9
	github.com/google/btree v1.0.1
	github.com/gorilla/websocket v1.4.2
	github.com/jessevdk/go-flags v1.4.0
	github.com/pion/datachannel v1.4.21
	github.com/pion/ice/v2 v2.1.12 // indirect
	github.com/pion/interceptor v0.0.15 // indirect
	github.com/pion/rtp v1.7.1 // indirect
	github.com/pion/srtp/v2 v2.0.5 // indirect
	github.com/pion/webrtc/v3 v3.0.32
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	go.etcd.io/bbolt v1.3.6
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

go 1.16

exclude bazil.org/fuse v0.0.0-20200419173433-3ba628eaf417

exclude github.com/willf/bitset v1.2.0
