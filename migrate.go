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

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `proteins` (`id` int(11) NOT NULL auto_increment, `protein` varchar(255) NOT NULL, `name` varchar(255)  NOT NULL, `sequence`  varchar(255) NOT NULL, PRIMARY KEY (id));")

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO proteins SET protein = ?, sequence = ?, name = ?")
	if err != nil {
		log.Fatal(err)
	}

	proteins := []protein{

		/*
		Letter U
		 */

		{
			"Phe", "UUU", "Phenylalanine",
		},
		{
			"Phe", "UUC", "Phenylalanine",
		},
		{
			"Leu", "UUA", "Leucine",
		},
		{
			"Leu", "UUG", "Leucine",
		},
		{
			"Ser", "UCU", "Serine",
		},
		{
			"Ser", "UCC", "Serine",
		},
		{
			"Ser", "UCA", "Serine",
		},
		{
			"Ser", "UCG", "Serine",
		},
		{
			"Tyr", "UAU", "Tyrosine",
		},
		{
			"Tyr", "UAC", "Tyrosine",
		},
		{
			"STOP", "UAA", "Ochre (STOP)",
		},
		{
			"STOP", "UAG", "Amber (STOP)",
		},
		{
			"Cys", "UGU", "Cysteine",
		},
		{
			"Cys", "UGC", "Cysteine",
		},
		{
			"STOP", "UGA", "Opal (STOP)",
		},
		{
			"Trp", "UGG", "Tryptophan",
		},


		/*
			Letter C
		*/


		{
			"Leu", "CUU", "Leucine",
		},
		{
			"Leu", "CUC", "Leucine",
		},
		{
			"Leu", "CUA", "Leucine",
		},
		{
			"Leu", "CUG", "Leucine",
		},
		{
			"Pro", "CCU", "Proline",
		},
		{
			"Pro", "CCC", "Proline",
		},
		{
			"Pro", "CCA", "Proline",
		},
		{
			"Pro", "CCG", "Proline",
		},
		{
			"His", "CAU", "Histidine",
		},
		{
			"His", "CAC", "Histidine",
		},
		{
			"Gln", "CAA", "Glutamine",
		},
		{
			"Gln", "CAG", "Glutamine",
		},
		{
			"Arg", "CGU", "Arginine",
		},
		{
			"Arg", "CGC", "Arginine",
		},
		{
			"Arg", "CGA", "Arginine",
		},
		{
			"Arg", "CGG", "Arginine",
		},


		/*
			Letter A
		*/


		{
			"Ile", "AUU", "Isoleucine",
		},
		{
			"Ile", "AUC", "Isoleucine",
		},
		{
			"Ile", "AUA", "Isoleucine",
		},
		{
			"Met", "AUG", "Methionine",
		},
		{
			"Thr", "ACU", "Threonine",
		},
		{
			"Thr", "ACC", "Threonine",
		},
		{
			"Thr", "ACA", "Threonine",
		},
		{
			"Thr", "ACG", "Threonine",
		},
		{
			"Asn", "AAU", "Asparagine",
		},
		{
			"Asn", "AAC", "Asparagine",
		},
		{
			"Lys", "AAA", "Lysine",
		},
		{
			"Lys", "AAG", "Lysine",
		},
		{
			"Ser", "AGU", "Serine",
		},
		{
			"Ser", "AGC", "Serine",
		},
		{
			"Arg", "AGA", "Arginine",
		},
		{
			"Arg", "AGG", "Arginine",
		},


		/*
			Letter G
		*/

		{
			"Val", "GUU", "Valine",
		},
		{
			"Val", "GUC", "Valine",
		},
		{
			"Val", "GUA", "Valine",
		},
		{
			"Val", "GUG", "Valine",
		},
		{
			"Ala", "GCU", "Alanine",
		},
		{
			"Ala", "GCC", "Alanine",
		},
		{
			"Ala", "GCA", "Alanine",
		},
		{
			"Ala", "GCG", "Alanine",
		},
		{
			"Asp", "GAU", "Aspartate",
		},
		{
			"Asp", "GAC", "Aspartate",
		},
		{
			"Glu", "GAA", "Glutamate",
		},
		{
			"Glu", "GAG", "Glutamate",
		},
		{
			"Gly", "GGU", "Glycine",
		},
		{
			"Gly", "GGC", "Glycine",
		},
		{
			"Gly", "GGA", "Glycine",
		},
		{
			"Gly", "GGG", "Glycine",
		},
	}

	for _, value := range proteins {
		e, err := stmt.Exec(value.Protein, value.Sequence, value.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e)
	}


}

type protein struct {
	Protein string
	Sequence string
	Name string
}