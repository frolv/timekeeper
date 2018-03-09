package apiv1

import (
	"fmt"
	"time"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"timekeeper/lib/cache"
	"timekeeper/lib/osrs"
	"timekeeper/models"
)

type SkillInfo struct {
	SkillID    uint    `json:"id"`
	SkillName  string  `json:"name"`
	Level      int     `json:"level"`
	Experience int     `json:"experience"`
	Rank       int     `json:"rank"`
	Hours      float64 `json:"hours"`
}

type StatsResponse struct {
	Username   string      `json:"username"`
	LastUpdate time.Time   `json:"lastUpdate"`
	Skills     []SkillInfo `json:"skills"`
}

func lookupStats(c *gin.Context) {
	username := c.Param("username")
	cacheKey := fmt.Sprintf("stats_%s", username)

	val, err := cache.Get(cacheKey)
	if err == nil {
		c.String(200, val)
		return
	}

	acc, err := models.GetAccount(username)
	if err != nil {
		c.JSON(404, gin.H{"status": "error", "errorMessage": err.Error()})
		return
	}

	dp, _ := models.GetLatestDatapoint(acc)
	res := StatsResponse{
		Username:   acc.Username,
		LastUpdate: dp.CreatedAt,
		Skills:     make([]SkillInfo, osrs.SkillCount),
	}

	for i, sk := range dp.SkillLevels {
		res.Skills[i] = SkillInfo{
			SkillID:    sk.SkillID,
			SkillName:  osrs.Skills[sk.SkillID].Name,
			Level:      sk.Level,
			Experience: sk.Experience,
			Rank:       sk.Rank,
			Hours:      sk.Hours,
		}
	}

	j, _ := json.Marshal(res)
	s := string(j)
	cache.Set(cacheKey, s, acc)

	c.String(200, s)
}
