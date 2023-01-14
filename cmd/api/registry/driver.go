package registry

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"os"
)

func NewDatabaseDriver() database.Driver {
	return database.NewPostgresDriver(NewDatabaseHandler())
}

func NewDatabaseHandler() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logrus.Info(psqlInfo)
	logrus.Info("Successfully connected!")

	return db
}
