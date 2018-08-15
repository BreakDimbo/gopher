package main

import "fmt"

func main() {
	data := []string{"one", "one", "one", "two", "three", "four", "four"}
	fmt.Printf("%q\n", elimdupEx45(data))
	fmt.Printf("%q\n", data)
}

func nonemtpy(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonemtpy2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func elimdupEx45(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	out := strings[:0]
	out = append(out, strings[0])
	for i := 1; i < len(strings); i++ {
		if strings[i] == strings[i-1] {
			continue
		}
		out = append(out, strings[i])
	}
	return out
}
