package queues

/*
Algorithm:
For the first character in the string to move to the back
and eliminate the first "opposing" character in the string

DDRRR - the first D moves to the back and takes out the first R
DRRD - the first D moves to the back and takes out the first R
RDD - the first R moves to the back and takes out the first D
DR - the first (and only) D moves to the back and takes out the first (and only) R
D - D wins the vote.
*/

// https://leetcode.com/problems/dota2-senate/

func predictPartyVictory(senate string) string {
	var length = len(senate)
	var radiants = make([]int, 0, length)
	var dires = make([]int, 0, length)
	var skipR, skipD = 0, 0
	for index, r := range senate {
		switch r {
		case 'R':
			if skipR > 0 {
				skipR--
			} else {
				radiants = append(radiants, index)
				skipD++
			}
		case 'D':
			if skipD > 0 {
				skipD--
			} else {
				dires = append(dires, index)
				skipR++
			}
		}
	}
	if skipR > len(radiants) {
		return "Dire"
	}
	if skipD > len(dires) {
		return "Radiant"
	}
	radiants = radiants[skipR:]
	dires = dires[skipD:]
	for index := length; len(radiants) > 0 && len(dires) > 0; index++ {
		if radiants[0] < dires[0] {
			dires = dires[1:]
			radiants = radiants[1:]
			radiants = append(radiants, index)
		} else {
			radiants = radiants[1:]
			dires = dires[1:]
			dires = append(dires, index)
		}
		index++
	}
	if len(radiants) > 0 {
		return "Radiant"
	}
	return "Dire"
}

// func predictPartyVictory(senate string) string {
// 	var radiants = list.New()
// 	var dires = list.New()
// 	for index, r := range senate {
// 		switch r {
// 		case 'R':
// 			radiants.PushBack(index)
// 		case 'D':
// 			dires.PushBack(index)
// 		}
// 	}
// 	for index := len(senate); radiants.Len() > 0 && dires.Len() > 0; index++ {
// 		if radiants.Front().Value.(int) < dires.Front().Value.(int) {
// 			dires.Remove(dires.Front())
// 			radiants.Remove(radiants.Front())
// 			radiants.PushBack(index)
// 		} else {
// 			radiants.Remove(radiants.Front())
// 			dires.Remove(dires.Front())
// 			dires.PushBack(index)
// 		}
// 	}
// 	if radiants.Len() > 0 {
// 		return "Radiant"
// 	}
// 	return "Dire"
// }
