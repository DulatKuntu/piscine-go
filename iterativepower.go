package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0{
		return nb
	} 
	result := nb
	for i := 1; i <= power-1; i++ {
		result = nb * result
	}
	return result
}
