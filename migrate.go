package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/joho/godotenv"
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

func main() {
	db := connect()

	stmt, err := db.Prepare("INSERT INTO proteins SET protein = ?, sequence = ?")
	if err != nil {
		log.Fatal(err)
	}

	proteins := []protein{

		/*
		Letter U
		 */

		{
			"Phe", "UUU",
		},
		{
			"Phe", "UUC",
		},
		{
			"Leu", "UUA",
		},
		{
			"Leu", "UUG",
		},
		{
			"Ser", "UCU",
		},
		{
			"Ser", "UCC",
		},
		{
			"Ser", "UCA",
		},
		{
			"Ser", "UCG",
		},
		{
			"Tyr", "UAU",
		},
		{
			"Tyr", "UAC",
		},
		{
			"STOP", "UAA",
		},
		{
			"STOP", "UAG",
		},
		{
			"Cys", "UGU",
		},
		{
			"Cys", "UGC",
		},
		{
			"STOP", "UGA",
		},
		{
			"Trp", "UGG",
		},


		/*
			Letter C
		*/


		{
			"Leu", "CUU",
		},
		{
			"Leu", "CUC",
		},
		{
			"Leu", "CUA",
		},
		{
			"Leu", "CUG",
		},
		{
			"Pro", "CCU",
		},
		{
			"Pro", "CCC",
		},
		{
			"Pro", "CCA",
		},
		{
			"Pro", "CCG",
		},
		{
			"His", "CAU",
		},
		{
			"His", "CAC",
		},
		{
			"Gin", "CAA",
		},
		{
			"Gin", "CAG",
		},
		{
			"Arg", "CGU",
		},
		{
			"Arg", "CGC",
		},
		{
			"Arg", "CGA",
		},
		{
			"Arg", "CGG",
		},


		/*
			Letter A
		*/


		{
			"Ile", "AUU",
		},
		{
			"Ile", "AUC",
		},
		{
			"Ile", "AUA",
		},
		{
			"Met", "AUG",
		},
		{
			"Thr", "ACU",
		},
		{
			"Thr", "ACC",
		},
		{
			"Thr", "ACA",
		},
		{
			"Thr", "ACG",
		},
		{
			"Asn", "AAU",
		},
		{
			"Asn", "AAC",
		},
		{
			"Lys", "AAA",
		},
		{
			"Lys", "AAG",
		},
		{
			"Ser", "AGU",
		},
		{
			"Ser", "AGC",
		},
		{
			"Arg", "AGA",
		},
		{
			"Arg", "AGG",
		},


		/*
			Letter G
		*/

		{
			"Val", "GUU",
		},
		{
			"Val", "GUC",
		},
		{
			"Val", "GUA",
		},
		{
			"Val", "GUG",
		},
		{
			"Ala", "GCU",
		},
		{
			"Ala", "GCC",
		},
		{
			"Ala", "GCA",
		},
		{
			"Ala", "GCG",
		},
		{
			"Asp", "GAU",
		},
		{
			"Asp", "GAC",
		},
		{
			"Glu", "GAA",
		},
		{
			"Glu", "GAG",
		},
		{
			"Gly", "GGU",
		},
		{
			"Gly", "GGC",
		},
		{
			"Gly", "GGA",
		},
		{
			"Gly", "GGG",
		},
	}

	for _, value := range proteins {
		e, err := stmt.Exec(value.Protein, value.Sequence)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e)
	}


}

type protein struct {
	Protein string
	Sequence string
}