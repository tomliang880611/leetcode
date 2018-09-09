package leetcode

func rotate(matrix [][]int) {
	if matrix == nil {
		return
	}
	n := len(matrix)
	if n == 0 {
		return
	}

	// i represent the i th layer
	// and n/2 specifies the total number of layers
	for i := 0; i < n/2; i++ {
		shapes := n - i*2 - 1
		// rotate the each shope
		for s := 0; s < shapes; s++ {
			N := n - 1

			first := matrix[i][i+s]
			second := matrix[i+s][N-i]
			third := matrix[N-i][N-i-s]
			forth := matrix[N-i-s][i]

			matrix[i][i+s] = forth
			matrix[i+s][N-i] = first
			matrix[N-i][N-i-s] = second
			matrix[N-i-s][i] = third
		}
	}
}
