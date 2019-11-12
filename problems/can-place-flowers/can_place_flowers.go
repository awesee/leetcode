package problem605

func canPlaceFlowers(flowerbed []int, n int) bool {
	sum, ec := 0, 1
	for _, v := range flowerbed {
		if v == 0 {
			ec++
		} else {
			sum += (ec - 1) / 2
			ec = 0
		}
	}
	sum += ec / 2
	return sum >= n
}
