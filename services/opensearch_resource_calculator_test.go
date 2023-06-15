package main

import (
	"fmt"
	"testing"
)

type testData struct {
	load     *SystemLoads
	testName string
}

var testDatas = []SystemLoads{
	*testLoad1000NodesWith4MbReportPerDay,
	*testLoad1000NodesWith2MbReportPerDay,
	*testLoad1000NodesWith1MbReportPerDay,
	*testLoad2000NodesWith4MbReportPerDay,
	*testLoad2000NodesWith2MbReportPerDay,
	*testLoad2000NodesWith1MbReportPerDay,
	*testLoad3000NodesWith4MbReportPerDay,
	*testLoad3000NodesWith2MbReportPerDay,
	*testLoad3000NodesWith1MbReportPerDay,
	*testLoad3000NodesWith4MbReport,
	*testLoad3000NodesWith2MbReport,
	*testLoad3000NodesWith1MbReport,
	*testLoad5000NodesWith4MbReportPerDay,
	*testLoad5000NodesWith2MbReportPerDay,
	*testLoad5000NodesWith1MbReportPerDay,
	*testLoad5000NodesWith4MbReport,
	*testLoad5000NodesWith2MbReport,
	*testLoad5000NodesWith1MbReport,
	*testLoad10000NodesWith4MbReportPerDay,
	*testLoad10000NodesWith2MbReportPerDay,
	*testLoad10000NodesWith1MbReportPerDay,
	*testLoad10000NodesWith4MbReport,
	*testLoad10000NodesWith2MbReport,
	*testLoad10000NodesWith1MbReport,
	*testLoad20000NodesWith4MbReportPerDay,
	*testLoad20000NodesWith2MbReportPerDay,
	*testLoad20000NodesWith1MbReportPerDay,
	*testLoad20000NodesWith4MbReport,
	*testLoad20000NodesWith2MbReport,
	*testLoad20000NodesWith1MbReport,
	*testLoad40000NodesWith4MbReportPerDay,
	*testLoad40000NodesWith2MbReportPerDay,
	*testLoad40000NodesWith1MbReportPerDay,
	*testLoad40000NodesWith4MbReport,
	*testLoad40000NodesWith2MbReport,
	*testLoad40000NodesWith1MbReport,
	*testLoad80000NodesWith4MbReport,
	*testLoad80000NodesWith2MbReport,
	*testLoad80000NodesWith1MbReport,
	*testLoad80000NodesWith2MbReportPerDay,
	*testLoad80000NodesWith1MbReportPerDay,
	*testLoad80000NodesWith4MbReport_perDay,
}

func TestGetOpensearchNodeCountRequired(t *testing.T) {
	nodesRequired, memory := getOpensearchNodesWihtMemory(testLoad80000NodesWith4MbReport)
	expected := int64(10)
	fmt.Println(nodesRequired)
	fmt.Println(memory)
	if nodesRequired != expected {
		t.Errorf("got %q, wanted %q", nodesRequired, expected)
	}
}

func TestGetOpensearchNodeCountForDifferentDataAndVariables(t *testing.T) {
	for _, test := range testDatas {
		t.Name()
		fmt.Println(test.name)
		nodesRequired, memory := getOpensearchNodesWihtMemory(&test)
		fmt.Println(nodesRequired)
		fmt.Println(memory)
	}
	if len(testDatas) < 1 {
		t.Error("Empty test list")
	}
}
