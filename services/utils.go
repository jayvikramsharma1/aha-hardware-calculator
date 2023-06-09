package main

import "strconv"

func getFrequencyToDayHours(frequency string) int {
	if frequency == FREQ_PER_HOUR {
		return 24
	} else {
		return 1
	}
}

func getDataSizeInKB(reportSize string) int64 {
	if len(reportSize) < 3 {
		return 0
	}
	size := reportSize[0 : len(reportSize)-2]
	d, err := strconv.ParseInt(size, 10, 32)
	if err != nil {
		return 0
	}
	sizeType := reportSize[len(reportSize)-2:]
	if sizeType == "MB" {
		return d * 1024
	} else {
		return d
	}
}
