package db

import (
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/pioz/faker"
	"github.com/segmentio/ksuid"
)

var (
	ultimateFrisbee = models.Sport{
		Base: models.Base{
			Name: "Ultimate Frisbee",
			ID:   ksuid.New().String(),
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
			Name: "Football",
			ID:   ksuid.New().String(),
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
			Name: "Basketball",
			ID:   ksuid.New().String(),
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
				Name: fmt.Sprintf("%s %s", fn, ln),
				ID:   ksuid.New().String(),
			},
			PowerScores: make(map[string]int),
			IsSeed:      true,
		}

		for _, s := range models.SkillsList {
			p.PowerScores[s] = faker.IntInRange(1, 10)
		}

		players[i] = p
	}

	return players
}
