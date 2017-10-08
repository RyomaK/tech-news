package scrape

import (
	"testing"
)

func TestEureka(t *testing.T) {
	actual := Eureka()
	t.Errorf("date : %v", actual[0])
	//t.Errorf("got %v\nwant %v", actual)
}
