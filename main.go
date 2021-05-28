package main

import (
	"log"

	"github.com/JL-OLEMAR/twittole/bd"
	"github.com/JL-OLEMAR/twittole/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
