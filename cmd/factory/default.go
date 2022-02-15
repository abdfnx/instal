package factory

import (
	"github.com/abdfnx/instal/ios"
)

type Factory struct {
	IOStreams *ios.IOStreams
}

func New() *Factory {
	f := &Factory{}

	f.IOStreams = ioStreams(f)

	return f
}

func ioStreams(f *Factory) *ios.IOStreams {
	io := ios.System()

	return io
}
