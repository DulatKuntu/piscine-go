package piscine

func BasicAtoi2(s string) int {
	var t int
	byt := []byte(s)
	if byt[0] == '-'{
		t = t * -1
	}
	for _, word := range byt {
		if word == '1' {
			t = t*10 + 1
		} else if word == '2' {
			t = t*10 + 2
		} else if word == '3' {
			t = t*10 + 3
		} else if word == '4' {
			t = t*10 + 4
		} else if word == '5' {
			t = t*10 + 5
		} else if word == '6' {
			t = t*10 + 6
		} else if word == '7' {
			t = t*10 + 7
		} else if word == '8' {
			t = t*10 + 8
		} else if word == '9' {
			t = t*10 + 9
		} else if word == '0' {
			t = t*10 + 0
		} else {
			return 0
		}
	}
	return t
}
