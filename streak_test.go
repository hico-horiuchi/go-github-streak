package streak

import "testing"

func TestGetContributions(t *testing.T) {
	_, err := GetContributions("hico-horiuchi")
	if err != nil {
		t.Fatal("GetContributions() expects error: %q", err.Error())
	}
}
