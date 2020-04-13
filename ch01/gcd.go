package main

import "fmt"

func main() {
	p, q := 45, 54
	fmt.Printf("%d, %d = %d", p, q, gcd(p, q))
}

func gcd(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}

func gcd2(p int, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	for r != 0 {
		fmt.Printf("p=%d, q=%d, r=%d\n", p, q, r)
		p = q
		q = r
		r = p % q
	}
	return q
}
