package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!")

}

func PostingArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "postingArticle!!!")
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "getArticles!!!")
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
