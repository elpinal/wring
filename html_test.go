package wring

import (
	"bytes"
	"strings"
	"testing"
)

func TestHTML(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{`<!DOCTYPE html>
<html lang="ja">
  <head>
    <title> testing... </title>
  </head>

  <body>
  <h1>  test  title </h1>
  </body>
</html>
	`, `<!DOCTYPE html><html lang="ja"><title>testing...</title><h1>test  title</h1>`},
	}
	for _, test := range tests {
		r := strings.NewReader(test.input)
		var b bytes.Buffer
		err := HTML(r, &b)
		if err != nil {
			t.Errorf("HTML(%q): %s", test.input, err)
		}
		if got := string(b.Bytes()); got != test.want {
			t.Errorf("HTML(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
