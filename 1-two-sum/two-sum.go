func twoSum(nums []int, target int) []int {
    visitedNumbers := make(map[int]int)

    for i, num := range(nums) {
        diff := target - num
        if idx, exists := visitedNumbers[diff]; exists {
            return []int{idx, i}
        }
        visitedNumbers[num] = i
    }

    return []int{}
}
