package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
)

func SetRequestHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", ACCESS_CONTROL_URL)
	return w
}

func GetEncryptedPassword(password string) string {
	passwordHashSlice := md5.Sum([]byte(password))
	return hex.EncodeToString(passwordHashSlice[:])
}

func IsUserExists(db *sql.DB, username string) bool {
	sqlStmt := ("select `username` from `users` where `username` = ?")
	err := db.QueryRow(sqlStmt, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}
