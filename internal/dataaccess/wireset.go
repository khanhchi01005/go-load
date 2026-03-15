package dataaccess

import (
	"goload/internal/dataaccess/database"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	database.WireSet,
)