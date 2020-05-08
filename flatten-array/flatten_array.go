// Package flatten provides functionality around flattening nested arrays
package flatten

// Flatten flattens nested arrays
func Flatten(i interface{}) []interface{} {
	result := make([]interface{}, 0)

	// actual value
	x, ok := i.([]interface{})
	if !ok {
		if i == nil {
			return result
		}
		return append(result, i)
	}

	// filter nils and flatten
	for _, y := range filterNils(x) {
		result = append(result, Flatten(y)...)
	}
	return result
}

func filterNils(i []interface{}) []interface{} {
	var result []interface{}
	for _, j := range i {
		if j == nil {
			continue
		}
		result = append(result, j)
	}
	return result
}
