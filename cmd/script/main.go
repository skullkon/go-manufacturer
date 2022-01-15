package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/skullkon/go-manufacturer/internal/db"
	"github.com/skullkon/go-manufacturer/internal/models"
	"net/http"
	"os"
	"strconv"
)

//

func Send(N int, db *db.Database) (int, error) {
	var ans models.Response
	client := http.DefaultClient
	var body bytes.Buffer
	var arr []int64

	res, err := db.GetPosts(N)
	if err != nil {
		logrus.Error("Error on db GetPosts Method: " + err.Error())
		return 0, err
	}

	if err := json.NewEncoder(&body).Encode(res); err != nil {
		logrus.Error("DB request res encoding error: " + err.Error())
		return 0, err
	}

	if len(res) == 0 {
		logrus.Info("All manufactures sent or nothing to send")
		return 0, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/manufacturer", &body)
	if err != nil {
		logrus.Error("Request building error: " + err.Error())
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Client sending request error: " + err.Error())
		return 0, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&ans); err != nil {
		logrus.Error("API resp encoding error: " + err.Error())
		return 0, err
	}

	for _, value := range res {
		arr = append(arr, value.Id)
	}

	if ans.Answer == false {
		logrus.Info("API did not confirm")
		return 0, nil
	}
	if err := db.Update(arr); err != nil {
		logrus.Info("DB update error: " + err.Error())
		return 0, nil
	}
	return len(arr), nil
}

func main() {
	counter := 1
	_ = godotenv.Load()
	DSN := os.Getenv("DSN")
	db, err := db.NewDB(DSN)
	if err != nil {
		logrus.Error("Error on db constructor: " + err.Error())
		return
	}
	//read N for
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------------------------------")
	fmt.Print("Input N: ")
	text, _, _ := reader.ReadLine()
	fmt.Println(text)
	n, err := strconv.Atoi(string(text))
	if err != nil {
		logrus.Error("N Atoi error: " + err.Error())
		return
	}
	logrus.Info(n)

	for counter > 0 {
		counter, _ = Send(n, db)
	}

}
