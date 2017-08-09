


func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(slc []string, rot int) {
	reverse(slc[rot:])
	reverse(slc[:rot])
	reverse(slc)
}