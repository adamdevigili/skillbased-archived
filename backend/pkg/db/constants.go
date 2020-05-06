package db

import (
	"github.com/adamdevigili/skillbased.io/pkg/models"
)

// Temporary in-memory data store
var (
	TeamsMem   = map[string]*models.Team{}
	SportsMem  = map[string]*models.Sport{}
	PlayersMem = map[string]*models.Player{}
)
