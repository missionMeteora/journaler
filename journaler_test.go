package journaler

import (
	"errors"
	"testing"
)

func TestJournaler(t *testing.T) {
	j := New("Main service", "Sub service")
	j.Success("Database entry posted")
	j.Notification("CPU temperatures are at 40*C")
	j.Warn("Update to remote server has failed")
	j.Error(errors.New("Danger, Will Robinson!"))
	j.Debug("debug test")

	// Testing custom labels
	SetLabel("error", "uh oh.")
	j.Error(errors.New("Danger, Will Robinson!"))

	// Testing custom output
	j.Output("Compliment", "green", "You smell nice.")
	j.Output("System", "default", "Server will be rebooting for maintenance in 45 minutes")
	j.Output("System", "yellow", "Server will be rebooting for maintenance in 5 minutes")
	j.Output("System", "red", "Server will be rebooting for maintenance in 5 seconds")

}
