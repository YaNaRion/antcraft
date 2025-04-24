package gateway

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"main/infra"
	"main/service/socket"
	"net/http"
	"time"

	. "github.com/onsi/gomega"

	socketio "github.com/doquangtan/socket.io/v4"
)

var _ = Describe("Gateway", func() {
	const port = "3030"
	var err error
	var (
		db      *infra.DB
		io      *socketio.Io
		gateway *Gateway
		mux     *http.ServeMux
		server  *http.Server
	)

	mux = http.NewServeMux()

	BeforeEach(func() {
		io = socketio.New()
		mux = http.NewServeMux()
		db = &infra.DB{
			DB:   nil,
			Conf: nil,
		}
		var socketService = socket.NewSocketService(db)
		gateway, err = factoryGateway(io, socketService)
		Expect(err).To(BeNil())
		io.OnConnection(gateway.handleConnection)
		mux.Handle("/socket.io", io.HttpHandler())
		server = &http.Server{
			Addr:    fmt.Sprintf("localhost:%s", port),
			Handler: mux,
		}
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("Server error: %v\n", err)
			}
		}()
	})

	AfterEach(func() {
		gateway.io.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Error shutting down server: %s\n", err)
		} else {
			fmt.Println("Server gracefully stopped")
		}
	})
	Describe("When a robot socket disconnects", func() {
	})
})
