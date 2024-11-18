package main

func SwapKeysAndValues(m map[string]string) map[string]string {
    newMap := make(map[string]string)

	for key, v := range m {
		newMap[v] = key
	}

	return newMap
}