package apiv1

import (
	"fmt"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"timekeeper/lib/cache"
	"timekeeper/lib/osrs"
	"timekeeper/tk"
)

type TrackInfo struct {
	SkillID    uint    `json:"id"`
	SkillName  string  `json:"name"`
	DeltaExp   int     `json:"deltaExperience"`
	DeltaRank  int     `json:"deltaRank"`
	DeltaHours float64 `json:"deltaHours"`
}

type TrackResponse struct {
	Username    string      `json:"username"`
	PeriodStart int64       `json:"periodStart"`
	PeriodEnd   int64       `json:"periodEnd"`
	Skills      []TrackInfo `json:"skills"`
}

func trackAccount(c *gin.Context) {
	username := c.Param("username")
	cacheKey := fmt.Sprintf("track_%s", username)

	val, err := cache.Get(cacheKey)
	if err == nil {
		c.Data(200, "application/json; charset=utf-8", []byte(val))
		return
	}

	acc, err := tk.GetAccount(username)
	if err != nil {
		code, json := errorToResponse(err)
		c.JSON(code, json)
		return
	}

	period, err := tk.ParsePeriod(c.DefaultQuery("period", "7d"))
	if err != nil {
		code, json := errorToResponse(err)
		c.JSON(code, json)
		return
	}

	start := time.Now().Add(-period)

	dps, err := tk.BoundaryPoints(acc, start)
	if err != nil {
		code, json := errorToResponse(err)
		c.JSON(code, json)
		return
	}

	res := TrackResponse{
		Username:    acc.Username,
		PeriodStart: dps[0].CreatedAt.Unix(),
		PeriodEnd:   dps[1].CreatedAt.Unix(),
		Skills:      make([]TrackInfo, osrs.SkillCount),
	}

	for i, _ := range dps[0].SkillLevels {
		sk0 := &dps[0].SkillLevels[i]
		sk1 := &dps[1].SkillLevels[i]
		res.Skills[i] = TrackInfo{
			SkillID:    sk0.SkillID,
			SkillName:  osrs.Skills[sk0.SkillID].Name,
			DeltaExp:   sk1.Experience - sk0.Experience,
			DeltaRank:  -(sk1.Rank - sk0.Rank),
			DeltaHours: sk1.Hours - sk0.Hours,
		}
	}

	j, _ := json.Marshal(res)
	s := string(j)

	// If the first datapoint in the period expires soon,
	// make sure that the cache entry expires with it.
	tte := dps[0].CreatedAt.Sub(start)
	if (tte < cache.DefaultExpiration) {
		cache.SetExpiration(cacheKey, s, tte, acc)
	} else {
		cache.Set(cacheKey, s, acc)
	}

	c.Data(200, "application/json; charset=utf-8", j)
}
