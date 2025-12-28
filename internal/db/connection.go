package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/axolotl-go/turso-chat/internal/config"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	cfg := config.Load()

	tursoURL := cfg.DBUrl
	authToken := cfg.DBToken

	dsn := fmt.Sprintf("%s?authToken=%s", tursoURL, authToken)

	sqlDB, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatal("Error to open connection:", err)
	}

	DB, err = gorm.Open(sqlite.Dialector{
		Conn: sqlDB,
	}, &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connect:", err)
	}

	fmt.Println("Connected")
}
