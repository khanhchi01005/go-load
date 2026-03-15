package database

import(
	"github.com/google/wire"
	migrations "goload/internal/dataaccess/database/migrations"
)

var WireSet = wire.NewSet(
	migrations.InitializeDB,
	migrations.InitializeGoquDB,
	NewAccountDataAccessor,
	NewAccountPasswordDataAccessor,
	
)