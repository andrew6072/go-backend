package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, x := range nums {
		m[x] = i
	}
	for i, x := range nums {
		value, exists := m[target-x]
		if i != value && exists {
			return []int{i, value}
		}
	}
	return nil
}
