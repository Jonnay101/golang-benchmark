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
