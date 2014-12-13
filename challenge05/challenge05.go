package challenge05

func SmallestMultiple(target int) int {
	factors := make([]int, 0)
	product := 1

	for i := 1; i <= target; i++ {
		subFactors := getAllFactors(i)
		factors = mergeFactors(factors, subFactors)
	}

	for _, value := range factors {
		product *= value
	}
	return product
}

func getAllFactors(num int) []int {
	factorSlice := make([]int, 1)
	factorSlice[0] = 1
	i := 2
	for num > 1 {
		if num%i == 0 {
			factorSlice = append(factorSlice, i)
			num /= i
			i = 2
		} else {
			i++
		}
	}
	return factorSlice
}

func mergeFactors(parentFactors, childFactors []int) []int {
	parentCopy := make([]int, len(parentFactors))
	copy(parentCopy, parentFactors)
	found := false

	for _, value := range childFactors {
		if len(parentCopy) == 0 {
			found = false
		} else {
			for i, pVal := range parentCopy {
				if value == pVal {
					parentCopy = append(parentCopy[:i], parentCopy[i+1:]...)
					found = true
					break
				} else {
					found = false
				}
			}
		}
		if !found {
			parentFactors = append(parentFactors, value)
			found = false
		}
	}
	return parentFactors
}
