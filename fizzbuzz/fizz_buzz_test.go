package fizzbuzz

import "testing"

func TestReturnString(t *testing.T) {
	expected := "2"
	if observed := FizzBuzz(2); observed != expected {
		t.Fatalf("FizzBuzz(2)) = %v, want %v", observed, expected)
	}
}
func TestFizz(t *testing.T) {
	expected := "Fizz"
	if observed := FizzBuzz(3); observed != expected {
		t.Fatalf("FizzBuzz(3)) = %v, want %v", observed, expected)
	}
}

func TestBuzz(t *testing.T) {
	expected := "Buzz"
	if observed := FizzBuzz(5); observed != expected {
		t.Fatalf("FizzBuzz(5)) = %v, want %v", observed, expected)
	}
}

func TestFizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	if observed := FizzBuzz(15); observed != expected {
		t.Fatalf("FizzBuzz(15)) = %v, want %v", observed, expected)
	}
}

// BenchmarkFizzBuzz() is a benchmarking function. These functions follow the
// form `func BenchmarkXxx(*testing.B)` and can be used to test the performance
// of your implementation. They may not be present in every exercise, but when
// they are you can run them by including the `-bench` flag with the `go test`
// command, like so: `go test -v --bench . --benchmem`
//
// You will see output similar to the following:
//
// BenchmarkFizzBuzz    2000000000           0.46 ns/op
//
// This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.
//
// While benchmarking can be useful to compare different iterations of the same
// exercise, keep in mind that others will run the same benchmarks on different
// machines, with different specs, so the results from these benchmark tests may
// vary.

// func BenchmarkFizzBuzz(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		FizzBuzz(15)
// 	}
// }
