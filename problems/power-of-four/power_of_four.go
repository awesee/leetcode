package problem342

func isPowerOfFour(num int) bool {
	return num&(num-1) == 0 && num&0x55555555 > 0
}
