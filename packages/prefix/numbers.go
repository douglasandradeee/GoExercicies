package prefix

import "strconv"

// Capital - represents the number of the telephone prefix in a state/province.
var (
	Capital        = 71
	test           = "test"
	TestWithPrefix = test + " " + strconv.Itoa(Capital)
)
