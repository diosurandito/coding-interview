package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	switch true {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}


// func getFizzBuzz(n int) string {
//     switch true {
//     case n%15 == 0:
//         return "FizzBuzz"
//     case n%3 == 0:
//         return "Fizz"
//     case n%5 == 0:
//         return "Buzz"
//     default:
//         return strconv.Itoa(n)
//     }

// }

// func fizzBuzz(n int32) {
//     var fb string
//     for i := 1; i <= int(n); i++ {
//         fb = getFizzBuzz(i)
//         fmt.Println(fb)
//     }
// }
