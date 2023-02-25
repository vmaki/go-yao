package helpers

func InArray(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}

	return -1
}

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	return ""
}
