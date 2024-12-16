// Package derrors defines internal error values to categorize the different
// types error semantics we support.
package derrors

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

var (
	// ErrNotFound indicates that a requested entity was not found (HTTP 404).
	ErrNotFound = errors.New("not found")

	// ErrInvalidArgument indicates that the input into the request is invalid in
	// some way (HTTP 400).
	ErrInvalidArgument = errors.New("invalid argument")

	// Unknown indicates that the error has unknown semantics.
	ErrUnknown = errors.New("unknown")
)

var codes = []struct {
	err  error
	code int
}{
	{ErrNotFound, http.StatusNotFound},
	{ErrInvalidArgument, http.StatusBadRequest},
}

// FromStatus generates an error according for the given status code. It uses
// the given format string and arguments to create the error string according
// to the fmt package. If format is the empty string, then the error
// corresponding to the code is returned unwrapped.
//
// If code is http.StatusOK, it returns nil.
func FromStatus(code int, format string, args ...any) error {
	if code == http.StatusOK {
		return nil
	}
	innerErr := ErrUnknown
	for _, e := range codes {
		if e.code == code {
			innerErr = e.err
			break
		}
	}
	if format == "" {
		return innerErr
	}
	return fmt.Errorf(format+": %w", append(args, innerErr)...)
}

// ToStatus returns a status code corresponding to err.
func ToStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	for _, e := range codes {
		if errors.Is(err, e.err) {
			return e.code
		}
	}
	return http.StatusInternalServerError
}

// Add adds context to the error.
// The result cannot be unwrapped to recover the original error.
// It does nothing when *errp == nil.
//
// Example:
//
//	defer derrors.Add(&err, "copy(%s, %s)", src, dst)
//
// See Wrap for an equivalent function that allows
// the result to be unwrapped.
func Add(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %v", fmt.Sprintf(format, args...), *errp)
	}
}

// Wrap adds context to the error and allows
// unwrapping the result to recover the original error.
//
// Example:
//
//	defer derrors.Wrap(&err, "copy(%s, %s)", src, dst)
//
// See Add for an equivalent function that does not allow
// the result to be unwrapped.
func Wrap(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}

// WrapStack is like Wrap, but adds a stack trace if there isn't one already.
func WrapStack(errp *error, format string, args ...any) {
	if *errp != nil {
		if se := (*StackError)(nil); !errors.As(*errp, &se) {
			*errp = NewStackError(*errp)
		}
		Wrap(errp, format, args...)
	}
}

// StackError wraps an error and adds a stack trace.
type StackError struct {
	Stack []byte
	err   error
}

// NewStackError returns a StackError, capturing a stack trace.
func NewStackError(err error) *StackError {
	// Limit the stack trace to 16K. Same value used in the errorreporting client,
	// cloud.google.com/go@v0.66.0/errorreporting/errors.go.
	var buf [16 * 1024]byte
	n := runtime.Stack(buf[:], false)
	return &StackError{
		err:   err,
		Stack: buf[:n],
	}
}

func (e *StackError) Error() string {
	return e.err.Error() // ignore the stack
}

func (e *StackError) Unwrap() error {
	return e.err
}

// WrapAndReport calls Wrap followed by Report.
func WrapAndReport(errp *error, format string, args ...any) {
	Wrap(errp, format, args...)
	if *errp != nil {
		Report(*errp)
	}
}

var reporter Reporter

// SetReporter the Reporter to use, for use by Report.
func SetReporter(r Reporter) {
	reporter = r
}

// Reporter is an interface used for reporting errors.
type Reporter interface {
	Report(err error, req *http.Request, stack []byte)
}

// Report uses the Reporter to report an error.
func Report(err error) {
	if reporter != nil {
		reporter.Report(err, nil, nil)
	}
}
