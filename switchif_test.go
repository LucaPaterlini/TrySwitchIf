package switchif

import "testing"

func BenchmarkTryIf(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TryIf(i%30)
		}
}

func BenchmarkTry2If(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TryIf(i%30)
	}
}

func BenchmarkTrySwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrySwitch(i%30)
	}
}

func BenchmarkTrySwitch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrySwitch(i%30)
	}
}