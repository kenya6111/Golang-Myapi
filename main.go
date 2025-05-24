package main

import (
	"fmt"
	"io"
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

func (t Test) getNameAndAge() string {
	return "名前は" + t.name + "です。年齢は" + strconv.Itoa(t.age) + "歳です。"
}

type Test2 struct {
	name string
	age  int
}

func (t Test2) getNameAndAge() string {
	return "名前は" + t.name + "です。年齢は" + strconv.Itoa(t.age) + "歳です。"
}

func displayNameAndAge(t Test) {
	fmt.Println(t.getNameAndAge())
}
func displayNameAndAge2(t Test2) {
	fmt.Println(t.getNameAndAge())
}

type TestInterface interface {
	getNameAndAge() string
}

func superDisplayNameAndAge(t TestInterface) {
	fmt.Println(t.getNameAndAge())
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

	student1 := Test{
		name: "kenya",
		age:  20,
	}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.age)
	student1.printName()
	student1.addAge1()
	fmt.Println(student1.age)
	student1.addAge2()
	fmt.Println(student1.age)
	fmt.Println(student1.getNameAndAge())
	displayNameAndAge(student1)

	student2 := Test2{
		name: "kenya",
		age:  20,
	}

	fmt.Println(student2.getNameAndAge())
	displayNameAndAge2(student2)

	superDisplayNameAndAge(student1)
	superDisplayNameAndAge(student2)
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

	fmt.Printf("%+v\n", article)

	testHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world \n")

	}
	http.HandleFunc("/test", testHandler)

	fmt.Printf("%T\n", handlers.HelloHandler)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostingNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.CommentHandler).Methods(http.MethodPost)
	r.HandleFunc("/test", testHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))

}
