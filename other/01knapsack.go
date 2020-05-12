package other

import "fmt"

func Knapsack01() {

	var N, V int

	n, _ := fmt.Scanf("%d %d", &N, &V)
	fmt.Println(n)

	fmt.Println(N)
	fmt.Println(V)

	v := make([]int, N)
	w := make([]int, N)
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d %d", &v[i], &w[i])
	}

	fmt.Println(v)
	fmt.Println(w)
	res := make([]int, V)
	for i := 1; i <= N; i++ {
		for j := V; j >= v[i]; j-- {
			if res[j-v[i]]+w[i] > res[j] {
				res[j] = res[j-v[i]] + w[i]
			}
		}
	}

	fmt.Println(res[V])
}
