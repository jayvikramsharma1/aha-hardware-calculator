package main

import (
	"fmt"
	"testing"
)

func TestGetOpensearchNodeCountRequired(t *testing.T) {
	nodesRequired := getOpensearchNodeCountRequired(testLoad)
	expected := int64(10)
	fmt.Println(nodesRequired)
	if nodesRequired != expected {
		t.Errorf("got %q, wanted %q", nodesRequired, expected)
	}
}
