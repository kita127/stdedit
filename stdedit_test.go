package stdedit

import (
	"testing"
)

func TestStdedit(t *testing.T) {

	testTable := []struct {
		input  []byte
		expect []byte
	}{
		{
			input:  []byte("hoge"),
			expect: []byte("hoge"),
		},
	}

	for _, tc := range testTable {
		got := Edit(tc.input)

		if string(got) != string(tc.expect) {
			t.Fatalf("not match expect=%s, got=%s", string(tc.expect), string(got))
		}
	}
}
