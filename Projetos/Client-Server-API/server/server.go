package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RESPONSE *http.Response

func init() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		DB = db
		DB.AutoMigrate(&USDBRL_DB{})
	}

	response, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		log.Fatal(err)
	}

	RESPONSE = response
}

func main() {
	http.HandleFunc("/conversao", conversaoHandler)
	http.HandleFunc("/save", saveHandler)

	http.ListenAndServe(":8080", nil)
}

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type USDBRL_DB struct {
	id      int `gorm:"PrimaryKey"`
	cotacao float64
	data    string
}

func conversaoHandler(w http.ResponseWriter, r *http.Request) {

	defer RESPONSE.Body.Close()

	dadosEmBytes, err := io.ReadAll(RESPONSE.Body)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(dadosEmBytes)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	dados, err := io.ReadAll(RESPONSE.Body)

	if err != nil {
		log.Fatal(err)
	}

	var cotacao Cotacao

	err = json.Unmarshal(dados, &cotacao)

	save(cotacao.USDBRL.Bid)

	w.Write([]byte("Cotacao salva"))
}

func save(cotacao string) {
	value, err := strconv.ParseFloat(cotacao, 64)

	if err != nil {
		log.Fatalf("Erro ao converter para float")
	}

	DB.Create(&USDBRL_DB{
		cotacao: value,
		data:    string(time.Now().Format("2006-01-02 15:04:05")),
	})
}
