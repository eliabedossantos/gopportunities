package config

import (
	"os"

	"github.com/eliabedossantos/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitialiizeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	//check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		//create the database file
		logger.Info("database file not found, creating...")

		//create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("database directory creation error: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("database file creation error: %v", err)
			return nil, err
		}
		file.Close()
		logger.Info("database file created")

	}

	//create db connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	//migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite auto migrate error: %v", err)
		return nil, err
	}

	//return DB
	return db, nil
}
