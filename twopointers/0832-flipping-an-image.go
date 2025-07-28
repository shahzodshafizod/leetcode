package twopointers

// https://leetcode.com/problems/flipping-an-image/

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)

	var l, r int

	for row := 0; row < n; row++ {
		for l, r = 0, n-1; l <= r; l, r = l+1, r-1 {
			image[row][l], image[row][r] = image[row][r]^1, image[row][l]^1
		}
	}

	return image
}
