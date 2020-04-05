package models

type SkillWeightMap map[string]float32

type Sport struct {
	Name              string         `json:"name"`
	ID                string         `json:"id"`
	SkillWeights      SkillWeightMap `json:"skills"`
	MaxPlayersPerTeam int            `json:"max_players_per_team"`
}

var (
	UltimateFrisbee = Sport{
		Name: "Ultimate Frisbee",
		ID:   "ultimate-id", //xid.New().String(),
		SkillWeights: SkillWeightMap{
			"handling": 0.9,
			"speed":    0.8,
			"stamina":  0.8,
			"height":   0.4,
		},
	}
)
