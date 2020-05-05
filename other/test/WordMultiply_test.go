package test

import (
	"fmt"
	"go-algorithm/other"
	"testing"
)

func TestMaxProduct(t *testing.T) {

	res := other.MaxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"})
	fmt.Println(res)

	res = other.MaxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"})
	fmt.Println(res)

	res = other.MaxProduct([]string{"a", "aa", "aaa", "aaaa"})
	fmt.Println(res)

	res = other.MaxProduct([]string{"edadc", "ebbfe", "aacdde", "dfe", "cb", "fddddff", "fabca", "adccac", "ece", "ccaf", "feba", "bcb", "edadc", "aea", "bacb", "acefa", "fcebffd", "dfeebca", "bedcbaa", "deaccc", "abedc", "dadff", "eef", "ddebbb", "abecab", "cd", "abdeee", "eedce", "deef", "dceaddd", "ced", "fbbf", "ba", "eefeda", "fb", "cddc", "adafbf", "dded", "aadbf", "caefbaf", "ccebf", "dbb", "ee", "dadcecc", "ddbcabb", "afeaa", "ec", "aad", "efde", "cbcda", "cdbdafd", "cddc", "ecaaa", "ae", "cfc", "bccac", "cdcc", "abbbf", "fcdface", "ddbcdc", "bfebb", "daed", "bc", "dc", "ecdfc", "eeb", "bb", "dad", "caecb", "fbe", "bbbc", "cacea", "dbc", "fbe", "bcfffbd", "aeda", "cff", "ddfc", "ea", "bdfd", "ccb", "cb", "ae", "ceabefa", "dcea", "cbaed", "bfedf", "fa", "ccd", "fece", "bceef", "acabca", "dafa", "fdeec", "dac", "cae", "adeeadb", "ecacc", "acfe", "de"})
	fmt.Println(res)

}

// BenchmarkMaxProduct-12    	1000000000	         0.000039 ns/op
// BenchmarkMaxProduct-12    	1000000000	         0.000033 ns/op
func BenchmarkMaxProduct(b *testing.B) {
	res := other.MaxProduct([]string{"edadc", "ebbfe", "aacdde", "dfe", "cb", "fddddff", "fabca", "adccac", "ece", "ccaf", "feba", "bcb", "edadc", "aea", "bacb", "acefa", "fcebffd", "dfeebca", "bedcbaa", "deaccc", "abedc", "dadff", "eef", "ddebbb", "abecab", "cd", "abdeee", "eedce", "deef", "dceaddd", "ced", "fbbf", "ba", "eefeda", "fb", "cddc", "adafbf", "dded", "aadbf", "caefbaf", "ccebf", "dbb", "ee", "dadcecc", "ddbcabb", "afeaa", "ec", "aad", "efde", "cbcda", "cdbdafd", "cddc", "ecaaa", "ae", "cfc", "bccac", "cdcc", "abbbf", "fcdface", "ddbcdc", "bfebb", "daed", "bc", "dc", "ecdfc", "eeb", "bb", "dad", "caecb", "fbe", "bbbc", "cacea", "dbc", "fbe", "bcfffbd", "aeda", "cff", "ddfc", "ea", "bdfd", "ccb", "cb", "ae", "ceabefa", "dcea", "cbaed", "bfedf", "fa", "ccd", "fece", "bceef", "acabca", "dafa", "fdeec", "dac", "cae", "adeeadb", "ecacc", "acfe", "de"})
	fmt.Println(res)
}

// BenchmarkMaxProductShort-12    	1000000000	         0.000009 ns/op
func BenchmarkMaxProductShort(b *testing.B) {
	res := other.MaxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"})
	fmt.Println(res)
}
