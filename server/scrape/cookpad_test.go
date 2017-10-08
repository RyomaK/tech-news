package scrape

import (
	"testing"
)

func TestCookpad(t *testing.T) {
	actual := CookPad(2017)
	t.Errorf("data : %v", actual[0])
}

func TestNewCookpad(t *testing.T) {
	actual := NewCookPad(2017)
	t.Errorf("data : %v", actual)
}
