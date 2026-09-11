package loggerratelimiter

import "testing"

// The operation sequence from the problem statement. A message is
// printed only if the same message has not been printed in the last ten
// seconds, so "foo" at t=11 is allowed but "bar" at t=8 is not.
func TestLoggerShouldPrintMessage(t *testing.T) {
	logger := Constructor()

	ops := []struct {
		timestamp int
		message   string
		want      bool
	}{
		{1, "foo", true},
		{2, "bar", true},
		{3, "foo", false},
		{8, "bar", false},
		{10, "foo", false},
		{11, "foo", true},
	}
	for _, op := range ops {
		if got := logger.ShouldPrintMessage(op.timestamp, op.message); got != op.want {
			t.Errorf("ShouldPrintMessage(%d, %q) = %v, want %v", op.timestamp, op.message, got, op.want)
		}
	}
}
