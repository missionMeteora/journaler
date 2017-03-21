package journaler

import "testing"

func TestJournaler(t *testing.T) {
	jj := New("Main service", "Sub service")
	jj.Success("Database entry posted")
	jj.Notification("CPU temperatures are at 40*C")
	jj.Warning("Update to remote server has failed")
	jj.Error("Danger, Will Robinson!")
	jj.Debug(map[string]string{"foo": "bar"})
}

func TestJournal(t *testing.T) {
	j.Success("Database entry posted")
	j.Notification("CPU temperatures are at 40*C")
	j.Warning("Update to remote server has failed")
	j.Error("Danger, Will Robinson!")
	j.Debug(map[string]string{"foo": "bar"})
}

func TestRoot(t *testing.T) {
	Success("Database entry posted")
	Notification("CPU temperatures are at 40*C")
	Warning("Update to remote server has failed")
	Error("Danger, Will Robinson!")
	Debug(map[string]string{"foo": "bar"})
}

func TestOutput(t *testing.T) {
	jj := New("Main service", "Sub service")
	jj.Output("Compliment", "green", "You smell nice.")
	jj.Output("System", "default", "Server will be rebooting for maintenance in 45 minutes")
	jj.Output("System", "yellow", "Server will be rebooting for maintenance in 5 minutes")
	jj.Output("System", "red", "Server will be rebooting for maintenance in 5 seconds")
}

func TestCustomLabel(t *testing.T) {
	jj := New("Main service", "Sub service")
	SetLabel("error", "uh oh.")
	jj.Error("Danger, Will Robinson!")
}
