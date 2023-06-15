package main

import (
	"fmt"
	"math"
)

func getOpensearchNodeCountRequiredZones(loads *SystemLoads, zoneFactor int) int64 {
	fmt.Println(getTotalSpaceRequiredWithErrorMarginInGB(loads))
	totalNodesRequired := float64(getTotalSpaceRequiredWithErrorMarginInGB(loads)) / float64(getMemoryRequiredForEachOpensearchNodes(loads)) / float64(zoneFactor)
	return int64(math.Ceil(totalNodesRequired))
}

func getOpensearchNodeCountRequiredHotZones(loads *SystemLoads) int64 {
	return getOpensearchNodeCountRequiredZones(loads, 30)
}

func getOpensearchNodeCountRequiredWarmZones(loads *SystemLoads) int64 {
	return getOpensearchNodeCountRequiredZones(loads, 160)
}

func getOpensearchNodeCount(loads *SystemLoads) int64 {
	if loads.DataRetentionPolicyInDays > 1 {
		nodesWithWarmZone := getOpensearchNodeCountRequiredWarmZones(loads)
		loads.DataRetentionPolicyInDays = 1
		nodesWithHotZone := getOpensearchNodeCountRequiredHotZones(loads)
		return nodesWithWarmZone + nodesWithHotZone
	} else {
		return getOpensearchNodeCountRequiredHotZones(loads)
	}

}
