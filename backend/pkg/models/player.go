package models

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// SkillWeights is a map relating an existing Sport's ID to a SkillMap (map of skill name to rating out of 10)
	Skills      map[string]SkillMap `json:"skills"`
	PowerScores map[string]float32  `json:"power_scores"`
}

type SkillMap map[string]int
