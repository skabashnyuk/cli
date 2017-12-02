package command

import (
	"io"

)

// Streams is an interface which exposes the standard input and output streams
type Streams interface {
	In() *InStream
	Out() *OutStream
	Err() io.Writer
}

// Cli represents the docker command line client.
type Cli interface {
	Out() *OutStream
	Err() io.Writer
	In() *InStream
	SetIn(in *InStream)
}

type DockerCli struct {
	in             *InStream
	out            *OutStream
	err            io.Writer
	defaultVersion string
}


// NewCli returns a NewCli instance with IO output and error streams set by in, out and err.
func NewCli(in io.ReadCloser, out, err io.Writer) *DockerCli {
	return &DockerCli{in: NewInStream(in), out: NewOutStream(out), err: err}
}