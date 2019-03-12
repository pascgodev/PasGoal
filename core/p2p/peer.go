package p2p

import (
	"fmt"
	"github.com/pasgo/pasgo/config"
	"net"
	"strconv"
)

type Peer struct {
	LocalPort int
	Listener  net.Listener
	ConnPool  []net.Conn
}

func InitPeer(port int) *Peer {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
	}

	localPeer := &Peer{
		LocalPort: port,
		Listener:  ln,
		ConnPool:  make([]net.Conn, 0),
	}

	localPeer.Bootstrap()

	return localPeer
}

// Init will connect to the bootstrap nodes and put them into ConnPool
func (p *Peer) Bootstrap() {
	peers := config.GetBootstrapPeers()
	for _, peerAddress := range peers {
		conn, err := net.Dial("tcp", peerAddress)
		if err != nil {
			fmt.Println("The Bootstrap peer", peerAddress, "is unreachable")
			continue
		}
		p.ConnPool = append(p.ConnPool, conn)
	}
}

func (p *Peer) Start() {
	fmt.Println("Local peer starts at port:", p.LocalPort)
	for {
		conn, err := p.Listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		p.ConnPool = append(p.ConnPool, conn)

		go p.handleConnection(conn)
	}
	//TODO: Then Say Hello to all conn
	//
}
