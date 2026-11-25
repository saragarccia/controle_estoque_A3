package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/hex"
	"errors"
	"net/http"
)

var rotasPublicas = map[string]bool{
	"GET /health": true,
}

func APIKey(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rotasPublicas[r.Method+" "+r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		chave := r.Header.Get("X-API-Key")

		if chave == "" {
			http.Error(w, "Chave de API não informada", http.StatusUnauthorized)
			return
		}

		hash := sha256.Sum256([]byte(chave))
		hashHex := hex.EncodeToString(hash[:])

		var hashArmazenado string

		err := db.QueryRow(
			"SELECT key_hash FROM api_clients WHERE key_hash = $1",
			hashHex,
		).Scan(&hashArmazenado)

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Chave de API inválida", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Erro ao verificar chave de API", http.StatusInternalServerError)
			return
		}

		if subtle.ConstantTimeCompare(
			[]byte(hashHex),
			[]byte(hashArmazenado),
		) != 1 {
			http.Error(w, "Chave de API inválida", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
