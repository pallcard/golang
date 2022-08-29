package main

import (
	"sort"
)

func diagonalSort(mat [][]int) [][]int {

	if len(mat) <= 1 || len(mat[0]) <= 1 {
		return mat
	}

	for i := 0; i < len(mat); i++ {
		ints := make([]int, 0)
		x := i
		for y := 0; x < len(mat) && y < len(mat[0]); y++ {
			ints = append(ints, mat[x][y])
			x++
		}
		sort.Ints(ints)
		x = i
		for y := 0; x < len(mat) && y < len(mat[0]); y++ {
			mat[x][y] = ints[y]
			x++
		}
	}
	//{75, 21, 13, 24, 8},
	//{24, 100, 40, 49, 62}

	for j := 1; j < len(mat[0]); j++ {
		ints := make([]int, 0)
		y := j
		for x := 0; x < len(mat) && y < len(mat[0]); y++ {
			ints = append(ints, mat[x][y])
			x++
		}
		sort.Ints(ints)
		y = j
		for x := 0; x < len(mat) && y < len(mat[0]); y++ {
			mat[x][y] = ints[x]
			x++
		}
	}
	return mat

}
