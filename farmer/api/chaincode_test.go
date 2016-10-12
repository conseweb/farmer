package api

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestChaincodeWrapper(t *testing.T) {
	cw := ChaincodeWrapper{
		Name:     "asdf",
		Method:   "deploy",
		Function: "init",
		Args:     []string{"a", "100", "b", "200"},
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "	")
	if err := enc.Encode(cw.ToJSONRPC()); err != nil {
		t.Error(err)
	}
	t.Logf("data: %s\n", buf.String())
	// t.Error("...")
}
