package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"controle_estoque/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run ./cmd/gerar-chave NOME")
		return
	}

	nome := os.Args[1]

	chave := make([]byte, 32)
	_, err := rand.Read(chave)
	if err != nil {
		log.Fatal("Erro ao gerar chave")
	}

	chaveTexto := hex.EncodeToString(chave)

	hash := sha256.Sum256([]byte(chaveTexto))
	hashHex := hex.EncodeToString(hash[:])

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO api_clients (name, key_hash) VALUES ($1, $2)",
		nome,
		hashHex,
	)
	if err != nil {
		log.Fatalf("Erro ao cadastrar chave no banco de dados: %v", err)
	}

	fmt.Println("Chave de API gerada com sucesso:")
	fmt.Println(chaveTexto)
	fmt.Println("Guarde essa chave em local seguro.")
}
