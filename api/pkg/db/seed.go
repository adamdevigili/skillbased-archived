package db

import (
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/labstack/gommon/log"
	"github.com/pioz/faker"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

var (
	ultimateFrisbee = models.Sport{
		Base: models.Base{
			Name:   "Ultimate Frisbee",
			ID:     ksuid.New().String(),
			IsSeed: true,
		},
		SkillWeights: models.SkillWeightMap{
			"handling": 0.9,
			"speed":    0.8,
			"stamina":  0.8,
			"height":   0.4,
		},
		MaxPlayersPerTeam:       15,
		MaxActivePlayersPerTeam: 7,
	}

	football = models.Sport{
		Base: models.Base{
			Name:   "Football",
			ID:     ksuid.New().String(),
			IsSeed: true,
		},
		SkillWeights: models.SkillWeightMap{
			"strength": 0.7,
			"speed":    0.8,
			"stamina":  0.8,
			"agility":  0.5,
		},
		MaxPlayersPerTeam:       50,
		MaxActivePlayersPerTeam: 11,
	}

	basketball = models.Sport{
		Base: models.Base{
			Name:   "Basketball",
			ID:     ksuid.New().String(),
			IsSeed: true,
		},
		SkillWeights: models.SkillWeightMap{
			"shooting": 0.9,
			"speed":    0.6,
			"stamina":  0.8,
			"height":   0.8,
			"passing":  0.5,
		},
		MaxPlayersPerTeam:       12,
		MaxActivePlayersPerTeam: 5,
	}

	initialSports = []models.Sport{
		ultimateFrisbee,
		basketball,
		football,
	}
)

var playersToGenerate = 30

func generateSeedPlayers() []*models.Player {
	faker.SetSeed(623)
	players := make([]*models.Player, playersToGenerate)

	for i := range players {
		fn, ln := faker.FirstName(), faker.LastName()
		p := &models.Player{
			FirstName: fn,
			LastName:  ln,
			Base: models.Base{
				Name:   fmt.Sprintf("%s %s", fn, ln),
				ID:     ksuid.New().String(),
				IsSeed: true,
			},
			PowerScores: make(map[string]int),
		}

		for _, s := range models.SkillsList {
			p.PowerScores[s] = faker.IntInRange(1, 10)
		}

		players[i] = p
	}

	return players
}

func generateSeedTeams(seedPlayers []*models.Player, db *gorm.DB) []*models.Team {
	teamSize := 5
	numTeams := len(seedPlayers) / teamSize
	teams := make([]*models.Team, numTeams)

	log.Infof("%+v", teams, teamSize, numTeams, len(seedPlayers))
	currPlayer := 0
	for i := range teams {
		t := &models.Team{
			Base: models.Base{
				Name:   faker.ColorName(),
				ID:     ksuid.New().String(),
				IsSeed: true,
			},
		}

		for i := currPlayer; i < currPlayer+i; i++ {
			// t.Players = append(t.Players, seedPlayers[i])
			db.Model(t).Association("Players").Append()
			db.Model(seedPlayers[i]).Association("Teams").Append(seedPlayers[i])
		}

		log.Infof("generated team: %+v", t)

		teams[i] = t
	}

	return teams
}
