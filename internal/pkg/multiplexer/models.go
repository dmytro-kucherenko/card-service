package multiplexer

import (
	"github.com/soheilhy/cmux"
)

type Server interface {
	Serve(cmux.CMux) error
	Stop()
}
