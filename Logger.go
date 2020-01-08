package logging

import "log"

// Logger defines the methods that any logging type must implement.  It is
// designed specifically to allow a *log.Logger to be resolved in dependency
// injection.
type Logger interface {
	Printf(format string, v ...interface{})
}

var _ Logger = &log.Logger{}
