package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/skullkon/go-manufacturer/internal/models"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/skullkon/go-manufacturer/internal/db"
)

//

func Send(N int, db *db.Database) {
	var ans models.Response
	client := &http.Client{}
	var body bytes.Buffer

	res, err := db.GetPosts(N)
	if err != nil {
		logrus.Error("Error on db GetPosts Method: " + err.Error())
		return
	}

	err = json.NewEncoder(&body).Encode(res)
	if err != nil {
		logrus.Error("DB request res encoding error: " + err.Error())
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/manufacturer", &body)
	if err != nil {
		logrus.Error("Request building error: " + err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Client sending request error: " + err.Error())
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&ans)
	if err != nil {
		logrus.Error("API resp encoding error: " + err.Error())
		return
	}
	logrus.Info(ans.Answer)
	if ans.Answer == true {

	}
}

func main() {
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
	lsdf := []int64{3, 4}
	err = db.Update(lsdf)
	if err != nil {
		logrus.Info(err)
		return
	}
	//Send(n, db)

}
