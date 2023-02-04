package golang

import (
	// golang package
	"testing"
	"time"

	// external package
	"github.com/stretchr/testify/assert"
)

func TestGolang_GetTimeGMT7(t *testing.T) {
	mockTime := time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local)
	loadLocationOri := loadLocation
	timeNowInOri := timeNowIn
	defer func() {
		loadLocation = loadLocationOri
		timeNowIn = timeNowInOri
	}()

	tests := []struct {
		name          string
		mockLocation  func(name string) (*time.Location, error)
		mockTimeNowIn func(loc *time.Location) time.Time
		want          time.Time
	}{
		{
			name: "success",
			mockLocation: func(name string) (*time.Location, error) {
				return &time.Location{}, nil
			},
			mockTimeNowIn: func(loc *time.Location) time.Time {
				return mockTime
			},

			want: mockTime,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			loadLocation = test.mockLocation
			timeNowIn = test.mockTimeNowIn

			g := &Golang{}

			got := g.GetTimeGMT7()
			assert.Equal(t, test.want, got)
		})
	}
}
