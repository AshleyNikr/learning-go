package popcount

func PopBit(x uint64) int {
	count := 0
	for i := uint(0); i < 64; i++ {
		if x&1 != 0 {
			count++
		}
		x = x >> 1
	}
	return count
}
