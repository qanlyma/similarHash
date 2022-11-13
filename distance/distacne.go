package distance

import "math/bits"

func MinHamming(testhash uint64, hashset []uint64) int {
	if len(hashset) == 0 {
		return 0
	}
	mindistance := 64

	for _, hash := range hashset {
		//distance := 100
		hamming := hash ^ testhash
		distance := popcnt(hamming)
		//fmt.Println(testhash, "   ", hash, "    ", distance)

		mindistance = min(mindistance, distance)
	}
	return mindistance
}

func popcnt(x uint64) int { return bits.OnesCount64(x) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
