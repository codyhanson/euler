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

func GetNextPrimeFaster(n int) int {
	i := n
	for {
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
		if IsPrimeFaster(i) {
			return i
		}
	}
}

func IsPrimeFaster(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	j := 5
	for j*j <= n {
		if n%j == 0 || n%(j+2) == 0 {
			return false
		}
		j += 6
	}
	return true
}

func IsPrime(p int) bool {
	for i := 2; i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}
