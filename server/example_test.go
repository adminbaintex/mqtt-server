package server

import (
	"log"
	"net"

	"github.com/adminbaintex/gomqtt/stream"
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

func ExampleServer() {

	server := NewServer(&exampleHandler{}, false)
	if err := server.ListenAndServe("localhost:1337"); err != nil {
		log.Println(err)
		return
	}

	c, _ := net.Dial("tcp", "localhost:1337")
	c.Close()

	<-done
	server.Stop()

	// Output:
	// remote addr CONNECTED
	// remote addr CLOSE
}
