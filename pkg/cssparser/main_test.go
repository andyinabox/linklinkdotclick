package cssparser

import "testing"

func Test_Parse(t *testing.T) {
	{
		css := `body {
			font-weight: bold;
		}`
		result, err := Parse([]byte(css), true)

		if err != nil {
			t.Fatal(err.Error())
		}

		if !result.Valid {
			t.Error("expected valid result")
		}

		if len(result.Output) == 0 {
			t.Error("recieved empty parse result")
		}
	}

	{
		css := `body {
			font-wight: bold
		}`
		result, err := Parse([]byte(css), true)
		if err == nil {
			t.Errorf("expected error for invalid css: %v", err)
		}

		if result.Valid {
			t.Error("expected invalid result")
		}
	}
}
