package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SkillWeightMap map[string]float32

type Sport struct {
	gorm.Model
	Name              string         `json:"name"`
	ID                string         `json:"id"`
	SkillWeights      SkillWeightMap `json:"skills"`
	MaxPlayersPerTeam int            `json:"max_players_per_team"`
}

type SportList struct {
	NumItems int     `json:"num_items"`
	Items    []Sport `json:"items"`
}

var (
	ultimateFrisbee = Sport{
		Name: "Ultimate Frisbee",
		ID:   "bqrc556hds3g8muin1ag", //xid.New().String(),
		SkillWeights: SkillWeightMap{
			"handling": 0.9,
			"speed":    0.8,
			"stamina":  0.8,
			"height":   0.4,
		},
	}

	football = Sport{
		Name: "Football",
		ID:   "bqrc7tmhds3g8muin1b0", //xid.New().String(),
		SkillWeights: SkillWeightMap{
			"strength": 0.7,
			"speed":    0.8,
			"stamina":  0.8,
			"agility":  0.5,
		},
	}

	basketball = Sport{
		Name: "Basketball",
		ID:   "bqrc7tmhds3g8muin1bg", //xid.New().String(),
		SkillWeights: SkillWeightMap{
			"shooting": 0.9,
			"speed":    0.6,
			"stamina":  0.8,
			"height":   0.8,
			"passing":  0.5,
		},
	}
)

var InitialSports = []Sport{
	ultimateFrisbee,
	basketball,
	football,
}
