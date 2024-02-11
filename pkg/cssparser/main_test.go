package cssparser

import "testing"

func Test_Parse(t *testing.T) {
	{
		css := `body {
			font-weight: bold;
		}`
		result, valid, err := Parse([]byte(css), &ParseOptions{})

		if err != nil {
			t.Fatal(err.Error())
		}

		if !valid {
			t.Error("expected valid result")
		}

		if len(result) == 0 {
			t.Error("recieved empty parse result")
		}
	}

	{
		css := `body {
			font-wight: bold
		}`
		_, valid, err := Parse([]byte(css), &ParseOptions{})
		if err != nil {
			t.Fatal(err.Error())
		}

		if valid {
			t.Error("expected invalid result")
		}
	}
}
