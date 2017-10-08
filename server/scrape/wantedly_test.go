package scrape

import (
	"testing"
)

func TestWantedly(t *testing.T) {
	actual := Wantedly()
	t.Errorf("data : %v", actual[0])
}
