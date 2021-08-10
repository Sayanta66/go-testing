package main

func add(n, add int) int {
	for i := 1; i <= add; i++ {
			n ++
	}
	return n
}

func multiply(n, multiplyBy int) int {
	result := n
	for i := 1; i < multiplyBy; i++ {
		result = add(result, n)
	}

	return result
}