package operator

import (
	"strconv"

	"github.com/douglasandradeee/GoExercicies/fundamentos/packages/prefix"
)

// OperatorName - Define the name of the operator.
// OperatorPrefixCapital - Define the prefix number of the operator
var (
	OperatorName          = "XPTO Telecom"
	OperatorPrefixCapital = strconv.Itoa(prefix.Capital) + " " + OperatorName
)
