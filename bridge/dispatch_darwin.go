package bridge

import (
	"github.com/progrium/shelldriver/dispatch"
)

// Dispatch uses the shell API to schedule work in the main UI thread
func Dispatch(fn func() error) error {
	return dispatch.Do(fn).Wait()
}
