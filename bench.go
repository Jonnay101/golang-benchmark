package bench

import "strings"

func getLastnamesAppend(fullnames []string) []string {
	lastNames := make([]string, 0)

	for _, fn := range fullnames {
		if fn == "" {
			continue
		}
		allNames := strings.Fields(fn)
		lastNameIdx := len(allNames) - 1

		lastNames = append(lastNames, allNames[lastNameIdx])
	}

	return lastNames
}

func getLastnamesIndex(fullnames []string) []string {
	lastNames := make([]string, len(fullnames))

	idx := -1

	for _, fn := range fullnames {
		if fn == "" {
			continue
		}
		allNames := strings.Fields(fn)
		lastNameIdx := len(allNames) - 1

		idx++
		lastNames[idx] = allNames[lastNameIdx]
	}

	return lastNames[:idx+1]
}

func addNumsLoop(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func addNumsRecursive(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + addNumsRecursive(nums[1:]...)
}
