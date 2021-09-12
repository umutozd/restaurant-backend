package utils

// StringSliceContains reports whether given slice contains element or not.
func StringSliceContains(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}
