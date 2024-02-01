package linkrepository

import "testing"

func Test_New(t *testing.T) {
	_ = New(&Config{":memory:"})
}
