package hizuke

import (
	"testing"
	"time"

	"github.com/agatan/timejump"
)

var newyear = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

func TestToday(t *testing.T) {
	timejump.Activate()
	defer timejump.Deactivate()

	timejump.Stop()
	timejump.Jump(newyear)

	expected := "20190101"
	actual := Today()
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}

func TestYesterday(t *testing.T) {
	timejump.Activate()
	defer timejump.Deactivate()

	timejump.Stop()
	timejump.Jump(newyear)

	expected := "20181231"
	actual := Yesterday()
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}
