package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func connect() *sql.DB {
	env := godotenv.Load(".env.local")

	if env != nil {
		log.Fatal(env)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))
	if err != nil {
		log.Panic(err)
	}

	return db
}

func GetProtein(mRNA string, w http.ResponseWriter, r *http.Request) string {
	var protein string
	db := connect()
	err := db.QueryRow("select protein from proteins where sequence = ?", mRNA).Scan(&protein)
	if err != nil {
		log.Panic(err)
	}

	db.Close()

	return protein
}