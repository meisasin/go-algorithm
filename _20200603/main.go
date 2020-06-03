package main

func main() {

}

func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W+1)
	for i := K; i <= N && i < K+W; i++ {
		dp[i] = 1.0
	}

	dp[K-1] = 1.0 * float64(min(N-K+1, W)) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+W+1]-dp[i+1])/float64(W)
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
