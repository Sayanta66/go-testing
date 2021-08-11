package main

func main() {
	println(add(100,100))
}

func add(n, add int) int {
	return n+add
}

func multiply(n, multiplyBy int) int {
	result := n
	for i := 1; i < multiplyBy; i++ {
		result = add(result, n)
	}
	return result
}