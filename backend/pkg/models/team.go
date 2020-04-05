package models

type Team struct {
	Name       string   `json:"name"`
	ID         string   `json:"id"`
	Sport      Sport    `json:"sport"`
	PowerScore float32  `json:"power_score"`
	Players    []Player `json:"players"`
}

type GenerateTeamRequest struct {
	SportID       string   `json:"sport_id"`
	Players       []Player `json:"players"`
	NumberOfTeams int      `json:"number_of_teams"`
}
