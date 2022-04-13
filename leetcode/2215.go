func findDifference(nums1 []int, nums2 []int) [][]int {

	m1 := make(map[int]int);
	m2 := make(map[int]int);
	var answer [][]int
	var answer0 []int
	var answer1 []int

	for _, val := range nums1 {
		m1[val] = 1
	}
	for _, val := range nums2 {
		m2[val] = 1
	}

	for key, _ := range m1 {
		if _, ok := m2[key]; !ok {
			answer0 = append(answer0, key)
		}
	}
	for key, _ := range m2 {
		if _, ok := m1[key]; !ok {
			answer1 = append(answer1, key)
		}
	}
	answer = append(answer, answer0)
	answer = append(answer, answer1)
	return answer
}