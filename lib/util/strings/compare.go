package strings

func ContainsSlice(required, contained []string) bool {
	containedMap := make(map[string]bool)
	for _, c := range contained {
		containedMap[c] = true
	}

	for _, r := range required {
		_, ok := containedMap[r]
		if !ok {
			return false
		}
	}
	return true
}
