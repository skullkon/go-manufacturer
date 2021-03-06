package v1

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/skullkon/go-manufacturer/internal/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	var batch []models.Manufacturer
	if err := json.NewDecoder(r.Body).Decode(&batch); err != nil {
		w.WriteHeader(500)
		logrus.Error("DB error :" + err.Error())
		return
	}

	if len(batch) == 0 {
		res.Answer = false
		resp, _ := json.Marshal(res)
		w.WriteHeader(200)
		test, err := w.Write(resp)
		logrus.Info(test)
		if err != nil {
			logrus.Error("Answer Marshall error in handler" + err.Error())
			return
		}
		return
	}
	res.Answer = true
	resp, _ := json.Marshal(res)
	w.WriteHeader(200)
	test, err := w.Write(resp)
	logrus.Info(test)
	if err != nil {
		logrus.Error("Answer Marshall error in handler" + err.Error())
		return
	}
}
