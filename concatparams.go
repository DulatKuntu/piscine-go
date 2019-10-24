package piscine

func ConcatParams(args []string) string {
	var res string
	for i, item := range args {
		if i > 0 {
			item = "\n" + item
		}
		res = res + item
	}
	return res
}
