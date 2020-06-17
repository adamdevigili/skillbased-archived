package core

import (
	"fmt"
	"sort"
	"time"

	"github.com/Pallinder/go-randomdata"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"
)

func GenerateTeams(req models.GenerateTeamRequest, sport models.Sport) models.TeamSet {
	numberOfTeams := req.NumberOfTeams
	players := req.Players

	teams := make([]models.Team, numberOfTeams)

	// Calculate power score for every player provided. If an ID was provided, check if that player has existing stats
	// for the target sport. Otherwise, generate for that player and update their entry.
	for _, player := range players {
		if player.ID == "" {
			player.PowerScores = map[string]float32{}
			player.PowerScores[sport.ID.String()] = calcPlayerPowerScoreForSport(player, sport)
			player.ID = xid.New().String()
			db.PlayersMem[player.ID] = &player
			log.Info(fmt.Sprintf("%s player score: %f", player.Name, player.PowerScores[sport.ID.String()]))
		} else {
			player, ok := db.PlayersMem[player.ID]
			if !ok {
				log.Warn(fmt.Sprintf("provided ID %s not present", player.ID))
			} else {
				if _, ok := player.PowerScores[sport.ID.String()]; !ok {
					player.PowerScores[sport.ID.String()] = calcPlayerPowerScoreForSport(*player, sport)
				} else {
					log.Warn(fmt.Sprintf("provided ID %s doesn't have stats for target sport", player.ID))
				}
			}
		}
	}

	// Sort all players by power score (descending)
	sort.Slice(req.Players, func(i, j int) bool {
		return req.Players[i].PowerScores[sport.ID.String()] > req.Players[j].PowerScores[sport.ID.String()]
	})

	// Iterate over sorted list, adding players to different teams
	for i, player := range players {
		if i == len(players)-1 {
			teams[numberOfTeams-1].Players = append(teams[numberOfTeams-1].Players, player)
			teams[numberOfTeams-1].PowerScore += player.PowerScores[sport.ID.String()]
		} else {
			teams[i%numberOfTeams].Players = append(teams[i%numberOfTeams].Players, player)
			teams[i%numberOfTeams].PowerScore += player.PowerScores[sport.ID.String()]
		}
	}

	for i := range teams {
		teams[i].Sport = sport
		teams[i].ID = xid.New().String()
		teams[i].Name = randomdata.SillyName()
		teams[i].CreatedAt = time.Now()
		db.TeamsMem[teams[i].ID] = &teams[i]
	}

	return models.TeamSet{
		Name:      req.Name,
		ID:        xid.New().String(),
		Teams:     teams,
		Sport:     sport,
		CreatedAt: time.Now(),
	}
}

func calcPlayerPowerScoreForSport(player models.Player, sport models.Sport) float32 {
	var powerScore float32

	for skill, weight := range sport.SkillWeights {
		powerScore += float32(player.Skills[sport.ID.String()][skill]) * weight
	}

	return powerScore
}
