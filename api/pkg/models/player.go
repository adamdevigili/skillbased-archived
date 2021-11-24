package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
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

	// TeamID string

	Teams []*Team `gorm:"many2many:player_teams;"`
	// Groups []*PlayerGroup `gorm:"many2many:player_groups;"`
}

// BeforeCreate is a GORM hook that is used to convert the Go map to a JSON struct to be stored in postgres
func (p *Player) BeforeCreate(db *gorm.DB) (err error) {
	if x, err := json.Marshal(p.PowerScores); err != nil {
		return err
	} else {
		p.PowerScoresDB = postgres.Jsonb{RawMessage: x}
	}

	return
}

// PlayerList is the wrapper object given back to the user
type PlayerList struct {
	NumItems int       `json:"num_items"`
	Items    []*Player `json:"items"`
}

// PlayerGroup is a collection of players, referenced by ID
type PlayerGroup struct {
	Base
	Players []*Player

	IsSeed bool `json:"-"`
}

type SkillMap map[string]int
