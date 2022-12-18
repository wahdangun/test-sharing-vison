package utils

import "test-sharing-vision/app/models"

// valdiation for article models
func ValidateArticle(article models.Article) (bool, map[string]interface{}) {
	// define error message
	var message = make(map[string]interface{})

	// check title
	if len(article.Title) < 20 {
		message["title"] = "Title at least 20 characters"
	}

	// check content
	if len(article.Content) < 200 {
		message["content"] = "Content at least 200 characters"
	}

	// check category
	if len(article.Category) < 3 {
		message["category"] = "Category at least 3 characters"
	}

	// check status
	if article.Status != "published" && article.Status != "draft" && article.Status != "thrash" {
		message["status"] = "Status is invalid"
	}

	// check if error message is empty
	if len(message) > 0 {
		return false, message
	}

	// return true if no error
	return true, nil
}
