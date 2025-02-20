package util

func CalculateProfit(principalAmount int64, roi float64) int64 {
	profit := float64(principalAmount) * roi / 100
	return int64(profit)
}
