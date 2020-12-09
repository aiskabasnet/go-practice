package Models

import (

	//my sql driver
	"go-practice/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllArticles(article *[]Article) (err error) {
	err = Config.DB.Find(article).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateArticle(article *Article) (err error) {
	if err = Config.DB.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func GetArticleById(article *Article, id string) (err error) {
	err = Config.DB.Where("id = ?", id).First(article).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateArticle(article *Article, id string) (err error) {
	err = Config.DB.Save(article).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(article *Article, id string) (err error) {
	err = Config.DB.Where("id = ?", id).Delete(article).Error
	if err != nil {
		return err
	}
	return nil
}
