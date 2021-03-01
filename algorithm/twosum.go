package main

func sum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		resolve := target - nums[i]

		if index, ok := m[resolve]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}
