package maths

// https://leetcode.com/problems/alice-and-bob-playing-flower-game/

func flowerGame(n int, m int) int64 {
	nev, nod := n/2, (n+1)/2
	mev, mod := m/2, (m+1)/2

	return int64(nev*mod) + int64(nod*mev)
	// return int64(m*n) / 2
}
