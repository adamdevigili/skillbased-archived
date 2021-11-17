package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Player struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	// Skills is a map relating an existing Sport's ID to a SkillMap (map of skill name to rating out of 10)
	Skills   map[string]SkillMap `gorm:"-" json:"skills"`
	SkillsDB postgres.Jsonb      `json:"-"`

	PowerScores   map[string]int `gorm:"-" json:"power_scores"`
	PowerScoresDB postgres.Jsonb `json:"-"`

	IsSeed bool `json:"-"`
}

// BeforeCreate is a GORM hook that is used to convert the Go map to a JSON struct to be stored in postgres
func (p *Player) BeforeCreate(scope *gorm.Scope) (err error) {
	if x, err := json.Marshal(p.PowerScores); err != nil {
		return err
	} else {
		p.PowerScoresDB = postgres.Jsonb{RawMessage: x}
	}

	return
}

type PlayerList struct {
	NumItems int      `json:"num_items"`
	Items    []Player `json:"items"`
}

type SkillMap map[string]int
