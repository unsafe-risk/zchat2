package main

import (
	"net"

	"github.com/lemon-mint/zchat2/internal/envaddr"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	ln      net.Listener
	stopCh  chan struct{}
	readyCh chan error
}

func (g *Server) run() {

}

func (g *Server) stop() {

}

func (g *Server) waitReady() error {
	err, nak := <-g.readyCh
	if !nak {
		return nil
	}
	return err
}

func startServer() (*Server, error) {
	addr := envaddr.Get(":53267")

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := Server{
		ln:      ln,
		stopCh:  make(chan struct{}),
		readyCh: make(chan error, 1),
	}

	go srv.run()
	err = srv.waitReady()
	if err != nil {
		srv.stop()
		return nil, err
	}

	return &srv, err
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano

	srv, err := startServer()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
	defer srv.stop()

}
