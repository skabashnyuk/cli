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

type CheCli struct {
	in             *InStream
	out            *OutStream
	err            io.Writer
	defaultVersion string
}



// DefaultVersion returns api.defaultVersion or DOCKER_API_VERSION if specified.
func (cli *CheCli) DefaultVersion() string {
	return cli.defaultVersion
}


// Out returns the writer used for stdout
func (cli *CheCli) Out() *OutStream {
	return cli.out
}

// Err returns the writer used for stderr
func (cli *CheCli) Err() io.Writer {
	return cli.err
}

// SetIn sets the reader used for stdin
func (cli *CheCli) SetIn(in *InStream) {
	cli.in = in
}

// In returns the reader used for stdin
func (cli *CheCli) In() *InStream {
	return cli.in
}

// NewCli returns a NewCli instance with IO output and error streams set by in, out and err.
func NewCli(in io.ReadCloser, out, err io.Writer) *CheCli {
	return &CheCli{in: NewInStream(in), out: NewOutStream(out), err: err}
}