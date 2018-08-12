// ex1.3
package main

import "testing"

var args []string

func init() {
	args = []string{".echo",
		"kksjd", "kdkishj", "kkdjksj",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi",
		"kdiiii", "kdkso", "kkdisi"}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
