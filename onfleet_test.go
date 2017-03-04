package onflt_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/tmc/onflt"
)

func TestBasics(t *testing.T) {
	f, err := os.Open("testdata/sample.json")
	if err != nil {
		t.Skip(err.Error())
		return
	}

	var sg onflt.WorkerSubgraph
	if err := json.NewDecoder(f).Decode(&sg); err != nil {
		t.Error(err)
	}
	spew.Dump(sg)
}
