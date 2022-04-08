package utils

// GroupBy 根据byKey将List中数据分组
func GroupBy(arr []map[string]interface{}, byKey string) map[interface{}][]interface{} {
	m := make(map[interface{}][]interface{})
	for _, a := range arr {
		k := a[byKey]
		if m[k] == nil {
			m[k] = make([]interface{}, 0)
		}
		m[k] = append(m[k], a)
	}
	return m
}
