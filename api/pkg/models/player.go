package models

type Player struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	// Skills is a map relating an existing Sport's ID to a SkillMap (map of skill name to rating out of 10)
	Skills      map[string]SkillMap `json:"skills"`
	PowerScores map[string]int      `json:"power_scores"`
	IsSeed      bool                `json:"-"`
}

type PlayerList struct {
	NumItems int      `json:"num_items"`
	Items    []Player `json:"items"`
}

type SkillMap map[string]int
