package core

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sort"
)

const (
	numberOfTeams  = 3
	speedWeight    = 0.9
	handlingWeight = 0.8
	staminaWeight  = 0.8
	heightWeight   = 0.4
)

type Player struct {
	PowerScore float32
	Name       string `json:"name"`

	Handling int `json:"handling"`
	Speed    int `json:"speed"`
	Stamina  int `json:"stamina"`
	Height   int `json:"height"`
}

type Players struct {
	Players []Player `json:"players"`
}

type Team struct {
	Name       string
	PowerScore float32
	Players    []Player
}

func main() {
	jsonFile, err := os.Open("players.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var players Players
	json.Unmarshal(byteValue, &players)

	makeTeams(players.Players, numberOfTeams)

	// for i, team := range madeTeams {
	// 	fmt.Println(fmt.Sprintf("Team %d: %+v", i+1, team))
	// }

}

func makeTeams(players []Player, numTeams int) {
	teams := make([]Team, numTeams)
	for i := range teams {
		teams[i].Name = fmt.Sprintf("Team %d", i)
	}

	// Calculate power score for every player
	for i := range players {
		players[i].PowerScore = calcPowerScore(players[i])
	}

	// Sort players by power score (descending)
	sort.Slice(players, func(i, j int) bool {
		return players[i].PowerScore > players[j].PowerScore
	})

	// Iterate over sorted list, adding players to different teams
	for i, player := range players {
		if i == len(players)-1 {
			teams[numTeams-1].Players = append(teams[numTeams-1].Players, player)
			teams[numTeams-1].PowerScore += player.PowerScore
		} else {
			teams[i%numTeams].Players = append(teams[i%numTeams].Players, player)
			teams[i%numTeams].PowerScore += player.PowerScore
		}
	}

	// Print teams
	for i, team := range teams {
		fmt.Println(fmt.Sprintf("--- Team %d, Size %d, PowerScore %0.2f ---", i+1, len(team.Players), team.PowerScore))
		for _, player := range team.Players {
			fmt.Println(fmt.Sprintf("%s - %0.2f", player.Name, player.PowerScore))
		}
		fmt.Println("\n")
	}
}

func calcPowerScore(player Player) float32 {
	var powerScore float32

	powerScore += float32(player.Speed) * speedWeight
	powerScore += float32(player.Handling) * handlingWeight
	powerScore += float32(player.Stamina) * staminaWeight
	powerScore += float32(player.Height) * heightWeight

	return powerScore
}

