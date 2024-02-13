package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSampleStats$
func TestSampleStats(t *testing.T) {
	for _, tc := range []struct {
		count      []int
		statistics []float64
	}{
		{
			count:      []int{0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			statistics: []float64{1.00000, 3.00000, 2.37500, 2.50000, 3.00000},
		},
		{
			count:      []int{0, 4, 3, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			statistics: []float64{1.00000, 4.00000, 2.1818181818181817, 2.00000, 1.00000},
		},
		{
			count:      []int{264, 912, 1416, 1903, 2515, 3080, 3598, 4099, 4757, 5270, 5748, 6451, 7074, 7367, 7847, 8653, 9318, 9601, 10481, 10787, 11563, 11869, 12278, 12939, 13737, 13909, 14621, 15264, 15833, 16562, 17135, 17491, 17982, 18731, 19127, 19579, 20524, 20941, 21347, 21800, 22342, 21567, 21063, 20683, 20204, 19818, 19351, 18971, 18496, 17974, 17677, 17034, 16701, 16223, 15671, 15167, 14718, 14552, 14061, 13448, 13199, 12539, 12265, 11912, 10931, 10947, 10516, 10177, 9582, 9102, 8699, 8091, 7864, 7330, 6915, 6492, 6013, 5513, 5140, 4701, 4111, 3725, 3321, 2947, 2357, 1988, 1535, 1124, 664, 206, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			statistics: []float64{0.00000, 89.00000, 42.826829, 42.00000, 40.00000},
		},
		{
			count:      []int{3504389, 165872, 3254354, 2382237, 3894994, 1543673, 3601959, 3360481, 2250927, 2851284, 2947090, 1779123, 260198, 212182, 3017153, 1284944, 2190801, 401287, 1920021, 1703434, 1205084, 3205642, 2674765, 3106625, 3167967, 1455650, 3368920, 2692039, 1181893, 2948239, 1660711, 3094668, 567515, 3775522, 1039368, 2283031, 3503447, 3755429, 3305177, 2268857, 3223278, 1993842, 3539984, 3869704, 1648441, 2565056, 1680059, 475978, 3449295, 3791341, 557654, 464406, 1391144, 3690228, 1142104, 2036721, 3066753, 3901671, 2619935, 1363446, 1409659, 3098974, 2107121, 332213, 2607616, 1403649, 1332423, 3798188, 1910436, 839122, 2815151, 4272, 1519268, 3140686, 1498862, 1382223, 1192301, 2034578, 2294722, 1518981, 1703643, 325862, 2494360, 1081354, 319389, 3031924, 1346686, 2722481, 1950744, 2096754, 2576991, 2326085, 2285739, 3806849, 1639354, 2802274, 3609362, 3152924, 75585, 1601883, 3590204, 2308693, 1276820, 2495766, 2960019, 2601223, 85520, 1782113, 2914644, 408357, 2776257, 682333, 3157360, 3148127, 1763907, 2487015, 105685, 777846, 3449894, 1955365, 2738762, 522010, 2962745, 2272673, 973485, 2412368, 1267064, 1490548, 842846, 1042437, 1807571, 2441072, 2981390, 1863242, 3720573, 629057, 2371556, 445186, 728437, 2112521, 3269653, 3318709, 299351, 459562, 675723, 3030255, 1147665, 1473484, 1011994, 738245, 3209204, 484421, 2009360, 3647712, 786402, 1859993, 1398030, 1819265, 1806565, 3607579, 163565, 1248474, 3259370, 2566820, 2433535, 2260583, 2712490, 1775470, 3397822, 2368342, 3589282, 2885514, 778401, 3729125, 2176775, 3533311, 3400615, 3057796, 2030669, 1458083, 3456717, 210671, 1014571, 2396817, 1883746, 2008262, 3641023, 3452500, 985086, 3317463, 765037, 2889971, 2801322, 3584987, 1223184, 3029068, 1298373, 3136128, 1750290, 3895010, 237932, 2730051, 3351724, 1485867, 712706, 2431367, 2209671, 1641282, 1689124, 1814332, 882359, 3363267, 3664679, 781595, 270431, 2520852, 1889952, 734963, 2241739, 2655887, 2119965, 2176304, 1695741, 735568, 252795, 1709312, 1661007, 2937669, 1660138, 1027780, 2458361, 2282011, 3745776, 119191, 1374332, 424207, 3614549, 2939366, 2584670, 2386823, 3117681, 1492822, 497427, 2046946, 1238796, 2972332, 527591, 3122102, 1891922, 3634296, 3886432, 3525414, 819099, 1961272, 150057, 3237702},
			statistics: []float64{0.00000, 255.00000, 125.49992539222552, 125.00000, 57.00000},
		},
		{
			count:      []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			statistics: []float64{66.00000, 124.00000, 86.75000, 78.50000, 66.00000},
		},
		{
			count:      []int{1110334, 2565423, 1381946, 3724624, 3760200, 528838, 1432290, 485973, 2341953, 410279, 3749046, 3602779, 490515, 3031556, 569097, 3267507, 2916268, 1095409, 27962, 1352652, 1790260, 1606817, 1991310, 582528, 423586, 1645670, 2202174, 1704485, 1819887, 1545977, 1762186, 3508264, 3888077, 2283285, 2554434, 3559466, 2054745, 3782219, 1820542, 1165412, 2685924, 3288012, 682160, 2574520, 512338, 329960, 453183, 3587259, 16751, 3384979, 2271587, 1196758, 3522370, 116500, 3474035, 2563261, 3493953, 1532291, 2617291, 1950003, 1560606, 143602, 3178964, 3643530, 2519930, 293750, 2173974, 3516247, 524341, 2379639, 1130999, 3748853, 2481883, 2260677, 917366, 2248720, 1961975, 1466102, 3223569, 2797794, 1277625, 2513196, 2183524, 2798935, 2855340, 1485570, 1166846, 610273, 2571762, 1561958, 3580685, 3021496, 2368626, 2996339, 1490718, 1950089, 1550742, 1186279, 2553557, 2540036, 999446, 2370946, 1406297, 2654376, 1783148, 1799926, 437456, 805455, 3749971, 1276815, 882626, 330879, 127485, 612652, 1846084, 505502, 292827, 59468, 2798440, 2098847, 353366, 1981580, 112258, 1325422, 701722, 1944604, 2453758, 1486133, 404897, 3082873, 368702, 2034500, 120397, 1387714, 2628666, 3893596, 2499882, 1041532, 3102036, 1055816, 670482, 2303950, 3879696, 453280, 1569263, 2730677, 3101683, 2562986, 494120, 930299, 2020650, 1720596, 3135172, 1093698, 966412, 19392, 2550951, 3576749, 191159, 1367674, 2327108, 2479869, 3289786, 1904683, 544242, 512872, 1791909, 2048044, 1015525, 3253376, 2257082, 331976, 2776845, 2920440, 2508327, 515084, 1001388, 3537902, 3348638, 3702541, 626417, 1830584, 1226214, 1864351, 3709336, 300465, 2088192, 2982053, 1878838, 1674936, 1129544, 2249506, 3813402, 3174388, 1103518, 3322125, 695021, 2204732, 2034486, 2322163, 264423, 3790129, 905572, 3792595, 3838897, 1284055, 3107253, 705843, 3499459, 3563249, 2589759, 644934, 2003109, 480792, 1227703, 358653, 1987309, 3765731, 535983, 2650477, 2854262, 191341, 3534573, 2605596, 1260958, 3728041, 554097, 815807, 3522111, 3607382, 1776086, 3335203, 1269522, 487583, 1999285, 1016251, 797391, 3391007, 1017812, 526082, 22539, 476415, 708537, 1758103, 3303593, 2089135, 3055736, 1538287, 2762321, 20559, 1809152, 2002236, 2076051, 3533401, 905767, 3440983},
			statistics: []float64{0.00000, 255.00000, 127.06772268478379, 129.00000, 135.00000},
		},
		{
			count:      []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			statistics: []float64{191.00000, 204.00000, 200.40000, 204.00000, 204.00000},
		},
	} {
		statistics := sampleStats(tc.count)
		assert.Equal(t, tc.statistics, statistics)
	}
}