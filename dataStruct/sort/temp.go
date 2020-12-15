package main

func bubbleSort1(in []int) {
	if len(in) < 2 {
		return
	}
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] < in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
}

func selectionSort1(in []int) {
	if len(in) < 2 {
		return
	}
	for i := 0; i < len(in); i++ {
		maxV := in[i]
		maxI := i
		for j := i; j < len(in)-1; j++ {
			if in[j+1] > maxV {
				maxV = in[j+1]
				maxI = j + 1
			}
		}
		in[i], in[maxI] = maxV, in[i]
	}
}

func insertionSort1(in []int) {
	if len(in) < 2 {
		return
	}
	for i := 0; i < len(in)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if in[j] < in[j-1] {
				in[j], in[j-1] = in[j-1], in[j]
			}
		}
	}
}

func quickSort1(in []int) {
	if len(in) < 2 {
		return
	}

	l, r, p := 0, len(in)-1, in[0]
	for i := 1; i <= r; {
		if in[i] > p {
			in[l], in[i] = in[i], in[l]
			i++
			l++
		} else {
			in[i], in[r] = in[r], in[i]
			r--
		}
	}

	quickSort1(in[:l])
	quickSort1(in[l+1:])
}
