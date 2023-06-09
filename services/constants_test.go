package main

var (
	testLoad = &SystemLoads{
		noOfNodes:                   8000,
		noOfComplianceScan:          1,
		frequencyOfComplianceScan:   FREQ_PER_HOUR,
		complianceScanReportSize:    "4MB",
		noOfClientRun:               1,
		frequencyOfClientRun:        FREQ_PER_HOUR,
		clientRunSize:               "300KB",
		noOfEventFeed:               1,
		frequencyOfEventFeed:        FREQ_PER_HOUR,
		eventFeedUpdateSize:         "2KB",
		DataRetentionPolicyInDays:   30,
		noOfShardsInOpensearchIndex: 2,
	}
)
