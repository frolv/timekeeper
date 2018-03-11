package osrs

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"timekeeper/lib/tkerr"
)

const osrsAPI = "http://services.runescape.com/m=hiscore_oldschool/index_lite.ws?player="

type SkillInfo struct {
	ID         uint
	Experience int
	Level      int
	Rank       int
	Hours      float64
}

func HiscoreLookup(username string) ([]SkillInfo, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	res, err := client.Get(osrsAPI + username)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, tkerr.Create(tkerr.InvalidAccount)
	} else if res.StatusCode != http.StatusOK {
		return nil, tkerr.Create(tkerr.OSAPIError)
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return parseHSResponse(string(contents))
}

// Parse the OSRS API response and return a slice of SkillInfo objects
// containing information about each of the player's skills.
func parseHSResponse(contents string) ([]SkillInfo, error) {
	skills := make([]SkillInfo, SkillCount)
	ss := strings.Split(string(contents), "\n")[:SkillCount]
	totallvl := 0
	i := SkillCount - 1

	// This is done backwards so that the total level can be calculated
	// and put into the entry for the Overall skill.
	for i >= 0 {
		fields := strings.Split(ss[i], ",")
		if len(fields) != 3 {
			return nil, tkerr.Create(tkerr.OSAPIError)
		}

		xp, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, tkerr.Create(tkerr.OSAPIError)
		}
		if xp == -1 {
			xp = 0
		}

		var lvl int
		if i == OverallID {
			lvl = totallvl
		} else {
			lvl = XPToLevel(xp)
			totallvl += lvl
		}

		rank, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, tkerr.Create(tkerr.OSAPIError)
		}

		skills[i] = SkillInfo{
			ID:         uint(i),
			Experience: xp,
			Level:      lvl,
			Rank:       rank,
			Hours:      0,
		}

		i--
	}

	return skills, nil
}
