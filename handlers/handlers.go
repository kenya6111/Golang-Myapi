package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// u, _ := url.Parse("http://localhost:8080?page=1&page=2&a=1&")
	// queryMap := u.Query()
	// fmt.Println(queryMap["page"])
	// fmt.Println(queryMap["a"])
	// fmt.Println(queryMap["b"])
	io.WriteString(w, "Hello, World!!")

}

func PostingArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "postingArticle!!!")
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

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
	// io.WriteString(w, "getArticles!!!")
}

func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "postingNice!!!")
}

func CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "comment!!!")
}
