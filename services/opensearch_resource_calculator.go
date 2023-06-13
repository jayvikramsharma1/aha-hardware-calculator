package main

import (
	"fmt"
	"math"
)

func getOpensearchNodeCountRequired(loads *SystemLoads) int64 {
	fmt.Println(getTotalSpaceRequiredWithErrorMarginInGB(loads))
	totalNodesRequired := float64(getTotalSpaceRequiredWithErrorMarginInGB(loads)) / float64(getMemoryRequiredForEachOpensearchNodes(loads)) / float64(30)
	return int64(math.Ceil(totalNodesRequired))
}
