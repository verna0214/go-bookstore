package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/verna0214/go-bookstore/pkg/utils"
	"github.com/verna0214/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Connect-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}