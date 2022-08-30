package main

func rotate(matrix [][]int) {
	// 先翻转[i,j] ==> [j,i]
	// 在行内反转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ { //遍历一半
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[i])-j-1] = matrix[i][len(matrix[i])-j-1], matrix[i][j]
		}
	}
}
