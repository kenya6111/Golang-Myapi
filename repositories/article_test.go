package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	// (以下略)
)

// SelectArticleDetail 関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	// dbUser := "docker"
	// dbPassword := "docker"
	// dbDatabase := "sampledb"
	// dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	// db, err := sql.Open("mysql", dbConn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()

	// 1. テスト結果として期待する値を定義
	// expected := models.Article{
	// 	ID:       1,
	// 	Title:    "firstPost",
	// 	Contents: "This is my first blog",
	// 	UserName: "saki",
	// 	NiceNum:  32,
	// }

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1", // テストのタイトル
			expected: models.Article{ // テストで期待する値
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  32,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}
	for _, test := range tests {
		// Run メソッドの第一引数にはサブテスト名、第二引数にはサブテストの内容を指定
		t.Run(test.testTitle, func(t *testing.T) {
			// 2. テスト対象となる関数を実行
			// -> 結果が got に格納される
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// 3. 2 の結果と 1 の値を比べる
			if got.ID != test.expected.ID {
				// 不一致だった場合にはテスト失敗
				t.Errorf("get %d but want %d\n", got.ID, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				// 不一致だった場合にはテスト失敗
				t.Errorf("get %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				// 不一致だった場合にはテスト失敗
				t.Errorf("get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
	// t.Fatal も t.Errorf も実行されずに関数が終わった場合にはテスト成功
}

// SelectArticleList 関数のテスト
func TestSelectArticleList(t *testing.T) {
	// テストで使うデータベースに接続
	// dbUser := "docker"
	// dbPassword := "docker"
	// dbDatabase := "sampledb"
	// dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	// db, err := sql.Open("mysql", dbConn)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer db.Close()
	// テスト対象の関数を実行
	expectedNum := 5
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	// SelectArticleList 関数から得た Article スライスの長さが期待通りでないなら FAIL にする
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}
	expectedArticleNum := 8
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
			`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}
