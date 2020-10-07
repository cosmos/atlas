package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/urfave/cli/v2"
)

// BufferReader is implemented by types that read from a string buffer.
type BufferReader interface {
	io.Reader
	Reset(string)
}

// BufferWriter is implemented by types that write to a buffer.
type BufferWriter interface {
	io.Writer
	Reset()
	Bytes() []byte
	String() string
}

// ApplyMockIO replaces stdout/err with buffers that can be used during testing.
// Returns an input BufferReader and an output BufferWriter.
func ApplyMockIO(app *cli.App) (BufferReader, BufferWriter) {
	mockIn := strings.NewReader("")
	mockOut := bytes.NewBufferString("")

	app.Writer = mockOut
	app.ErrWriter = mockOut

	return mockIn, mockOut
}

// ExecCmd executes a command in a test environment with a given Context and set
// of arguments. If an error occurs, it is written to the App's ErrWriter.
func ExecCmd(ctx context.Context, app *cli.App, args []string) error {
	err := app.RunContext(ctx, args)
	if err != nil {
		fmt.Fprintln(app.ErrWriter, err.Error())
	}

	return err
}
