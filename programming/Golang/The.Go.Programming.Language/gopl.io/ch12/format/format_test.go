package format

import (
	"strings"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d = 1 * time.Nanosecond

	if got, want := Any(x), "1"; got != want {
		t.Errorf("Any(%v) = %v", x, got)
	}

	if got, want := Any(d), "1"; got != want {
		t.Errorf("Any(%v) = %v", x, got)
	}

	if got, want := Any([]int64{x}), "[]int64 0x"; !strings.Contains(got, want) {
		t.Errorf("Any(%v) = %v", x, got)
	}

	if got, want := Any([]time.Duration{d}), "[]time.Duration 0x"; !strings.Contains(got, want) {
		t.Errorf("Any(%v) = %v", x, got)
	}
}
