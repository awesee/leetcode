package problem191

func hammingWeight(num uint32) int {
	num = (num & 0x55555555) + (num>>1)&0x55555555
	num = (num & 0x33333333) + (num>>2)&0x33333333
	num = (num & 0x0F0F0F0F) + (num>>4)&0x0F0F0F0F
	num = (num * 0x01010101) >> 24
	return int(num)
}
