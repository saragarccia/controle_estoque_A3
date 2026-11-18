package main

import (
	"fmt"
	"controle_estoque/config"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco:", err)
		return
	}

	defer db.Close()

	fmt.Println("Conexão com o PostgreSQL realizada com sucesso!")
}