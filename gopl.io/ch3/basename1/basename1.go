package basename

// remove directory compoment and a .suffix

func basename1(s string) string {
	// Discard the slash / and before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func comm(s string) string {
	if len(s) < 3 {
		return s
	}
	return comm(s[:len(s)-3]) + "," + s[len(s)-3:]
}
