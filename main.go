package main

import (
	"database/sql"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	// "test/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func getHandler(w http.ResponseWriter, r *http.Request) { //これをアプリケーション層に書くのかな？多分
	fmt.Fprintf(w, "Hello World")
}

func postHandler(w http.ResponseWriter, r *http.Request) { //これをアプリケーション層に書くのかな？多分
	var requestData RequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil { //JSONデータを構造体にデコード
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	responseData := ResponseData{ //%sはString型
		Message: fmt.Sprintf("Hello %s! Your email is %s.", requestData.Name, requestData.Email),
	}

	w.Header().Set("Content-Type", "application/json") //レスポンスがJSON形式であることを示す
	w.WriteHeader(http.StatusOK) //レスポンスのHTTPステータスコードを設定
	json.NewEncoder(w).Encode(responseData) //responseData構造体をJSON形式に変換してレスポンスとしてクライアントに送信
}

var db *sql.DB

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", getHandler)
	mux.HandleFunc("POST /user", postHandler)

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    var err error
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    log.Println("Successfully connected to the database")

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
