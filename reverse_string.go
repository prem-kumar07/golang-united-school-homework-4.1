package reverse_string

func ReverseString(input string) (output string) {
	//Place your code here
	byte_str := []rune(input)
	for i := len(byte_str)/2 - 1; i >= 0; i-- {
		rev := len(byte_str) - 1 - i
		byte_str[i], byte_str[rev] = byte_str[rev], byte_str[i]
	}
	output= string(byte_str)
	return output
}
