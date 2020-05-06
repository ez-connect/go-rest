package core

func IndexOf(values []string, value string) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}

	return -1
}
