package multiplexer

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const defaultTimeout = 10 * time.Second

type Instance struct {
	instance cmux.CMux
	servers  []Server
	port     uint16
	timeout  time.Duration
	logger   log.Logger
}

func New(port uint16) (*Instance, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}

	server := cmux.New(listener)

	return &Instance{server, make([]Server, 0), port, defaultTimeout, nil}, nil
}

func (multiplexer *Instance) Port() uint16 {
	return multiplexer.port
}

func (multiplexer *Instance) Timeout() time.Duration {
	return multiplexer.timeout
}

func (multiplexer *Instance) WithTimeout(timeout time.Duration) *Instance {
	multiplexer.timeout = timeout

	return multiplexer
}

func (multiplexer *Instance) WithLogger(logger log.Logger) *Instance {
	multiplexer.logger = logger

	return multiplexer
}

func (multiplexer *Instance) WithFiber(server *fiber.App) *Instance {
	multiplexer.servers = append(multiplexer.servers, newFiberAdapter(server))

	return multiplexer
}

func (multiplexer *Instance) WithGRPC(server *grpc.Server) *Instance {
	multiplexer.servers = append(multiplexer.servers, newGRPCAdapter(server))

	return multiplexer
}

func (multiplexer *Instance) Stop() {
	var wg sync.WaitGroup
	wg.Add(len(multiplexer.servers))

	for _, server := range multiplexer.servers {
		s := server
		go func() {
			s.Stop()
			wg.Done()
		}()
	}

	wg.Wait()
	multiplexer.instance.Close()
}

func (multiplexer *Instance) Serve() error {
	errs := make(chan error)

	for _, server := range multiplexer.servers {
		s := server
		go func() {
			err := s.Serve(multiplexer.instance)
			errs <- err
		}()
	}

	go func() {
		err := multiplexer.instance.Serve()
		errs <- err
	}()

	err := <-errs
	multiplexer.Stop()

	return err
}

func (multiplexer *Instance) ServeGracefully() error {
	errs := make(chan error, 1)
	defer close(errs)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		err := multiplexer.Serve()
		errs <- err
	}()

	if multiplexer.logger != nil {
		multiplexer.logger.Info("Serving servers: ", len(multiplexer.servers))
	}

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		if multiplexer.logger != nil {
			multiplexer.logger.Warn("Stopping servers")
		}

		done := make(chan struct{})
		go func() {
			multiplexer.Stop()
			close(done)
		}()

		select {
		case <-done:
			if multiplexer.logger != nil {
				multiplexer.logger.Info("Graceful stop completed")
			}

			return nil
		case <-time.After(multiplexer.timeout):
			if multiplexer.logger != nil {
				multiplexer.logger.Error("Graceful timeout reached")
			}

			return nil
		}
	}
}
