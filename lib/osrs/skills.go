package osrs

import "math"

// Definitions of OSRS skills.
const (
	OverallID = iota
	AttackID
	DefenceID
	StrengthID
	HitpointsID
	RangedID
	PrayerID
	MagicID
	CookingID
	WoodcuttingID
	FletchingID
	FishingID
	FiremakingID
	CraftingID
	SmithingID
	MiningID
	HerbloreID
	AgilityID
	ThievingID
	SlayerID
	FarmingID
	RunecraftID
	HunterID
	ConstructionID
)

const SkillCount = 24

type Skill struct {
	ID   uint
	Name string
}

var Skills []Skill

func init() {
	Skills = []Skill{
		{OverallID, "Overall"},
		{AttackID, "Attack"},
		{DefenceID, "Defence"},
		{StrengthID, "Strength"},
		{HitpointsID, "Hitpoints"},
		{RangedID, "Ranged"},
		{PrayerID, "Prayer"},
		{MagicID, "Magic"},
		{CookingID, "Cooking"},
		{WoodcuttingID, "Woodcutting"},
		{FletchingID, "Fletching"},
		{FishingID, "Fishing"},
		{FiremakingID, "Firemaking"},
		{CraftingID, "Crafting"},
		{SmithingID, "Smithing"},
		{MiningID, "Mining"},
		{HerbloreID, "Herblore"},
		{AgilityID, "Agility"},
		{ThievingID, "Thieving"},
		{SlayerID, "Slayer"},
		{FarmingID, "Farming"},
		{RunecraftID, "Runecraft"},
		{HunterID, "Hunter"},
		{ConstructionID, "Construction"},
	}
}

// Return the virtual level for the given amount of experience points.
func XPToLevel(xp int) int {
	n := 1.0
	xp = xp*4 + 1
	for xp >= 0 {
		xp -= int(math.Floor(n + 300*math.Pow(2, n/7.0)))
		n++
	}

	return int(n) - 1
}
