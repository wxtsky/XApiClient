package utils

func DeepCopy(source interface{}) interface{} {
	switch source := source.(type) {
	case []interface{}:
		copiedSlice := make([]interface{}, len(source))
		for i, v := range source {
			copiedSlice[i] = DeepCopy(v)
		}
		return copiedSlice
	case map[string]interface{}:
		copiedMap := make(map[string]interface{})
		for k, v := range source {
			copiedMap[k] = DeepCopy(v)
		}
		return copiedMap
	}
	return source
}
