package hizuke

import (
	"time"

	"github.com/agatan/timejump"
)

const layout string = "20060102"

// Today returns today string in yyyyMMdd format.
func Today() string {
	return today().Format(layout)
}

// Yesterday returns yesterday string in yyyyMMdd format.
func Yesterday() string {
	return yesterday().Format(layout)
}

func today() time.Time {
	return timejump.Now()
}

func yesterday() time.Time {
	return today().Add(-1 * time.Duration(24) * time.Hour)
}
