package adapter

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var connectionDB = make(map[string]interface{})

func LoadMultipleDatabase() {
	// slice list db into array string
	listDb := viper.GetString("userservice.db.lists")
	db := strings.Split(listDb, ",")

	// slice db name connection into array string
	listDbName := viper.GetString("userservice.db.name")
	dbName := strings.Split(listDbName, ",")

	if len(db) > 0 {
		for index, database := range db {
			connection, err := Database(dbName[index], database)
			if err != nil {
				panic(err)
			}
			connectionDB[dbName[index]] = connection
		}
	}
}

func LoadMigrationsDatabase() {
	// connection strings
	connstrings := viper.GetString("userservice.db.migrations.connstrings")

	// folder migration sql
	folder := viper.GetString("userservice.db.migrations.folder")

	m, err := migrate.New(
		folder,
		connstrings)

	if err != nil {
		log.Fatalf("Unable to instantiate the database schema migrator - %v", err)
	}

	// Migrate up to the latest active version
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Unable to migrate up to the latest database schema - %v", err)
	}
}

func Database(dbName, dsn string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		err = fmt.Errorf("error to connection db %s. %v", dbName, err)
		panic(err)
	}

	maxIdle := viper.GetInt("db.maxidle")
	maxConn := viper.GetInt("db.maxconn")
	// set max idle and connection to database
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxConn)

	err = db.Ping()

	return db, err
}

func DBConnection(dbName string) *sqlx.DB {
	return connectionDB[dbName].(*sqlx.DB)
}
