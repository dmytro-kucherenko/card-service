package multiplexer

import (
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type grpcAdapter struct {
	server *grpc.Server
}

func newGRPCAdapter(server *grpc.Server) Server {
	return &grpcAdapter{server}
}

func (adapter *grpcAdapter) Serve(server cmux.CMux) error {
	listener := server.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	return adapter.server.Serve(listener)
}

func (adapter *grpcAdapter) Stop() {
	adapter.server.GracefulStop()
}
