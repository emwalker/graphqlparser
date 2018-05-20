package parser

import (
	"encoding/json"
	"testing"
)

func TestJson(t *testing.T) {
	ast, _ := Parse(`query { organization { name } }`)
	buf := ast.JSON()
	out := make(map[string]interface{})

	if err := json.Unmarshal(buf, &out); err != nil {
		t.Errorf("unable to parse json: %s, %s", err, string(buf))
	}
}

func TestEmptyQuery(t *testing.T) {
	_, err := Parse(`query { }`)

	if err == nil {
		t.Log("expected an error")
		t.Fail()
	}
}
