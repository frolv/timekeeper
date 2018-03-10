package tk

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsePeriod(t *testing.T) {
	var d time.Duration
	var err error

	d, err = ParsePeriod("5M")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(5*time.Minute), d, "5M should equal 5 minutes")

	d, err = ParsePeriod("3h")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(3*time.Hour), d, "3h should equal 3 hours")

	d, err = ParsePeriod("1d")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(24*time.Hour), d, "1d should equal 1 day")

	d, err = ParsePeriod("2w")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(14*24*time.Hour), d, "2w should equal 2 weeks")

	d, err = ParsePeriod("1m")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(31*24*time.Hour), d, "1m should equal 31 days")

	d, err = ParsePeriod("1y")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(365*24*time.Hour), d, "1y should equal 365 days")

	d, err = ParsePeriod("1z")
	assert.NotNil(t, err, "z is an invalid period unit")

	d, err = ParsePeriod("3")
	assert.NotNil(t, err, "periods must provide a unit")

	d, err = ParsePeriod("1000y")
	assert.NotNil(t, err, "periods cannot not overflow a time.Duration")
}
