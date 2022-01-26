package main

import (
	"log"
	"net"

	"github.com/adminbaintex/gomqtt/stream"
	"github.com/adminbaintex/mqtt-server/server"
)

var done = make(chan struct{})

type exampleHandler struct{}

func (handler *exampleHandler) ServeMQTT(conn net.Conn, s stream.Stream) {
	defer func() {
		log.Println(conn.RemoteAddr(), "CLOSED")
		s.Close()
		close(done)
	}()
	log.Println(conn.RemoteAddr(), "CONNECTED")
}

func main() {

	s := server.NewServer(&exampleHandler{}, false)
	if err := s.ListenAndServe("localhost:1337"); err != nil {
		log.Println(err)
		return
	}

	c, _ := net.Dial("tcp", "localhost:1337")
	c.Close()

	<-done
	if err := s.Stop(); err != nil {
		log.Println(err)
	}

	log.Println("Server stopped correctly")

	// Output:
	// remote addr CONNECTED
	// remote addr CLOSE
}
