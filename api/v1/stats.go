package apiv1

import (
	"fmt"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"timekeeper/lib/cache"
	"timekeeper/lib/osrs"
	"timekeeper/tk"
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
	LastUpdate int64       `json:"lastUpdate"`
	Skills     []SkillInfo `json:"skills"`
}

func lookupStats(c *gin.Context) {
	username := c.Param("username")
	cacheKey := fmt.Sprintf("stats_%s", username)

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

	dp, _ := tk.LatestDatapoint(acc)
	res := StatsResponse{
		Username:   acc.DisplayName,
		LastUpdate: dp.CreatedAt.Unix(),
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

	c.Data(200, "application/json; charset=utf-8", j)
}
