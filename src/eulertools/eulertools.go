package eulertools

func GetNextPrime(i int) int {
	j := i + 1
	for {
		isPrime := true
		for c := 2; c < j; c++ {
			if j%c == 0 {
				// not prime
				isPrime = false
				break
			}
		}

		if isPrime {
			return j
		}
		//otherwise keep searching
		j++
	}

}

func IsPrime(p int) bool {
	for i := 2; i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}
