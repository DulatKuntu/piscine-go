package piscine

func AlphaCount(str string) int {
	counter := 0
	byt := []byte(str)
	for index, letter := range byt {
		index = index
		if (letter > 64 && letter < 91) || letter > 96 && letter < 123 {
			counter++
		}
	}
	return counter
}
