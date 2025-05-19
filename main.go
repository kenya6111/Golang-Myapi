package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/handlers"
)

type Test struct {
	name string
	age  int
}

func (t *Test) printName() {
	fmt.Println("name: ", t.name)
}

func (t Test) addAge1() {
	t.age += 10
}
func (t *Test) addAge2() {
	t.age += 10
}

// Carのinterface
// Carの機能が関数として入っている
type Car interface {
	run(int) string
	stop()
}

// MyCarの構造体を定義
// 構造体は頭大文字
type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	createdAt time.Time `json:"created_at"`
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
	// t := Test{name: "John", age: 22}
	// fmt.Println(t)
	// t.printName()

	// t.addAge1()
	// fmt.Println(t) // ｔをコピーして渡しているので、addAge1メソッド内で＋10しても、元のtは変わらない
	// t.addAge2()    // tはポインタ型なので、addAge2メソッド内で＋10しても、元のtは変わる
	// fmt.Println(t)

	// var x, y interface{}
	// fmt.Printf("%#v", x) // -> nil
	// x = 1
	// x = 2.1
	// y = []int{1, 2, 3}
	// y = "hello"
	// y = 2
	// fmt.Printf("%v", y) // -> 2

	// myCar := &MyCar{
	// 	name:  "マイカー",
	// 	speed: 0,
	// }

	// fmt.Println(myCar) // &{マイカー 0}

	// // MyCarの構造体を持ち、Carインターフェースを持つ変数の定義
	// var objCar Car = myCar
	// fmt.Println(objCar.run(50)) // 50kmで走ります
	// objCar.stop()               // 停止します

	// http.HandleFunc("/", helloHandler)
	// http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/article", handlers.PostingArticleHandler)
	// http.HandleFunc("/article/list", handlers.GetArticleListHandler)
	// http.HandleFunc("/article/1", handlers.GetArticleHandler)
	// http.HandleFunc("/article/nice", handlers.PostingNiceHandler)
	// http.HandleFunc("/comment", handlers.CommentHandler)
	// log.Println("Starting server on :8080")
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		createdAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment2",
		createdAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "this is the test article",
		UserName:    "kenya",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	// fmt.Printf("%+v\n", article)
	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonData)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostingArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostingNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.CommentHandler).Methods(http.MethodPost)

	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))

}
