package v1

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/skullkon/go-manufacturer/internal/models"
)

type response struct {
	Answer bool `json:"answer"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var res response
	var batch []models.Manufacturer
	err := json.NewDecoder(r.Body).Decode(&batch)
	if err != nil {
		w.WriteHeader(500)
		logrus.Error("DB error :" + err.Error())
		return
	}
	res.Answer = true
	resp, _ := json.Marshal(res)
	w.Write(resp)
}
