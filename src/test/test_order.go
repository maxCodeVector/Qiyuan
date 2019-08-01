package test

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	orderList := getAllOrders()

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != orderList[i].Content ||
			v.ID != orderList[i].ID ||
			v.Title != orderList[i].Title {

			t.Fail()
			break
		}
	}
}