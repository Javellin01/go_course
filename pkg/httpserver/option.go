package httpserver

import "time"

type Option func(*server)

// Addr optionally specifies the TCP address for the server to listen on,
// in the form "host:port". If empty, ":http" (port 80) is used.
// The service names are defined in RFC 6335 and assigned by IANA.
func Addr(addr string) Option {
	return func(s *server) {
		s.server.Addr = addr
	}
}

// ReadTimeout is the maximum duration for reading the entire
// request, including the body. A zero or negative value means
// there will be no timeout.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout is the maximum duration before timing out
// writes of the response. It is reset whenever a new
// request's header is read. Like ReadTimeout, it does not
// let Handlers make decisions on a per-request basis.
// A zero or negative value means there will be no timeout.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.WriteTimeout = timeout
	}
}

// IdleTimeout is the maximum amount of time to wait for the
// next request when keep-alives are enabled. If IdleTimeout
// is zero, the value of ReadTimeout is used. If both are
// zero, there is no timeout.
func IdleTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.IdleTimeout = timeout
	}
}
