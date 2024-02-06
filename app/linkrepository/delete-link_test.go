package linkrepository

import "testing"

func Test_DeleteLink(t *testing.T) {
	r := NewLinkRepository(t)
	{
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic with missing user ID")
			}
		}()
		_, _ = r.DeleteLink(0, 1)
	}
}
