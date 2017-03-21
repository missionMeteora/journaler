package journaler

import (
	"errors"
	"testing"
)

func TestJournaler(t *testing.T) {
	j := New("Main service", "Sub service")
	j.Success("Database entry posted")
	j.Warn("Update to remote server has failed")
	j.Error(errors.New("Danger, Will Robinson!"))
}
