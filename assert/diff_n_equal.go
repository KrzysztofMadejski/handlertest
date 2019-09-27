package assert

import (
	"reflect"
)

type Differ func(minusPrefixed, plusPrefixed interface{}) string
type Equalizer func(x, y interface{}) bool

// TODO global vars might not be the best idea, but how else allow to set it without the need to set it in every test?
var differ Differ
var equalizer = Equalizer(func(x, y interface{}) bool {
	// default equalizer
	return reflect.DeepEqual(x, y)
})

// SetDiff set diffing function likely using dependencies
func SetDiff(d Differ) {
	differ = d
}

// SetEqual set equality checking function likely using dependencies
func SetEqual(e Equalizer) {
	equalizer = e
}

// Implementations using dependencies

// 1. github.com/google/go-cmp/cmp

//import (
//	"github.com/google/go-cmp/cmp"
//	"github.com/google/go-cmp/cmp/cmpopts"
//)

//// GoCmpDiffer using https://godoc.org/github.com/google/go-cmp/cmp
//func GoCmpDiffer(minusPrefixed, plusPrefixed interface{}) string {
//	return cmp.Diff(minusPrefixed, plusPrefixed)
//}
//
//// GoCmpDifferShortStringsAlso ...
//// By default cmp.Diff does not break into lines strings shorter than 64 chars and 4 lines
//// This implementation will keep it the same for all strings
////
//// Ref https://github.com/google/go-cmp/blob/6eaffb0bbd93e7b7eae6cb6de180383b034db0d2/cmp/report_slices.go#L55-L56 and #L100
//func GoCmpDifferShortStringsAlso(minusPrefixed, plusPrefixed interface{}) string {
//	return cmp.Diff(minusPrefixed, plusPrefixed,
//		cmpopts.AcyclicTransformer("SplitString", func(s string) []string { return strings.Split(s, "\n") }))
//}
//
//func GoCmpEqualizer(x, y interface{}) bool {
//	return cmp.Equal(x, y)
//}

// 2. https://godoc.org/gotest.tools/assert#Equal

// TODO research for diffing https://godoc.org/gotest.tools/assert#Equal
