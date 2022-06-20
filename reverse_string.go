package reverse_string

func ReverseString(input string) (output string) {
	bytes := []byte(input)

	var reversedBytes []byte
	for k, _ := range bytes {
		reversedBytes = append(reversedBytes, bytes[len(bytes)-1-k])
	}
	output = string(reversedBytes)
	return output
}
