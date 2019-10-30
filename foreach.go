package piscine

func ForEach(f func(int), arr []int) {
	for _, letter := range arr {
		f(letter)
	}
}
