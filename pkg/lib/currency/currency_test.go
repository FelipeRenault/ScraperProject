package currency

import "testing"

func Test_ParsePrice(t *testing.T) {
	parsed, err := ParsePrice("R$ 1.956,90")
	if err != nil {
		t.Error(err)
	}
	t.Log(parsed)
}
