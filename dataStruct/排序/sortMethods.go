package main

// O(n2)
func bubbleSort(in []int) {
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] < in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
}

// O(n2)
func selectionSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		maxV := in[i]
		maxI := i
		for j := i; j < len(in)-1; j++ {
			if maxV < in[j+1] {
				maxV = in[j+1]
				maxI = j + 1
			}
		}
		if maxV != in[i] {
			in[i], in[maxI] = in[maxI], in[i]
		}
	}
}

// O(n2)
func insertionSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		insertV := in[i+1]
		insertI := i + 1
		for j := i; j+1 > 0; j-- {
			if insertI == 0 {
				break
			}
			if insertV > in[insertI-1] {
				in[insertI], in[insertI-1] = in[insertI-1], in[insertI]
				insertI--
			}
		}
	}
}

// O(N*logN)
func quickSort(in []int) {
	if len(in) < 2 {
		return
	}

	l, r, p := 0, len(in)-1, in[0]
	for i := 1; i <= r; {
		if in[i] < p {
			in[l], in[i] = in[i], in[l]
			l++
			i++
		} else {
			in[r], in[i] = in[i], in[r]
			r--
		}
	}

	quickSort(in[:l])
	quickSort(in[l+1:])
}
