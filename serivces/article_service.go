package service

import (
	"fmt"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()

	fmt.Println(44)
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

	return article, nil

}

func PostArticleService(article models.Article) (models.Article, error) {
	// TODO : 実装
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}

	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err

	}
	return newArticle, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
