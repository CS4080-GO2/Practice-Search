package main

import (
	"fmt"
	"time"
)

// Maybe sort 'comparable' objects instead?
func main() {
	test := []int{0, 1, 6, 8, 9, 10, 12, 14, 15, 16, 18, 19, 21, 25, 26, 29, 31, 32, 34, 36, 37, 38, 39, 41, 44, 46, 51, 52, 56, 57, 61, 63, 66, 68, 69, 72, 73, 74, 75, 76, 77, 81, 87, 88, 93, 95, 98, 99, 100, 101, 102, 103, 104, 105, 107, 108, 109, 110, 112, 113, 114, 115, 116, 118, 119, 120, 121, 129, 130, 135, 136, 137, 140, 142, 143, 144, 146, 149, 153, 155, 157, 158, 159, 160, 161, 162, 164, 167, 168, 170, 173, 174, 179, 180, 181, 184, 185, 191, 195, 196, 197, 198, 199, 200, 202, 203, 206, 207, 208, 209, 210, 212, 223, 226, 228, 230, 231, 232, 233, 234, 238, 239, 241, 242, 244, 245, 247, 249, 250, 252, 253, 255, 258, 259, 260, 261, 263, 264, 266, 268, 269, 275, 280, 283, 284, 285, 286, 287, 290, 291, 292, 296, 297, 299, 300, 301, 302, 307, 309, 310, 312, 313, 315, 317, 320, 325, 327, 330, 331, 335, 339, 340, 342, 345, 349, 351, 352, 353, 356, 357, 361, 362, 363, 365, 368, 369, 370, 371, 374, 376, 377, 378, 382, 383, 384, 387, 391, 393, 394, 396, 398, 399, 402, 403, 404, 407, 415, 418, 421, 424, 426, 427, 428, 429, 430, 434, 435, 438, 439, 440, 441, 442, 444, 446, 448, 449, 450, 454, 456, 457, 462, 463, 471, 472, 473, 474, 476, 480, 481, 489, 493, 495, 496, 497, 498, 499, 501, 503, 504, 505, 508, 509, 511, 512, 515, 516, 518, 524, 531, 532, 534, 536, 537, 538, 539, 540, 541, 542, 544, 546, 555, 557, 558, 559, 560, 561, 562, 563, 564, 565, 573, 574, 576, 577, 579, 580, 581, 582, 583, 584, 586, 587, 588, 589, 590, 595, 598, 599, 602, 604, 605, 606, 607, 608, 609, 610, 615, 616, 617, 624, 628, 630, 631, 634, 636, 640, 641, 642, 643, 646, 647, 651, 652, 653, 655, 656, 657, 660, 661, 663, 665, 668, 669, 670, 671, 675, 679, 680, 682, 686, 688, 702, 703, 707, 708, 709, 710, 711, 712, 713, 715, 716, 717, 718, 720, 722, 723, 726, 727, 728, 729, 732, 733, 734, 736, 737, 739, 744, 746, 753, 754, 760, 762, 766, 767, 771, 772, 773, 774, 775, 777, 778, 779, 780, 785, 787, 788, 791, 792, 795, 797, 799, 800, 801, 802, 804, 805, 806, 808, 809, 810, 811, 812, 814, 817, 819, 821, 822, 826, 827, 829, 830, 831, 832, 835, 838, 840, 841, 842, 843, 844, 846, 848, 849, 850, 854, 855, 856, 857, 862, 863, 865, 866, 867, 868, 869, 871, 872, 873, 874, 875, 876, 877, 880, 882, 883, 892, 894, 896, 898, 899, 901, 904, 905, 906, 907, 908, 909, 910, 913, 914, 916, 918, 921, 922, 924, 926, 927, 930, 933, 934, 937, 939, 940, 942, 944, 945, 946, 947, 948, 950, 952, 953, 955, 960, 961, 963, 964, 966, 973, 974, 979, 981, 982, 983, 987, 992, 995, 997, 999, 1000}

	fmt.Print("Binary Search:\n")
	start := time.Now()
	result := fullTest(test)
	end := time.Now()
	fmt.Printf("worked?: %v took: %d ms\n", result, (end.Nanosecond()/1000000 - start.Nanosecond()/1000000))

	fmt.Print("Exponential Search:\n")

	newStart := time.Now()
	result = fullExpoTest(test)
	newEnd := time.Now()
	fmt.Printf("worked?: %v took: %d ms", result, (newEnd.Nanosecond()/1000000 - newStart.Nanosecond()/1000000))
}

func fullTest(arr []int) bool {
	size := len(arr)
	for i, v := range arr {
		fmt.Println(i)
		if binarySearch(arr, v, 0, size) != i {
			return false
		}
	}
	return true
}

func fullExpoTest(arr []int) bool {
	for i, v := range arr {
		fmt.Println(i)
		if expoSearch(arr, v) != i {
			return false
		}
	}
	return true
}

func binarySearch(arr []int, element int, min int, max int) int {
	found := false
	for !found {
		middle := ((max - min) / 2) + min
		eval := arr[middle] - element
		if eval == 0 {
			// Found element
			found = true
			return middle
		} else if eval > 0 {
			// arr[max/2] > element, so go lower in array
			max = middle - 1
		} else {
			// arr[max/2] < element, so go higher in array
			min = middle + 1
		}
		//fmt.Printf("%d %d %d\n", min, max, middle)
	}
	return -1
}

func expoSearch(arr []int, element int) int {
	//return the first index if it's the target element
	if arr[0] == element {
		return 0;
	}
	
	i := 1
	for i < len(arr) && arr[i] <= element {
		i = i*2
	}

	min := i/2
	max := i
	//ensure that i <= len(arr)
	if len(arr) < i{
		max = len(arr)-1
	}

	return binarySearch(arr, element, min, max)

}
