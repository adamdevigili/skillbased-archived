package models

import (
	"fmt"

	"github.com/pioz/faker"
	"github.com/rs/xid"
)

var (
	ultimateFrisbee = Sport{
		Base: Base{
			Name: "Ultimate Frisbee",
			ID:   xid.New(),
		},
		SkillWeights: SkillWeightMap{
			"handling": 0.9,
			"speed":    0.8,
			"stamina":  0.8,
			"height":   0.4,
		},
	}

	football = Sport{
		Base: Base{
			Name: "Football",
			ID:   xid.New(),
		},
		SkillWeights: SkillWeightMap{
			"strength": 0.7,
			"speed":    0.8,
			"stamina":  0.8,
			"agility":  0.5,
		},
	}

	basketball = Sport{
		Base: Base{
			Name: "Basketball",
			ID:   xid.New(),
		},
		SkillWeights: SkillWeightMap{
			"shooting": 0.9,
			"speed":    0.6,
			"stamina":  0.8,
			"height":   0.8,
			"passing":  0.5,
		},
	}

	InitialSports = []Sport{
		ultimateFrisbee,
		basketball,
		football,
	}
)

var playersToGenerate = 30

func GenerateSeedPlayers() []*Player {
	faker.SetSeed(623)
	players := make([]*Player, playersToGenerate)

	for i := range players {
		fn, ln := faker.FirstName(), faker.LastName()
		p := &Player{
			FirstName: fn,
			LastName:  ln,
			Base: Base{
				Name: fmt.Sprintf("%s %s", fn, ln),
				ID:   xid.New(),
			},
			PowerScores: make(map[string]int),
			IsSeed:      true,
		}

		for _, s := range skillsList {
			p.PowerScores[s] = faker.IntInRange(1, 10)
		}

		players[i] = p
	}

	return players
}
