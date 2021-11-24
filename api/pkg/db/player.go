package db

import (
	"errors"
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"gorm.io/gorm"
)

func InsertPlayer(db *gorm.DB, player *models.Player) error {
	return db.Create(player).Error
}

func GetPlayer(db *gorm.DB, id string) (*models.Player, error) {
	player := &models.Player{}
	if err := db.Where("id = ?", id).First(player).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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
	if err := db.Where("id = ?", id).First(&models.Player{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("player not found")
	} else {
		fmt.Println("ya")
		db.Where("id = ?", id).Delete(&models.Player{})
		return nil
	}
}
