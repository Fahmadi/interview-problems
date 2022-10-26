package main

func main() {
	nums := []int{2, 3, 2, 3, 4}
	singleNumber(nums)
}

func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		if hash[num] == 0 {
			hash[num] = 1
		} else {
			hash[num] = hash[num] + 1
		}
	}
	var sNumber int
	for h := range hash {
		if hash[h] == 1 {
			sNumber = h
		}
	}

	return sNumber

}
