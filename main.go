package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// Connect to the PostgreSQL database
	connStr := goDotEnvVariable("CONNSTR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Perform a sample query
	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var nombre string
		var clave string
		if err := rows.Scan(&nombre, &clave); err != nil {
			panic(err)
		}
		fmt.Printf("Nombre: %s, Clave: %s\n", nombre, clave)
	}
}
