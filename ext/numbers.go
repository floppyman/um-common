package ext

import (
	"fmt"
	"strconv"
)

// SplitFloat32 tries to split a 'num' into its two components the number and precision, or an error if it fails
//
//goland:noinspection GoUnusedExportedFunction
func SplitFloat32(num float32) (int32, int32, error) {
	return SplitFloat32WithPrecision(num, -1)
}

// SplitFloat32WithPrecision tries to split a 'num' into its two components the number and precision, 'prec' can be used to define the precision (number of decimals), or an error if it fails
//
//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func SplitFloat32WithPrecision(num float32, prec int) (int32, int32, error) {
	var numSec int32
	var numNSec int32
	if _, err := fmt.Sscanf(strconv.FormatFloat(float64(num), 'f', prec, 32), "%d.%d", &numSec, &numNSec); err != nil {
		return 0, 0, err
	}
	return numSec, numNSec, nil
}

// SplitFloat64 tries to split a 'num' into its two components the number and precision, or an error if it fails
//
//goland:noinspection GoUnusedExportedFunction
func SplitFloat64(num float64) (int64, int64, error) {
	return SplitFloat64WithPrecision(num, -1)
}

// SplitFloat64WithPrecision tries to split a 'num' into its two components the number and precision, 'prec' can be used to define the precision (number of decimals), or an error if it fails
//
//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func SplitFloat64WithPrecision(num float64, prec int) (int64, int64, error) {
	var numSec int64
	var numNSec int64
	if _, err := fmt.Sscanf(strconv.FormatFloat(num, 'f', prec, 64), "%d.%d", &numSec, &numNSec); err != nil {
		return 0, 0, err
	}
	return numSec, numNSec, nil
}

// IntToString converts 'num' into a base(10) string of the number
//
//goland:noinspection GoUnusedExportedFunction
func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

// Int32ToString converts 'num' into a base(10) string of the number
//
//goland:noinspection GoUnusedExportedFunction
func Int32ToString(num int32) string {
	return strconv.FormatInt(int64(num), 10)
}

// Int64ToString converts 'num' into a base(10) string of the number
//
//goland:noinspection GoUnusedExportedFunction
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Float32ToString converts 'num' into a base(10) string of the number
//
//goland:noinspection GoUnusedExportedFunction
func Float32ToString(num float32) string {
	return strconv.FormatFloat(float64(num), 'f', -1, 32)
}

// Float64ToString converts 'num' into a base(10) string of the number
//
//goland:noinspection GoUnusedExportedFunction
func Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
