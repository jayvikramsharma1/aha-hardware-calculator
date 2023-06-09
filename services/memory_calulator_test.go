package main

import (
	"fmt"
	"testing"
)

func TestGetMemoryRequiredForEachOpensearchNodes(t *testing.T) {
	memoryRequired := getMemoryRequiredForEachOpensearchNodes(testLoad)
	expected := int64(64)
	fmt.Println(memoryRequired)
	if memoryRequired != expected {
		t.Errorf("got %q, wanted %q", memoryRequired, expected)
	}
}
