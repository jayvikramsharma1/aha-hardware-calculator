package main

type SystemLoads struct {
	noOfNodes                   int64
	noOfComplianceScan          int
	frequencyOfComplianceScan   string
	complianceScanReportSize    string
	noOfClientRun               int
	frequencyOfClientRun        string
	clientRunSize               string
	noOfEventFeed               int
	frequencyOfEventFeed        string
	eventFeedUpdateSize         string
	DataRetentionPolicyInDays   int
	noOfShardsInOpensearchIndex int
}

const FREQ_PER_HOUR = "PerHour"
const FREQ_PER_DAY = "PerDay"
