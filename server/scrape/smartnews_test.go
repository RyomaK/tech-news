package scrape

import (
	"testing"
)

func TestSmartNews(t *testing.T) {
	actual := SmartNews()
	t.Errorf("data : %v", actual[0])
}
