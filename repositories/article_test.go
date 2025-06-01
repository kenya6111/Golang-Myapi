package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	// (以下略)
)

// SelectArticleDetail 関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// 1. テスト結果として期待する値を定義
	expected := models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  32,
	}

	// 2. テスト対象となる関数を実行
	// -> 結果が got に格納される
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	// 3. 2 の結果と 1 の値を比べる
	if got.ID != expected.ID {
		// 不一致だった場合にはテスト失敗
		t.Errorf("get %d but want %d\n", got.ID, expected.ID)
	}

	if got.Title != expected.Title {
		// 不一致だった場合にはテスト失敗
		t.Errorf("get %s but want %s\n", got.Title, expected.Title)
	}

	if got.Contents != expected.Contents {
		// 不一致だった場合にはテスト失敗
		t.Errorf("get %s but want %s\n", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}

	// t.Fatal も t.Errorf も実行されずに関数が終わった場合にはテスト成功
}
