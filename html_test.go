package wring

import (
	"bytes"
	"strings"
	"testing"
)

func TestHTML(t *testing.T) {
	s := `<!DOCTYPE html>
<html lang="ja">
  <head>
    <title> testing... </title>
  </head>

  <body>
  <h1>  test  title </h1>
  </body>
</html>
	`
	want := `<!DOCTYPE html><html lang="ja"><title>testing...</title><h1>test  title</h1>`
	r := strings.NewReader(s)
	var b bytes.Buffer
	err := HTML(r, &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	got := string(b.Bytes())
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
