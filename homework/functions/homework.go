package main

func Map(data []int, action func(int) int) []int {
	if len(data) == 0 {
		return data
	}

	for i, item := range data {
		data[i] = action(item)
	}

	return data
}

func Filter(data []int, action func(int) bool) []int {
	if len(data) == 0 {
		return data
	}

	filteredData := make([]int, 0)
	for _, item := range data {
		if action(item) {
			filteredData = append(filteredData, item)
		}
	}

	return filteredData
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	result := initial

	for _, item := range data {
		result = action(item, result)
	}

	return result
}
