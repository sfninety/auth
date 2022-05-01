package handler

import "testing"

func TestParsePhoneNumber(t *testing.T) {
	scenarios := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{
			"valid with leading 0",
			"+4407931013928",
			"+447931013928",
		},
		{
			"valid without leading 0",
			"+447931013928",
			"+447931013928",
		},
		{
			"invalid length",
			"+44793101392",
			errInvalidPhoneNumber,
		},
		{
			"no leading country code",
			"7931013928",
			errInvalidPhoneNumber,
		},
	}

	for test, sc := range scenarios {
		result, err := parsePhoneNumber(sc.input)
		if err != nil {
			if err != sc.expected {
				t.Fatalf("test %d: %v returned invalid when not expected: \n input: %v \n output: %v \n expected: %v", test, sc.name, sc.input, err.Error(), sc.expected)
				return
			}
			return
		}
		if result != sc.expected {
			t.Fatalf("test %d: %v returned wrong result: \n input: %v \n output: %v \n expected: %v", test, sc.name, sc.input, result, sc.expected)
		}
	}
}
