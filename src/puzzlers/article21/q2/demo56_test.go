package q2

import (
	"fmt"
	"testing"
)

var expectedPrimes = []int{
	2, 3, 5, 7, 11, 13, 17, 19,
	23, 29, 31, 37, 41, 43, 47, 53,
	59, 61, 67, 71, 73, 79, 83, 89,
	97, 101, 103, 107, 109, 113, 127, 131,
	137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223,
	227, 229, 233, 239, 241, 251, 257, 263,
	269, 271, 277, 281, 283, 293, 307, 311,
	313, 317, 331, 337, 347, 349, 353, 359,
	367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457,
	461, 463, 467, 479, 487, 491, 499, 503,
	509, 521, 523, 541, 547, 557, 563, 569,
	571, 577, 587, 593, 599, 601, 607, 613,
	617, 619, 631, 641, 643, 647, 653, 659,
	661, 673, 677, 683, 691, 701, 709, 719,
	727, 733, 739, 743, 751, 757, 761, 769,
	773, 787, 797, 809, 811, 821, 823, 827,
	829, 839, 853, 857, 859, 863, 877, 881,
	883, 887, 907, 911, 919, 929, 937, 941,
	947, 953, 967, 971, 977, 983, 991, 997,
}

func TestGetPrimesWithAbnormalParam(t *testing.T) {
	t.Parallel()
	invalidParams := []int{1, 0, -1, -2, -3}
	for _, ip := range invalidParams {
		result := GetPrimes(ip)
		if result == nil {
			t.Errorf("The result is nil, but it should not be.")
		}
		if len(result) > 0 {
			t.Errorf("The result is not empty, but it should be.")
		}
	}
}

func TestGetPrimesWith1000(t *testing.T) {
	t.Parallel()
	max := 1000
	primes := GetPrimes(max)
	for i, prime := range primes {
		expectedPrime := expectedPrimes[i]
		if prime != expectedPrime {
			t.Errorf("%dth prime number %d is not the expected value %d",
				i, prime, expectedPrime)
		}
	}
	if t.Failed() == false {
		t.Logf("The primes less than %d are all correct.", max)
	}
}

func TestGetPrimesParallel(t *testing.T) {
	for _, max := range []int{1, 2, 3, 4, 5} {
		max := max * 200
		// 这些子测试会被并发地执行，
		// 并且只有它们都执行完毕之后当前的测试函数才会执行完成。
		// 当前的测试函数并不会与其他测试函数一起被并发的执行。
		t.Run(fmt.Sprintf("TestGetPrimesWith%d", max),
			func(t *testing.T) {
				t.Parallel()
				primes := GetPrimes(max)
				err := comparePrimes(primes)
				if err != nil {
					t.Error(err)
				} else {
					t.Logf("The primes less than %d are all correct.", max)
				}
			})
	}
}

func comparePrimes(primes []int) (err error) {
	for i, prime := range primes {
		expectedPrime := expectedPrimes[i]
		if prime != expectedPrime {
			err = fmt.Errorf(
				"%dth prime number %d is not the expected value %d",
				i, prime, expectedPrime)
			break
		}
	}
	return
}
