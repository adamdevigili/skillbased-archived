package db

import (
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func InsertPlayer(db *gorm.DB, player *models.Player) error {
	return db.Create(player).Error
}

func GetPlayer(db *gorm.DB, id string) (*models.Player, error) {
	player := &models.Player{}
	if db.Where("id = ?", id).First(player).RecordNotFound() {
		return nil, fmt.Errorf("player not found")
	}

	return player, nil
}

func ListPlayers(db *gorm.DB) (*models.PlayerList, error) {
	playerList := &models.PlayerList{}
	db.Order("name").Find(&playerList.Items)

	playerList.NumItems = len(playerList.Items)
	return playerList, nil
}

func UpdatePlayer(db *gorm.DB, player *models.Player) (*models.Player, error) {
	if err := db.Save(player).Error; err != nil {
		return nil, err
	} else {
		s := &models.Player{}
		db.Where("id = ?", player.ID).First(player)
		return s, nil
	}
}

func DeletePlayer(db *gorm.DB, id string) error {
	if db.Where("id = ?", id).First(&models.Player{}).RecordNotFound() {
		return fmt.Errorf("player not found")
	} else {
		fmt.Println("ya")
		db.Where("id = ?", id).Delete(&models.Player{})
		return nil
	}
}

func deleteAllSeedPlayers(db *gorm.DB) {
	log.Info("deleteing all seed players from existing DB")
	db.Where("is_seed = ?", "t").Delete(&models.Player{})
}
