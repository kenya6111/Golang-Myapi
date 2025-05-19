Go では main 以外のパッケージ名は、基本的にそのファイルが格納されているディレクトリ名と
同名にする必要があります

main.go にあったときには、ハンドラは main 関数のスコープ内でハンドラを定義していたため、
「変数 helloHandler に代入演算子:=でハンドラ関数を代入する」という記法を使うことができま
した。しかし、どこかの関数の中ではない場所で代入演算子:=は使えません。そのため、main 関
数のスコープ外である handler パッケージの中では、「新しく関数を定義して、それをパッケージ
外部に公開する」という形でハンドラ関数を定義しました。

Go では、他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大
文字から始める必要があります。逆にいうと、小文字から始まる名前のものは他パッケージには
非公開なものになります。

// localhost:8080 にてサーバーを起動
log.Fatal(http.ListenAndServe(":8080", nil))
http.ListenAndServe 関数の返り値 error は、サーバー起動時に起こったエラーがあったときに
返ってきます。もしそれが起こった場合にエラーの内容を拾ってログとして出力するために、返
り値のエラーを log.Fatal 関数に渡しています。
// 一行で書いたパターン
log.Fatal(http.ListenAndServe(":8080", nil))
// 二行で書いたパターン
err := http.ListenAndServe(":8080", nil)
log.Fatal(err)
log.Fatal 関数は引数に渡された内容をログとして出力しますが、log.Println 関数と異なると
ころとしては、これが実行されると「実行時に異常が起こった」というステータスを伴ってプロ
グラムが終了されるということです。そのため「プログラム実行に支障をきたすような重大なエ
ラーが発生した際に、該当エラーをログ出力させた上でその時点でプログラムを終わらせる」と
いうときに使用されます。


curl コマンドで HTTP メソッドを指定するには、-X オプションを使用します。
$ curl http://localhost:8080/hello -X GET
Hello, world!


今回は if 文での条件分岐を req.Method == "GET"とメソッド名をハードコーディングせずに、
req.Method == http.MethodGet と定数を使って記述しました。
これはもし将来「GET というメソッドの名前を、GET2 と改名します！」という変更15が加わって
しまったときに、いちいちソースコード内の GET という文字列を GET2 に書き換える手間をなくす
ために重要です。もしもそのような変更が加わったときには、Go の net/http パッケージにおけ
る定数 MethodGet の定義も GET2 に置き換わるアップデートが加わるはずですので、この定数を
使用していた場合にはそのような苦労をしなくて済むのです。

$ curl http://localhost:8080/hello -X DELETE -w '%{http_code}\n'
Invalid method
405
curl コマンドに-w '%{http_code}\n' というオプションをつけたことによって、レスポンスの
ステータスコードを明示するようにしています。きちんと GET のときにだけ正常応答 200 番が
が得られていて、他は 405 番のステータスコードが返ってきていることが確認できました。


[コラム] http.ListenAndServe 関数の第二引数と Go のデフォルトルータ
http.ListenAndServe 関数の第二引数というのは、実は「サーバーの中で使うルータを指定する」
部分なのです。ここにルータが渡されず nil だった場合には、Go の HTTP サーバーがデフォル
トで持っているルータが自動的に採用されます。


func Atoi(s string) (int, error)
Atoi 関数の第二戻り値 error は、int 型に直せないような値が引数として与えられたときにエ
ラーを返すためのものです。int 型に直せないような値がもしここに入るということは、リクエス
23https://pkg.go.dev/github.com/gorilla/mux#Vars
24https://pkg.go.dev/strconv#Atoi
65
トのパスパラメータの値がおかしいということです。そのため、その場合には「リクエストが不
正だ」という意味の 400 番エラーを返すことにします。25
articleID, err := strconv.Atoi(mux.Vars(req)["id"])
if err != nil {
// 400 番エラー (BadRequest) を返す
http.Error(w, "Invalid query parameter", http.StatusBadRequest)
return
}