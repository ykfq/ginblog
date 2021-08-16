package models

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Chcek that the lenght of the list of articles returned is the
	// same as the lenght of the global varible holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}

}
