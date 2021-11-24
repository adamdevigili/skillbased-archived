package db

import (
	"errors"
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"gorm.io/gorm"
)

func InsertSport(db *gorm.DB, sport *models.Sport) error {
	return db.Create(sport).Error
}

func GetSport(db *gorm.DB, id string) (*models.Sport, error) {
	sport := &models.Sport{}
	if err := db.Where("id = ?", id).First(sport).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("sport not found")
	}

	return sport, nil
}

func ListSports(db *gorm.DB) (*models.SportList, error) {
	sportList := &models.SportList{}
	db.Order("name").Find(&sportList.Items)

	sportList.NumItems = len(sportList.Items)
	return sportList, nil
}

func UpdateSport(db *gorm.DB, sport *models.Sport) (*models.Sport, error) {
	if err := db.Save(sport).Error; err != nil {
		return nil, err
	} else {
		s := &models.Sport{}
		db.Where("id = ?", sport.ID).First(sport)
		return s, nil
	}
}

func DeleteSport(db *gorm.DB, id string) error {
	if err := db.Where("id = ?", id).First(&models.Sport{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("sport not found")
	} else {
		fmt.Println("ya")
		db.Where("id = ?", id).Delete(&models.Sport{})
		return nil
	}
}
