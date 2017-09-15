package ranges_test

import (
	"testing"
	"time"

	"github.com/dolanor/ranges"
)

func TestPeriod(t *testing.T) {
	start := time.Date(2017, time.Month(8), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2017, time.Month(8), 1, 1, 0, 0, 0, time.UTC)
	got := ranges.Period(start, end, 10*time.Minute)
	expected := []time.Time{
		time.Date(2017, time.Month(8), 1, 0, 00, 0, 0, time.UTC),
		time.Date(2017, time.Month(8), 1, 0, 10, 0, 0, time.UTC),
		time.Date(2017, time.Month(8), 1, 0, 20, 0, 0, time.UTC),
		time.Date(2017, time.Month(8), 1, 0, 30, 0, 0, time.UTC),
		time.Date(2017, time.Month(8), 1, 0, 40, 0, 0, time.UTC),
		time.Date(2017, time.Month(8), 1, 0, 50, 0, 0, time.UTC),
	}
	notExpected := time.Date(2017, time.Month(8), 1, 1, 0, 0, 0, time.UTC)

	lgot, lexp := len(got), len(expected)
	if lgot != lexp {
		t.Fatalf("interval size don't match\n\tgot:      %s\n\texpected: %s", lgot, lexp)
	}

	for i := 0; i < lgot; i++ {
		if !got[i].Equal(expected[i]) {
			t.Fatalf("interval don't match\n\tgot:      %s\n\texpected: %s", got, expected)
		}
		if got[i].Equal(notExpected) {
			t.Fatalf("interval includes non expected date\n\tgot:      %s\n\texpected: %s", got, expected)
		}
	}
}
