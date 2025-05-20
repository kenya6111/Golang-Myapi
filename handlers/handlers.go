package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!!")

}

func PostingArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		_, errMsg := fmt.Printf("failed to encode json (page %d)\n", page)
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		_, errMsg := fmt.Printf("failed to encode json (page %d)\n", articleID)
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func CommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
