package helpers

func InArray(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}

	return -1
}
