package tk

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"timekeeper/lib/tkerr"
)

var periodRegex *regexp.Regexp

func init() {
	periodRegex = regexp.MustCompile(`^\d+[Mhdwmy]$`)
}

// Parse a string period into a time duration.
// Periods take the form number + unit, where number is an arbitrary integer,
// and unit is one of [Mhdwmy], representing minutes, hours, days, months, and
// years, respectively.
func ParsePeriod(period string) (time.Duration, error) {
	if !periodRegex.MatchString(period) {
		return 0, tkerr.Create(tkerr.InvalidPeriod)
	}

	// Regex guarantees this will succeed
	c, _ := strconv.Atoi(period[:len(period)-1])
	r := period[len(period)-1]
	var d time.Duration

	switch r {
	case 'M':
		d = time.Minute
	case 'h':
		d = time.Hour
	case 'd':
		d = time.Hour * 24
	case 'w':
		d = time.Hour * 24 * 7
	case 'm':
		d = time.Hour * 24 * 31
	case 'y':
		d = time.Hour * 24 * 365
	}

	// Guard against overflow
	if int64(c) > math.MaxInt64/int64(d) {
		return 0, tkerr.Create(tkerr.InvalidPeriod)
	}

	return time.Duration(c) * d, nil
}
