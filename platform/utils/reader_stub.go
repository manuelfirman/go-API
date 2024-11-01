package utils

import "io"

// stubReadCloser is a stub implementation.
type StubReadCloser struct {
	reader io.Reader
	err    error
}

// Read reads from the stubReadCloser and returns the specified error if set.
// If no error is set, it reads from the underlying reader.
func (m *StubReadCloser) Read(p []byte) (n int, err error) {
	if m.err != nil {
		return 0, m.err
	}
	return m.reader.Read(p)
}

// It is implemented to satisfy the io.ReadCloser interface.
func (m *StubReadCloser) Close() error {
	return nil
}

// NewStubReadCloser creates a new StubReadCloser.
func NewStubReadCloserWithErr(reader io.Reader, err error) *StubReadCloser {
	if err == nil {
		err = io.ErrUnexpectedEOF // Usar io.ErrUnexpectedEOF por defecto si el error es nil
	}
	return &StubReadCloser{
		reader: reader,
		err:    err,
	}
}
