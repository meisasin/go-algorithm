package main

func main() {

}

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
