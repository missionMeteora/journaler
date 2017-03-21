package journaler

import (
	"errors"
	"testing"
)

func TestJournaler(t *testing.T) {
	jj := New("Main service", "Sub service")
	jj.Success("Database entry posted")
	jj.Notification("CPU temperatures are at 40*C")
	jj.Warn("Update to remote server has failed")
	jj.Error(errors.New("Danger, Will Robinson!"))
	jj.Debug(map[string]string{"foo": "bar"})
	j.Debug("Parent debug test")
	// Testing custom labels
	SetLabel("error", "uh oh.")
	jj.Error(errors.New("Danger, Will Robinson!"))

	// Testing custom output
	jj.Output("Compliment", "green", "You smell nice.")
	jj.Output("System", "default", "Server will be rebooting for maintenance in 45 minutes")
	jj.Output("System", "yellow", "Server will be rebooting for maintenance in 5 minutes")
	jj.Output("System", "red", "Server will be rebooting for maintenance in 5 seconds")

}
