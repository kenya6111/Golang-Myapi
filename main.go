package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/yourname/reponame/handlers"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// articleID := 1
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	article_id := 1
	const sqlGetNice = `
	select nice from articles where article_id=?;`
	row := tx.QueryRow(sqlGetNice, article_id)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, niceNum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	tx.Commit()

	// const sqlStr = `
	// 	select * from articles
	// 	where article_id = ?;
	// 	;
	// `
	// rows, err := db.Query(sqlStr, articleID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer rows.Close()

	// row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// article := models.Article{
	// 	Title:    "first article",
	// 	Contents: "this is the test article",
	// 	UserName: "kenya",
	// 	NiceNum:  0,
	// }

	// const sqlStr = `
	// insert into articles (title,contents,username, nice,created_at) values (?,?,?,?,now())`

	// result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName, article.NiceNum)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return

	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())
	// fmt.Println("inserted article:", article)
	// articleArray := make([]models.Article, 0)
	// var article models.Article
	// var createdTime sql.NullTime
	// // for rows.Next() {
	// 	var article models.Article
	// 	var createdTime sql.NullTime

	// 	err := rows.Scan(&article.ID, &article.Title, &article.Contents,
	// 		&article.UserName, &article.NiceNum, &createdTime)
	// 	if createdTime.Valid {
	// 		article.CreatedAt = createdTime.Time
	// 	}

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		articleArray = append(articleArray, article)
	// 	}
	// }
	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName,
	// 	&article.NiceNum, &createdTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }
	// fmt.Println("----------")
	// fmt.Printf("%+v\n", article)
	// fmt.Println("----------")

	// comment1 := Comment{
	// 	CommentID: 1,
	// 	ArticleID: 1,
	// 	Message:   "test comment1",
	// 	CreatedAt: time.Now(),
	// }
	// comment2 := Comment{
	// 	CommentID: 2,
	// 	ArticleID: 1,
	// 	Message:   "test comment2",
	// 	CreatedAt: time.Now(),
	// }
	// article = Article{
	// 	ID:          1,
	// 	Title:       "first article",
	// 	Contents:    "this is the test article",
	// 	UserName:    "kenya",
	// 	NiceNum:     1,
	// 	CommentList: []Comment{comment1, comment2},
	// 	CreatedAt:   time.Now(),
	// }
	// fmt.Printf("%+v\n", article)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostingNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.CommentHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8081", r))

}
