package _023

// 2023-07-11 21:13:43
// 1911. 最大子序列交替和
// 下标从0开始的数组的 交替和 定义为 偶数 下标处元素之 和 减去 奇数 下标处元素之 和
// 比方说，数组 [4,2,5,3] 的交替和为 (4 + 5) - (2 + 3) = 4
// 输入数组 nums ，返回 nums 中 最大子序列交替和

// 动态规划
// 子序列交替和可以理解为偶数下标为正，奇数下标为负，求最大子序列和
// dp[i][0] 表示以 nums[i] 结尾的交替和的最大值，且 nums[i] 必须选取
// dp[i][1] 表示以 nums[i] 结尾的交替和的最大值，且 nums[i] 必须不选取
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	dp := make([][2]int64, n)
	dp[0][0] = int64(nums[0])
	dp[0][1] = 0
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+nums[i])
	for i := 1; i < n; i++ {
		dp[i][0] = max64(dp[i-1][0], dp[i-1][1]+int64(nums[i]))
		dp[i][1] = max64(dp[i-1][1], dp[i-1][0]-int64(nums[i]))
	}
	return max64(dp[n-1][0], dp[n-1][1])
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
