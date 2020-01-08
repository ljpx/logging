package logging

import (
	"fmt"

	"github.com/ljpx/test"
)

// DummyLogger is a dummy implementation of Logger that simply records the
// messages logged to it.
type DummyLogger struct {
	messages map[string]int
}

var _ Logger = &DummyLogger{}

// NewDummyLogger creates a new, empty DummyLogger.
func NewDummyLogger() *DummyLogger {
	return &DummyLogger{
		messages: make(map[string]int),
	}
}

// Printf records a message into the DummyLogger.
func (d *DummyLogger) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	d.messages[msg]++
}

// AssertLogged asserts that the message specified was logged to the
// DummyLogger.
func (d *DummyLogger) AssertLogged(t test.T, format string, v ...interface{}) {
	t.Helper()

	msg := fmt.Sprintf(format, v...)
	c, ok := d.messages[msg]
	if !ok || c < 1 {
		t.Fatalf("expected message:\n\n%v\n\nto have been logged, but hadn't", msg)
	}
}
