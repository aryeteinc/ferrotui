package main

import (
	"fmt"

	"github.com/aryeteinc/ferrotui/internal/db"
)

func main() {
	fmt.Print("Iniciando ferrotui...")
	_, err := db.Connect()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Base de Datos Conectada...")
}
