package handlers

import (
	"io"
	"net/http"
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
	io.WriteString(w, "getArticle!!!")
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "postingNice!!!")
}

func CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "comment!!!")
}
