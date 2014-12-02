package challenge09

/*
  a^2 + b^2 = c^2, a + b + c = 1000
  substitute c by a and b, then we have
  1000*(a+b) - ab = 500000
*/
func PythagoreanTripletProduct() int {
	a, b, c := 0, 1, 0

	for b = 1; b < 500; b++ {
		// 1000a + 1000b - ab = 500000
		// (1000 - b)*a = 500000 - 1000b
		// a = (500000 - 1000b)/(1000 - b)
		if (500000-1000*b)%(1000-b) == 0 {
			a = (500000 - 1000*b) / (1000 - b)
			c = 1000 - a - b
			return (a * b * c)
		}
	}
	return -1
}
