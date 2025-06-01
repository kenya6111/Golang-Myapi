package handlers

import (
	"database/sql"
	"encoding/json"
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
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	article, err := InsertArticle(db, reqArticle)

	if err != nil {
		http.Error(w, "Failed to insert article", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

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

	articleList, err := SelectArticleList(db, page)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(page)

	// articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(articleID)

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	article, err := SelectArticleDetail(db, articleID)
	if err != nil {
		http.Error(w, "Failed to get article detail", http.StatusInternalServerError)
		return
	}
	// article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostingNiceHandler(w http.ResponseWriter, req *http.Request) {
	// article := models.Article1
	queryMap := req.URL.Query()

	var articleID int
	if p, ok := queryMap["id"]; ok && len(p) > 0 {
		var err error
		articleID, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		articleID = 1
	}
	fmt.Println(articleID)
	fmt.Println("articleID" + fmt.Sprint(articleID))

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = UpdateNiceNum(db, articleID)

	if err != nil {
		http.Error(w, "Failed to update nice number", http.StatusInternalServerError)
		return
	}

	// json.NewEncoder(w).Encode(article)
	// var reqArticle models.Article
	// if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
	// 	http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	// }

	// article := reqArticle
	// json.NewEncoder(w).Encode(article)

}

func CommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	// (問 1) 構造体 `models.Article`を受け取って、それをデータベースに挿入する処理
	const sqlStr = `insert into articles (title,contents,username,nice ,created_at) values(?,?,?,?,now())`
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName, article.NiceNum)
	if err != nil {
		fmt.Println("Error inserting article:", err)
		return models.Article{}, err
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	return article, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	select article_id, title, contents, username, nice
	from articles
	limit ? offset ?;
`
	articleArray := make([]models.Article, 0)

	rows, err := db.Query(sqlStr, 5, 0)
	if err != nil {
		fmt.Println(err)
		return []models.Article{}, err
	}

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []models.Article{}, err
		} else {
			articleArray = append(articleArray, article)
		}
	}

	return articleArray, nil
}

// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
	select *
	from articles
	where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)

	if err := row.Err(); err != nil {
		fmt.Println("Error querying article:", err)
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println("Error scanning article:", err)
		return models.Article{}, err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	// (問 3) 指定 ID の記事データをデータベースから取得して、それを models.Article 構造体の形で返す処理
	return article, nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	const sqlGetNice = `
	select nice
	from articles
	where article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, articleID)
	fmt.Println("articleID : " + string(rune(articleID)))
	fmt.Println(row)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	fmt.Println("niceNum : " + string(rune(niceNum)))

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	// (問 4) 指定された ID の記事のいいね数を+1 するようにデータベースの中身を更新する処理
	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

// 新規投稿をDBにinsertする関数
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`
	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
