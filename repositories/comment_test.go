package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	// (以下略)
)

// SelectCommentList関数のテスト
func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "CommentInsertTest",
	}

	expectedCommentID := 25
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}
	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where comment_id = ?
		`
		testDB.Exec(sqlStr, newComment.CommentID)
		testDB.Exec("ALTER TABLE comments AUTO_INCREMENT = 3")

	})
}
