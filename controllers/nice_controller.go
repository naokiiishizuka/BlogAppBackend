package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers/services"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

type NiceController struct {
	services services.Niceservicer
}

func NewNiceController(s services.Niceservicer) *NiceController {
	return &NiceController{services: s}
}

func (n *NiceController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqNice models.Nice
	if err := json.NewDecoder(req.Body).Decode(&reqNice); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	nice, err := n.services.PostNiceSerice(reqNice)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(nice)
}
