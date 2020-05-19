package main

func makeResultReceiver(length int) []interface{} {
	values := make([]interface{}, length)
	valuePointers := make([]interface{}, length)
	for i := 0; i < length; i++ {
		valuePointers[i] = &values[i]
	}
	return valuePointers
}
