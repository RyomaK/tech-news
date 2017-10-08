package scrape

import (
	"testing"
)

func TestCyberAgent(t *testing.T) {
	actual := CyberAgent()
	t.Errorf("data : %v", actual[0])
}
