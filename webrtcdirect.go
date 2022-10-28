package libp2pwebrtcdirect

import (
	logging "github.com/ipfs/go-log"
	smux "github.com/libp2p/go-libp2p/core/network"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("webrtcdirect-tpt")

var webrtcma, _ = ma.NewMultiaddr("/p2p-webrtc-direct")
var httpma, _ = ma.NewMultiaddr("/http")

var _ tpt.Transport = &Transport{}
var _ tpt.CapableConn = &Conn{}
var _ tpt.Listener = &Listener{}

var _ smux.MuxedStream = &Stream{}
