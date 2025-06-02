package service

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func getArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 3. 2 で得たコメント一覧を、1 で得た Article 構造体に紐付ける
	article.CommentList = append(article.CommentList, commentList...)

	return article, err

}

func PostArticleService(article models.Article) (models.Article, error) {
	// TODO : 実装
	return models.Article{}, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	// TODO : 実装
	return nil, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	// TODO : 実装
	return models.Article{}, nil
}

// PostCommentHandler で使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func PostCommentService(comment models.Comment) (models.Comment, error) {
	// TODO : 実装
	return models.Comment{}, nil
}
