package Controllers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestAllArticle(t *testing.T){
	AllArticle := AllArticle
	assert.NotNil(t, AllArticle, "Should not be nil")
}


func TestCreateArticle(t *testing.T){
	CreateArticle := CreateArticle
	assert.NotNil(t, CreateArticle, "Should not be nil")
}


func TestUpdateArticle(t *testing.T){
	UpdateArticle := UpdateArticle
	assert.NotNil(t, UpdateArticle, "Should not be nil")
}


func TestDeleteArticle(t *testing.T){
	DeleteArticle := DeleteArticle
	assert.NotNil(t, DeleteArticle, "Should not be nil")
}


func TestArticle(t *testing.T){
	Article := Article
	assert.NotNil(t, Article, "Should not be nil")
}


