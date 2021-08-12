package dbcon

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	User     string `json:"user"`
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Sslmode  string `json:"sslmode"`
}

func readConfigFile(filename string) (Config, error) {
	var setup Config
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(file, &setup)
	return setup, err
}

func ConectaBD() *sql.DB {
	wdir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	configFileName := wdir + "/config/dbsetup.json"
	loadConfig, _ := readConfigFile(configFileName)
	conexao := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", loadConfig.User, loadConfig.Dbname, loadConfig.Password, loadConfig.Host, loadConfig.Sslmode)
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
