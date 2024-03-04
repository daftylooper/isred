package Server

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan []byte
}

func (s *Server) GetMsgch() chan []byte {
	return s.msgch
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		//setting 10 arbitrarily
		msgch: make(chan []byte, 10),
	}
}

// sets listener upon start
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	//if listener closes, then don't set listener in struct
	defer ln.Close()
	s.ln = ln

	go s.AcceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}

		fmt.Println("New connection to server:", conn.RemoteAddr())

		go s.ReadLoop(conn)
	}
}

func (s *Server) ReadLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error", err)
			continue
		}

		s.msgch <- buf[:n]

		conn.Write([]byte("roger that!\n"))
	}
}
