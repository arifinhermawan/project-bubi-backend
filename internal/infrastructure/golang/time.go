package golang

import (
	// golang package
	"time"
)

const (
	locationAsiaJakarta = "Asia/Jakarta"
)

var (
	// for mocking purpose
	loadLocation = time.LoadLocation
	timeNowIn    = time.Now().In
)

// GetTimeGMT7 will get current time in GMT+7
func (g *Golang) GetTimeGMT7() time.Time {
	location, _ := loadLocation(locationAsiaJakarta)
	return timeNowIn(location)
}
