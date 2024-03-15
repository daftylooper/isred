package Server

import (
	"fmt"
	"isred/Buffer"
	"net"
	"time"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan []byte
	replybuf   *Buffer.Buffer
}

func (s *Server) GetMsgch() chan []byte {
	return s.msgch
}

func NewServer(listenAddr string, replybuf *Buffer.Buffer) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		//setting 10 arbitrarily
		msgch:    make(chan []byte, 10),
		replybuf: replybuf,
	}
}

// sets listener upon start
func (s *Server) Start() error {
	fmt.Println("Starting Up Server!")
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
		go s.WriteLoop(conn)
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
	}
}

func (s *Server) WriteLoop(conn net.Conn) {
	for {
		reply, _ := s.replybuf.GetCommand()
		// if err != nil {
		// 	fmt.Println("Couldn't Read from Reply Buffer:", err)
		// }
		if reply != "" {
			fmt.Println("server!", reply)
			conn.Write([]byte(reply + "\n"))
		}
		time.Sleep(10 * time.Millisecond)
	}
}
