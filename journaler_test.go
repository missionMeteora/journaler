package journaler

import "testing"

func TestJournaler(t *testing.T) {
	test(t, New("Main service"))
}

func TestDeepJournaler(t *testing.T) {
	test(t, New("Main service", "Sub service", "Child process"))
}

func TestJournal(t *testing.T) {
	test(t, j)
}

func TestRoot(t *testing.T) {
	Success("Database entry posted")
	Notification("CPU temperatures are at 40*C")
	Warning("Update to remote server has failed")
	Error("Danger, Will Robinson!")
	Debug(map[string]string{"foo": "bar"})
}

func TestCustomLabel(t *testing.T) {
	jj := New("Main service", "Sub service")
	SetLabel("error", "uh oh.")
	jj.Error("Danger, Will Robinson!")
}

type Tester interface {
	Success(val interface{})
	Notification(val interface{})
	Warning(val interface{})
	Error(val interface{})
	Debug(val interface{})
	Output(label, color string, val interface{})
}

func test(t *testing.T, ti Tester) {
	ti.Success("Database entry posted")
	ti.Notification("CPU temperatures are at 40*C")
	ti.Warning("Update to remote server has failed")
	ti.Error("Danger, Will Robinson!")
	ti.Debug(map[string]string{"foo": "bar"})
	ti.Output("Compliment", "green", "You smell nice.")
	ti.Output("System", "default", "Server will be rebooting for maintenance in 45 minutes")
	ti.Output("System", "yellow", "Server will be rebooting for maintenance in 5 minutes")
	ti.Output("System", "red", "Server will be rebooting for maintenance in 5 seconds")
}
