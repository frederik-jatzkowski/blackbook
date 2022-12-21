package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	// read config from env
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgDb := os.Getenv("POSTGRES_DB")

	// try to open connection
	db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable",
				pgUser,
				pgPassword,
				pgDb,
				pgPort,
			),
		),
		&gorm.Config{},
	)
	if err != nil {
		return db, err
	}

	// try auto migration
	err = db.AutoMigrate(&Group{})
	if err != nil {
		return db, fmt.Errorf("error while automigrating groups: %s", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return db, fmt.Errorf("error while automigrating users: %s", err)
	}

	err = db.AutoMigrate(&Invitation{})
	if err != nil {
		return db, fmt.Errorf("error while automigrating invitations: %s", err)
	}

	return db, err

}
