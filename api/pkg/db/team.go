package db

import (
	"errors"
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"gorm.io/gorm"
)

func InsertTeam(db *gorm.DB, team *models.Team) error {
	return db.Create(team).Error
}

func GetTeam(db *gorm.DB, id string) (*models.Team, error) {
	team := &models.Team{}
	if err := db.Where("id = ?", id).First(team).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("team not found")
	}

	return team, nil
}

func ListTeams(db *gorm.DB) (*models.TeamList, error) {
	teamList := &models.TeamList{}
	db.Order("name").Find(&teamList.Items)

	// Fetch many to many relation
	for _, t := range teamList.Items {
		db.Model(&t).Association("players")
	}

	teamList.NumItems = len(teamList.Items)
	return teamList, nil
}

func UpdateTeam(db *gorm.DB, team *models.Team) (*models.Team, error) {
	if err := db.Save(team).Error; err != nil {
		return nil, err
	} else {
		s := &models.Team{}
		db.Where("id = ?", team.ID).First(team)
		return s, nil
	}
}

func DeleteTeam(db *gorm.DB, id string) error {
	if err := db.Where("id = ?", id).First(&models.Team{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("team not found")
	} else {
		fmt.Println("ya")
		db.Where("id = ?", id).Delete(&models.Team{})
		return nil
	}
}
