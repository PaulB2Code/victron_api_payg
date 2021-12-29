package victronapipayg

import (
	"testing"
)

func Test_open(t *testing.T) {

	apiVictron, _ := NewVictronAPI()

	response, err := apiVictron.GenerateCustomToken("HQ211849AU6", 874130296, "ABB6308F7CC1D6B75DCC447135C678DD", 1, "set_time", 1, 1)
	if response.Token != 332678297 {
		t.Fatalf("NO MATCH %v", err)
	}
	if response.Counter != 3 {
		t.Fatalf("NO MATCH")
	}
}
