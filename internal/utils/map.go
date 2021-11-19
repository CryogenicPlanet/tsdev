package utils

func MergeStringMaps(maps ...map[string]string) (result map[string]string) {
	result = make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
