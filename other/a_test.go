package other

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	val := "aaaaaaaa"
	ans := val[0:1]
	fmt.Println(ans)

	fmt.Println(strings.ReplaceAll("abba", "a", "ccc"))
}

func TestCopy(t *testing.T) {

	compute1 := Compute{
		name:  "Compute1",
		money: Money{18.00},
	}

	fmt.Println(compute1)
	fmt.Println("----------")
	compute2 := compute1
	compute2.name = "Compute2"
	compute2.money.price = 68.99
	fmt.Println(compute1.money.price)
	fmt.Println(compute2.money.price)

}

type Money struct {
	price float64
}
type Compute struct {
	name  string
	money Money
}

func TestNum(t *testing.T) {
	fmt.Println('8')

	fmt.Println('2'+'9' > '9'+'1')
	fmt.Println(string('7'))

	fmt.Println(90)
	ans := string([]byte{56})
	fmt.Println(ans)

	fmt.Println(strconv.Itoa(0))

	count := '1' - '0' + '0' - '0'
	fmt.Println(count)

	fmt.Println('0')
}

func addBinary(a string, b string) string {

	if len(a) > len(b) {
		a, b = b, a
	}

	ans := ""
	jin := false
	for i := len(a) - 1; i >= 0; i-- {
		count := a[i] - '0' + b[i] - '0'
		aa := a[i] - 48
		bb := '9' - 48
		cc := a[i]

		if jin {
			count += 1
		}

		if count >= 2 {
			jin = true
			count = count - 2
		} else {
			jin = false
		}
		fmt.Println("Count is ", count)
		//ans += strconv.Itoa(count)
		fmt.Println("Ans is ", ans)

	}

	for i := len(b) - 1; i >= len(a); i-- {
		count := b[i] - '0'
		if jin {
			count += 1
		}

		if count >= 2 {
			jin = true
			count = count - 2
		} else {
			jin = false
		}
		fmt.Println("Count is ", count)

		//ans += strconv.Itoa(count)
		fmt.Println("Ans is ", ans)

	}

	if jin {
		ans += "1"
	}

	return ans
}
