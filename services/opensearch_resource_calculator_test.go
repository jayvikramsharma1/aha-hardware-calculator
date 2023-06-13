package main

import (
	"fmt"
	"testing"
)

type testData struct {
	load     *SystemLoads
	testName string
}

var testDatas = []testData{
	//{testLoad, "NormalTest"},
	//{testLoad3000NodesWith4MbReport, "Test with 3000 nodes and 4Mb of compliance report size"},
	//{testLoad3000NodesWith2MbReport, "Test with 3000 nodes and 2Mb of compliance report size"},
	//{testLoad3000NodesWith1MbReport, "Test with 3000 nodes and 1Mb of compliance report size"},
	//{testLoad5000NodesWith4MbReport, "Test with 5000 nodes and 4Mb of compliance report size"},
	//{testLoad5000NodesWith2MbReport, "Test with 5000 nodes and 2Mb of compliance report size"},
	//{testLoad5000NodesWith1MbReport, "Test with 5000 nodes and 1Mb of compliance report size"},
	//{testLoad10000NodesWith4MbReport, "Test with 10000 nodes and 4Mb of compliance report size"},
	//{testLoad10000NodesWith2MbReport, "Test with 10000 nodes and 2Mb of compliance report size"},
	//{testLoad10000NodesWith1MbReport, "Test with 10000 nodes and 1Mb of compliance report size"},
	//{testLoad20000NodesWith4MbReport, "Test with 20000 nodes and 4Mb of compliance report size"},
	//{testLoad20000NodesWith2MbReport, "Test with 20000 nodes and 2Mb of compliance report size"},
	//{testLoad20000NodesWith1MbReport, "Test with 20000 nodes and 1Mb of compliance report size"},
	//{testLoad40000NodesWith4MbReport, "Test with 40000 nodes and 4Mb of compliance report size"},
	//{testLoad40000NodesWith2MbReport, "Test with 40000 nodes and 2Mb of compliance report size"},
	//{testLoad40000NodesWith1MbReport, "Test with 40000 nodes and 1Mb of compliance report size"},
	//{testLoad80000NodesWith4MbReport, "Test with 80000 nodes and 4Mb of compliance report size"},
	//{testLoad80000NodesWith2MbReport, "Test with 80000 nodes and 2Mb of compliance report size"},
	{testLoad80000NodesWith1MbReport, "Test with 80000 nodes and 1Mb of compliance report size"},
}

func TestGetOpensearchNodeCountRequired(t *testing.T) {
	nodesRequired := getOpensearchNodeCountRequired(testLoad)
	expected := int64(10)
	fmt.Println(nodesRequired)
	if nodesRequired != expected {
		t.Errorf("got %q, wanted %q", nodesRequired, expected)
	}
}

func TestGetOpensearchNodeCountForDifferentDataAndVariables(t *testing.T) {
	for _, test := range testDatas {
		t.Name()
		fmt.Println(test.testName)
		nodesRequired := getOpensearchNodeCountRequired(test.load)
		fmt.Println(nodesRequired)
	}
	if len(testDatas) < 1 {
		t.Error("Empty test list")
	}
}
