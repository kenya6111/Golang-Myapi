package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
