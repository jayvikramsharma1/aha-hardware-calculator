package main

import (
	"fmt"
	"testing"
)

func TestGetTotalDataSizeInGB(t *testing.T) {

	got := getTotalDataSizeInGB(testLoad)
	want := int64(16105)
	fmt.Println(got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetTotalSpaceRequiredWithErrorMarginInGB(t *testing.T) {

	got := getTotalSpaceRequiredWithErrorMarginInGB(testLoad)
	want := int64(20131)
	fmt.Println(got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
