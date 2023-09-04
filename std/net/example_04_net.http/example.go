package main

import (
	"context"
	"net"
	"net/http"
)

type conn struct {
	server *Server
	cancelCtx context.CancelFunc
	rwc net.Conn
}

type Server struct {
	Addr string
	BaseContext func(net.Listener) context.Context
	ConnContext func(ctx context.Context, c net.Conn) context.Context
	// More fields
}

func (srv *Server) ListenAndServe() error {
	addr := "localhost:8011"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}

func (srv *Server) Serve(l net.Listener) error {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, http.ServerContextKey, srv)
	for {
		rw, err := l.Accept()
		if err != nil {
			return err
		}

		c := srv.newConn(rw)
		go c.serve(ctx)
	}
}


func (c *conn) serve(ctx context.Context) {
	// omitted
	// for {
	// 	w, err := c.readRequest(ctx)
	// 	if err != nil {
	// 		return 
	// 	}

	// 	serverHandler{c.server}.ServeHTTP(w, w.req)
		
	// 	w.cancelCtx()
	// }
}
func (srv *Server) newConn(rwc net.Conn) *conn {
	c := &conn{
		server: srv,
		rwc:    rwc,
	}
	return c
}

