### 官方题解 [@link](https://leetcode-cn.com/problems/jewels-and-stones/solution/bao-shi-yu-shi-tou-by-leetcode-solution/)

![1.png](./source/1.png)
```Golang
func numJewelsInStones(J string, S string) (sum int) {
    for _, s := range S {
        for _, j := range J {
            if s == j {
                sum++
                break
            }
        }
    }
    return
}
```
![2.png](./source/2.png)
![3.png](./source/3.png)
```Golang
func numJewelsInStones(J string, S string) (sum int) {
    jewelsSet := map[byte]bool{}
    for i := range J {
        jewelsSet[J[i]] = true
    }
    for i := range S {
        if jewelsSet[S[i]] {
            sum++
        }
    }
    return
}
```
![4.png](./source/4.png)