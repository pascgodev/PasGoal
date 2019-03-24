package p2p

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/pasgo/pasgo/config"
	"net"
)

type RemotePeer struct {
	RemoteIP   string
	RemotePort uint16
	PeerKey    PeerKey
	LastTime   uint32
	LocalPort  uint16
}

type PeerKey uint32

type LocalPeer struct {
	LocalPort     int
	LocalListener *net.TCPListener
	ConnPool      []*net.TCPConn
	RemotePeers   []*RemotePeer
	LocalKey      PeerKey

	addConn chan *net.TCPConn
}

func InitPeer(port int, withBootstrapPeers bool) *LocalPeer {

	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: port,
	})
	if err != nil {
		glog.Warningln(err)
	}

	localPeer := &LocalPeer{
		LocalPort:     port,
		LocalListener: listener,
		ConnPool:      []*net.TCPConn{},
		RemotePeers:   []*RemotePeer{},
	}

	if withBootstrapPeers == true {
		localPeer.Bootstrap()
	}

	return localPeer
}

// Init will connect to the bootstrap nodes and put them into ConnPool
func (p *LocalPeer) Bootstrap() {
	peerConfigList := config.GetBootstrapPeers()
	for _, peerConfig := range peerConfigList {
		addr, err := net.ResolveTCPAddr("tcp4", peerConfig)
		if err != nil {
			fmt.Println("The Bootstrap peer", peerConfig, "is wrong")
			continue
		}

		conn, err := net.DialTCP("tcp4", nil, addr)
		if err != nil {
			fmt.Println("The Bootstrap peer", peerConfig, "is unreachable")
			continue
		}
		p.ConnPool = append(p.ConnPool, conn)
	}
	//TODO: get more peers (additional) from DB
}

func (p *LocalPeer) Start() {
	fmt.Println("Local peer starts at port:", p.LocalPort)
	for _, conn := range p.ConnPool {
		go p.handleConnection(conn)
		p.SendHello(conn)
	}

	p.addConn = make(chan *net.TCPConn, 1)
	// Run routines waiting for new peers
	go p.waitForPeerFromPeer()

	go p.waitForPeerFromListener()
}

func (p *LocalPeer) waitForPeerFromListener() {
	for {
		conn, err := p.LocalListener.AcceptTCP()
		if err != nil {
			glog.Warningln(err)
		}
		glog.Infoln("New Conn from", conn.RemoteAddr().String())
		p.addConn <- conn
	}
}

func (p *LocalPeer) waitForPeerFromPeer() {
	for {
		select {
		case newConn := <-p.addConn:
			glog.Infoln("add new conn", newConn.RemoteAddr())
			p.ConnPool = append(p.ConnPool, newConn)
			go p.handleConnection(newConn)
		}
	}
}

func (p *LocalPeer) CloseAll() {
	for _, conn := range p.ConnPool {
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed to close the conn", conn, "for", err)
		}
	}
}
