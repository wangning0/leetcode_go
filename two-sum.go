func twoSum(nums []int, target int) []int {
    for index, num := range nums {
        for _index, _num := range nums[index + 1:] {
            if nums[index] + nums[_index] == target {
                return []int{index, _index}
            }
        }
    }
}
