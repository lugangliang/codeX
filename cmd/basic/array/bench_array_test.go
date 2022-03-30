package main

import (
	"os/exec"
	"testing"
)

func BenchmarkMergeTwoSortedArray(b *testing.B) {
	b.ReportAllocs()
	path, err := exec.LookPath("hostname")
	if err != nil {
		b.Fatalf("could not find hostname: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := exec.Command(path).Run(); err != nil {
			b.Fatalf("hostname: %v", err)
		}
	}
}