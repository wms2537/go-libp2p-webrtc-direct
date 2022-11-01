package libp2pwebrtcdirect

import (
	"errors"
	"net"
	"strconv"

	ma "github.com/multiformats/go-multiaddr"
	mafmt "github.com/multiformats/go-multiaddr-fmt"
	manet "github.com/multiformats/go-multiaddr/net"
)

var webrtcDirectMA = ma.StringCast("/http/p2p-webrtc-direct")

var webrtcDirectMatcher = mafmt.And(mafmt.IP, mafmt.Base(ma.P_TCP), mafmt.Base(ma.P_HTTP), mafmt.Base(ma.P_P2P_WEBRTC_DIRECT))

func toWebrtcDirectMultiaddr(na net.Addr) (ma.Multiaddr, error) {
	addr, err := manet.FromNetAddr(na)
	if err != nil {
		return nil, err
	}
	if _, err := addr.ValueForProtocol(ma.P_TCP); err != nil {
		return nil, errors.New("not a TCP address")
	}
	return addr.Encapsulate(webrtcDirectMA), nil
}

func stringToWebrtcDirectMultiaddr(str string) (ma.Multiaddr, error) {
	host, portStr, err := net.SplitHostPort(str)
	if err != nil {
		return nil, err
	}
	port, err := strconv.ParseInt(portStr, 10, 32)
	if err != nil {
		return nil, err
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return nil, errors.New("failed to parse IP")
	}
	return toWebrtcDirectMultiaddr(&net.TCPAddr{IP: ip, Port: int(port)})
}
