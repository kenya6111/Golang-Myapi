package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
)

type Test struct {
	name string
	age  int
}

func HelloHandler(w http.ResponseWriter, req *http.Request) { //curl http://localhost:8080/hello
	io.WriteString(w, "Hello, World!!")

}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// 1. バイトスライス reqBodybuffer を何らかの形で用意
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)
	// 2. Read メソッドでリクエストボディを読み出し
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}
	// 3. ボディを Close する
	defer req.Body.Close()
	var reqArticle models.Article
	if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	article := reqArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	var reqBodyBuffer []byte
	if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "failed to read request body\n", http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

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
