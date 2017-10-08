package scrape

import (
	"testing"
)

func TestAjito(t *testing.T) {
	actual := Ajito()
	t.Errorf("data : %v", actual[0])
}

func TestGetDate(t *testing.T) {
	actual := getDate("Wed, Sep 13, 2017")
	t.Errorf("data : %v", actual)
}
