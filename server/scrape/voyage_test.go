package scrape

import (
	"testing"
)

func TestVoyage(t *testing.T) {
	actual := Voyage(2017)
	t.Errorf("data : %v", actual[0])
}
