package models

import "time"

// Team is a representation of a group of players for a single sport
type Team struct {
	Base

	// Sport is the "parent" sport that this team was created for
	// Sport Sport `json:"sport"`

	// PowerScore is the overall "ranking" of this team based on it's player's skills for the given sport, and the
	// weights that sport gives them
	PowerScore float32 `json:"power_score"`

	// Players is the list of players for this team
	Players []*Player `gorm:"many2many:player_teams;"`
	// Players []*Player
}

// TeamSet is the "output" of the team generation algorithm, which represents the collection of teams for
// a specific sport or season
type TeamSet struct {
	// User provided name to reference
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Sport     Sport     `json:"sport"`
	Teams     []*Team   `json:"teams"`
}

type TeamList struct {
	NumItems int     `json:"num_items"`
	Items    []*Team `json:"items"`
}

type GenerateTeamRequest struct {
	Name    string `json:"name"`
	SportID string `json:"sport_id"`

	Players       []*Player `json:"players"`
	NumberOfTeams int       `json:"number_of_teams"`
}
