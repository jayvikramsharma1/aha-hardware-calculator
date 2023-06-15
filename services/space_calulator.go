package main

func getTotalDataSizeInKB(load *SystemLoads) int64 {

	totalSpaceForComplianceScan := load.noOfNodes * (int64(load.noOfComplianceScan) * int64(getFrequencyToDayHours(load.frequencyOfComplianceScan))) * getDataSizeInKB(load.complianceScanReportSize)
	totalSpaceForClientRun := load.noOfNodes * (int64(load.noOfClientRun) * int64(getFrequencyToDayHours(load.frequencyOfClientRun))) * getDataSizeInKB(load.clientRunSize)
	totalSpaceForEventFeed := load.noOfNodes * (int64(load.noOfEventFeed) * int64(getFrequencyToDayHours(load.frequencyOfEventFeed))) * getDataSizeInKB(load.eventFeedUpdateSize)
	return (totalSpaceForComplianceScan + totalSpaceForClientRun + totalSpaceForEventFeed) * int64(load.DataRetentionPolicyInDays) * 2
	//return (totalSpaceForComplianceScan + totalSpaceForClientRun + totalSpaceForEventFeed) * 2
}

func getTotalSpaceRequiredWithErrorMarginInGB(load *SystemLoads) int64 {
	totalSpaceInGB := int64(float64(getTotalDataSizeInGB(load)) * float64(1.25))
	//fmt.Println(totalSpaceInGB)
	return totalSpaceInGB
}

func getTotalDataSizeInGB(load *SystemLoads) int64 {
	spaceInGb := (getTotalDataSizeInKB(load) / 1024) / 1024
	return spaceInGb
}
