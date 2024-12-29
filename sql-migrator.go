

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var dbConfig = struct {
	User     string
	Password string
	Host     string
	Database string
}{
	User:     "root",
	Password: "password",
	Host:     "localhost",
	Database: "my_database",
}



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}
	dbConfig.User = os.Getenv("DB_USER")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Database = os.Getenv("DB_NAME")
	if dbConfig.User == "" || dbConfig.Password == "" || dbConfig.Host == "" || dbConfig.Database == "" {
		log.Fatal("Database configuration is missing in .env or environment variables.")
	}
	if len(os.Args) < 2 {
		log.Fatal("Please provide the migrations directory as an argument.")
	}

	migrationsDir := os.Args[1]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Database)
fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := initializeDbVersionTable(db); err != nil {
		log.Fatalf("Failed to initialize db_version table: %v", err)
	}

	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			if err := applyMigration(db, filepath.Join(migrationsDir, file.Name())); err != nil {
				log.Printf("Failed to apply migration %s: %v", file.Name(), err)
			}
		}
	}

	log.Println("All migrations executed successfully.")
}

func initializeDbVersionTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS db_version (
			id INT AUTO_INCREMENT PRIMARY KEY,
			migration_name VARCHAR(255) NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.Exec(query)
	return err
}

func hasMigrationBeenApplied(db *sql.DB, migrationName string) (bool, error) {
	query := `SELECT 1 FROM db_version WHERE migration_name = ? LIMIT 1`
	row := db.QueryRow(query, migrationName)
	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}

func applyMigration(db *sql.DB, migrationPath string) error {
	migrationName := filepath.Base(migrationPath)

	applied, err := hasMigrationBeenApplied(db, migrationName)
	if err != nil {
		return err
	}
	if applied {
		log.Printf("Skipping already applied migration: %s", migrationName)
		return nil
	}

	sqlBytes, err := ioutil.ReadFile(migrationPath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	if _, err := db.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	_, err = db.Exec(`INSERT INTO db_version (migration_name) VALUES (?)`, migrationName)
	if err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	log.Printf("Successfully applied migration: %s", migrationName)
	return nil
}



