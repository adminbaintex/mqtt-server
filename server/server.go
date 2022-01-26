package server

import (
	"log"
	"net"

	"github.com/adminbaintex/gomqtt/stream"
	"github.com/armon/go-proxyproto"
)

// MQTTHandler will receive new connections as streams.
type MQTTHandler interface {
	ServeMQTT(net.Conn, stream.Stream)
}

// Server manages multiple Configurations and yields new connection as
// streams to the Handler.
type Server struct {
	// The Handler that receives new Streams.
	handler MQTTHandler

	// The currently running configurations.
	listener net.Listener

	// Use Proxy protocol
	ProxyProcotol bool
}

// NewServer returns a new Server.
func NewServer(handler MQTTHandler, proxyProcotol bool) *Server {
	return &Server{handler: handler, ProxyProcotol: proxyProcotol}
}

// ListenAndServe will run a simple TCP server.
func (s *Server) ListenAndServe(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if s.ProxyProcotol {
		// Wrap listener in a proxyproto listener
		s.listener = &proxyproto.Listener{Listener: l}
	} else {
		s.listener = l
	}

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Println(err)
				return
			}
			go s.handler.ServeMQTT(conn, stream.NewNetStream(conn))
		}
	}()

	return nil
}

// Stop will stop listening to new connections
func (s *Server) Stop() error {
	return s.listener.Close()
}
