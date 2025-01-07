package processrestart

import "github.com/khulnasoft/khulnasoft/lib/errors"

// CanRestart reports whether the current set of Khulnasoft processes can
// be restarted.
func CanRestart() bool {
	return usingGoremanServer
}

// Restart restarts the current set of Khulnasoft processes associated with
// this server.
func Restart() error {
	if !CanRestart() {
		return errors.New("reloading site is not supported")
	}
	if usingGoremanServer {
		return restartGoremanServer()
	}
	return errors.New("unable to restart processes")
}
