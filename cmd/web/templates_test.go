package main

import (
	"testing"
	"time"

	"github.com/jagfiend/snippetbox/internal/assert"
)

// example of simple unit test
func TestHumanDate(t *testing.T) {
	tm := time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC)
	got := humanReadableDate(tm)
	want := "17 Mar 2024 at 10:15"
	assert.Equal(t, got, want)
}

// example of 'table based' tests, tests multiple scenarios
func TestHumnDates(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			"UTC",
			time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			"17 Mar 2024 at 10:15",
		},
		{
			"empty",
			time.Time{},
			"",
		},
		{
			"CET",
			time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			"17 Mar 2024 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanReadableDate(tt.tm)
			assert.Equal(t, got, tt.want)
		})
	}
}
