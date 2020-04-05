package db

import (
	"github.com/adamdevigili/balancer.team/pkg/models"
)

var (
	TeamsMem   = map[string]*models.Team{}
	SportsMem  = map[string]*models.Sport{"ultimate-id": &models.UltimateFrisbee}
	PlayersMem = map[string]*models.Player{}
)
