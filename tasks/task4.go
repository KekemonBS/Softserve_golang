package tasks

func Task4() func() (n int) {
	var prev int
	var curr int
	var next int
	return func() (fibNextNum int) {
		next = curr + prev

		if prev == 0 {
			curr = 1
		}

		prev = curr
		curr = next
		return next
	}
}
