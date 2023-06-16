package main

import (
	"math"
)

func getOpensearchNodeCountRequiredZones(loads *SystemLoads, zoneFactor int, memory int) int64 {
	//fmt.Println(getTotalSpaceRequiredWithErrorMarginInGB(loads))
	totalNodesRequired := float64(getTotalSpaceRequiredWithErrorMarginInGB(loads)) / float64(memory) / float64(zoneFactor)
	return int64(math.Ceil(totalNodesRequired))
}

func getOpensearchNodeCountRequiredHotZones(loads *SystemLoads, memory int) int64 {
	return getOpensearchNodeCountRequiredZones(loads, 30, memory)
}

func getOpensearchNodeCountRequiredWarmZones(loads *SystemLoads, memory int) int64 {
	return getOpensearchNodeCountRequiredZones(loads, 100, memory)
}

func getOpensearchNodeCount(loads *SystemLoads, memory int) int64 {
	if loads.DataRetentionPolicyInDays > 1 {
		nodesWithWarmZone := getOpensearchNodeCountRequiredWarmZones(loads, memory)
		loads.DataRetentionPolicyInDays = 2
		nodesWithHotZone := getOpensearchNodeCountRequiredHotZones(loads, memory)
		return nodesWithWarmZone + nodesWithHotZone
	} else {
		return getOpensearchNodeCountRequiredHotZones(loads, memory)
	}

}

func getOpensearchNodeCountWithMemoryRequiredForEachNode(loads *SystemLoads, memory int) (int64, int) {
	count := getOpensearchNodeCount(loads, memory)
	if count < 3 && memory > 1 {
		count, memory = getOpensearchNodeCountWithMemoryRequiredForEachNode(loads, memory/2)
	}
	if count < 3 && memory <= 1 {
		return 3, memory
	}
	return count, memory
}

func getOpensearchNodesWihtMemory(loads *SystemLoads) (int64, int) {
	return getOpensearchNodeCountWithMemoryRequiredForEachNode(loads, 64)
}
