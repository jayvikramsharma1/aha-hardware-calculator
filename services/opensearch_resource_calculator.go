package main

func getOpensearchNodeCountRequired(loads *SystemLoads) int64 {
	return getTotalSpaceRequiredWithErrorMarginInGB(loads) / getMemoryRequiredForEachOpensearchNodes(loads) / 30
}
