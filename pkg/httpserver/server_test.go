package httpserver_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"gitlab.llo.su/adspend/blank-service/pkg/httpserver"
)

func TestServer(t *testing.T) {
	opts := []httpserver.Option{
		httpserver.Addr(":8080"),
		httpserver.ReadTimeout(15 * time.Second),
		httpserver.WriteTimeout(15 * time.Second),
		httpserver.IdleTimeout(30 * time.Second),
	}
	srv := httpserver.New(nil, opts...)

	done := make(chan struct{})

	go srv.Start()

	go func() {
		t.Log("\nStarting a graceful shutdown of the server")

		if err := srv.Shutdown(context.TODO()); err == nil {
			t.Log("\nServer graceful shutdown")
		} else {
			t.Errorf("\nNo error expected, but there is an error: %s", err)
		}

		close(done)
	}()

	err := <-srv.Notify()
	if err != http.ErrServerClosed {
		t.Errorf(
			"\nErrors not equals:\n- expected: %s\n- actual: %s",
			http.ErrServerClosed,
			err,
		)
	}

	<-done
}
