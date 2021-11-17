package models

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type SkillWeightMap map[string]float32

type Sport struct {
	Base
	SkillWeights      SkillWeightMap `gorm:"-" json:"skills"` // Ignore this field for DB
	SkillWeightsDB    postgres.Jsonb `json:"-"`               // Ignore this field for JSON
	MaxPlayersPerTeam int            `json:"max_players_per_team"`
}

// BeforeCreate is a GORM hook that is used to convert the Go map to a JSON struct to be stored in postgres
func (s *Sport) BeforeCreate(scope *gorm.Scope) (err error) {
	if x, err := json.Marshal(s.SkillWeights); err != nil {
		return err
	} else {
		s.SkillWeightsDB = postgres.Jsonb{RawMessage: x}
	}

	return
}

//func (s *Sport) BeforeUpdate() (err error) {
//	if x, err := json.Marshal(s.SkillWeights); err != nil {
//		return err
//	} else {
//		s.SkillWeightsDB = postgres.Jsonb{RawMessage: x}
//	}
//
//	return
//}

// AfterFind is a GORM hook that is used to convert the JSON struct in postgres to the Go map to be returned to the user
// and operated on by the balancer algorithm
func (s *Sport) AfterFind() (err error) {
	fmt.Println("AfterFind called on " + s.ID.String())
	if err := json.Unmarshal(s.SkillWeightsDB.RawMessage, &s.SkillWeights); err != nil {
		return err
	}

	return
}

type SportList struct {
	NumItems int     `json:"num_items"`
	Items    []Sport `json:"items"`
}
