package main

import "fmt"

var data = [][]int{
	{1487799425, 14, 1},
	{1487799425, 4, 0},
	{1487799425, 2, 0},
	{1487800378, 10, 1},
	{1487801478, 18, 1},
	{1487801478, 18, 0},
	{1487801478, 12, 0},
	{1487901013, 1, 0},
	{1487901211, 7, 1},
	{1487901211, 7, 0}}

// FindBusiestPeriod returns the timestamp during which the mall had the highest incoming traffic.
func FindBusiestPeriod(data [][]int) int {
	dataMap := map[int]int{}
	highestVolume := 0
	timestamp := 0

	for _, tData := range data {
		if _, ok := dataMap[tData[0]]; !ok {
			if tData[2] == 0 {
				dataMap[tData[0]] = -tData[1]
			} else {
				dataMap[tData[0]] = tData[1]
			}
		} else if tData[2] == 0 {
			dataMap[tData[0]] -= tData[1]
		} else {
			dataMap[tData[0]] += tData[1]
		}
		fmt.Println(dataMap)
	}

	for ts, v := range dataMap {
		if v > highestVolume {
			highestVolume = v
			timestamp = ts
		}
	}

	fmt.Println(timestamp)
	return timestamp
}

func main() {
	FindBusiestPeriod(data)
}

// {1487799425, 14, 1},
// {1487799425, 4, 0},
// {1487799425, 2, 0},
// {1487800378, 10, 1},
//  ts
// {1487801478, 18, 0},
//               v

// when 0 -> map[ts] = -v
// when 1 -> map[ts] = v

// {1487801478, 18, 1},
// {1487801478, 12, 0},
// {1487901013, 1, 0},
// {1487901211, 7, 1},
// {1487901211, 7, 0}
