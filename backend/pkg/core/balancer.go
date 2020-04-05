package core

import (
	"fmt"
	"sort"

	"github.com/adamdevigili/balancer.team/pkg/db"
	"github.com/adamdevigili/balancer.team/pkg/models"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"
)

func GenerateTeams(players []models.Player, sport models.Sport, numberOfTeams int) []models.Team {
	teams := make([]models.Team, numberOfTeams)

	// Calculate power score for every player provided. If an ID was provided, check if that player has existing stats
	// for the target sport. Otherwise, generate for that player and update their entry.
	for i := range players {
		if players[i].ID == "" {
			players[i].PowerScores = map[string]float32{}
			players[i].PowerScores[sport.ID] = calcPlayerPowerScoreForSport(players[i], sport)
			players[i].ID = xid.New().String()
			db.PlayersMem[players[i].ID] = &players[i]
			log.Info(fmt.Sprintf("%s player score: %f", players[i].Name, players[i].PowerScores[sport.ID]))
		} else {
			player, ok := db.PlayersMem[players[i].ID]
			if !ok {
				log.Warn(fmt.Sprintf("provided ID %s not present", players[i].ID))
			} else {
				if _, ok := player.PowerScores[sport.ID]; !ok {
					players[i].PowerScores[sport.ID] = calcPlayerPowerScoreForSport(players[i], sport)
				} else {
					log.Warn(fmt.Sprintf("provided ID %s doesn't have stats for target sport", players[i].ID))
				}
			}
		}
	}

	// Sort all players by power score (descending)
	sort.Slice(players, func(i, j int) bool {
		return players[i].PowerScores[sport.ID] > players[j].PowerScores[sport.ID]
	})

	// Iterate over sorted list, adding players to different teams
	for i, player := range players {
		if i == len(players)-1 {
			teams[numberOfTeams-1].Players = append(teams[numberOfTeams-1].Players, player)
			teams[numberOfTeams-1].PowerScore += player.PowerScores[sport.ID]
		} else {
			teams[i%numberOfTeams].Players = append(teams[i%numberOfTeams].Players, player)
			teams[i%numberOfTeams].PowerScore += player.PowerScores[sport.ID]
		}
	}

	return teams
}

func calcPlayerPowerScoreForSport(player models.Player, sport models.Sport) float32 {
	var powerScore float32

	for skill, weight := range sport.SkillWeights {
		powerScore += float32(player.Skills[sport.ID][skill]) * weight
	}

	return powerScore
}
