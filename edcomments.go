package main

import (
	"flag"
	"github.com/leonus96/edcommets/migration"
	"log"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB.")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Fin de la migración. :3")
	}
}
