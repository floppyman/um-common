package ext

import (
	"fmt"
	"strconv"
)

type SliceFuncType[TItem any, TOutput any] func(TItem) TOutput

// SliceSelect selects each element of 'arr' into another slice where each element can be transformed using 'transformer' function
//
//goland:noinspection GoUnusedExportedFunction
func SliceSelect[TItem any, TValue any](arr []TItem, transformer SliceFuncType[TItem, TValue]) []TValue {
	res := make([]TValue, len(arr))
	for index, elm := range arr {
		val := transformer(elm)
		res[index] = val
	}
	return res
}

// SliceSelectGroups selects each element of 'arr' into another slice where each element can be transformed using 'transformer' function, and it also groups the results in separate slices with the maximum element count defined in 'maxPerGroup'
//
//goland:noinspection GoUnusedExportedFunction
func SliceSelectGroups[TItem any, TValue any](arr []TItem, maxPerGroup int, transformer SliceFuncType[TItem, TValue]) [][]TValue {
	groups := make([][]TValue, 0)
	group := make([]TValue, maxPerGroup)
	groupCount := 0

	for _, elm := range arr {
		if groupCount == maxPerGroup {
			groups = append(groups, group)
			group = make([]TValue, maxPerGroup)
			groupCount = 0
		}
		val := transformer(elm)
		group[groupCount] = val
		groupCount++
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups
}

// SliceToMapList converts the 'arr' slice into a map[TKey][]TValue, using 'keySelector' to define the key, and 'valueTransformer' to transform the value
// If the same key exists more than once, then the last one will be the one that exists in the output map.
//
//goland:noinspection GoUnusedExportedFunction
func SliceToMapList[TItem any, TKey int | int8 | int32 | int16 | int64 | string, TValue any](arr []TItem, keySelector SliceFuncType[TItem, TKey], valueTransformer SliceFuncType[TItem, TValue]) map[TKey][]TValue {
	res := make(map[TKey][]TValue)

	for _, elm := range arr {
		selKey := keySelector(elm)
		selVal := valueTransformer(elm)

		_, ok := res[selKey]
		if !ok {
			res[selKey] = make([]TValue, 0)
		}
		res[selKey] = append(res[selKey], selVal)
	}

	return res
}

// SliceToMap converts the 'arr' slice into a map[TKey]TValue, using 'keySelector' to define the key, and 'valueTransformer' to transform the value.
// If the same key exists more than once, then the last one will be the one that exists in the output map.
//
//goland:noinspection GoUnusedExportedFunction
func SliceToMap[TItem any, TKey int | int8 | int32 | int16 | int64 | string, TOutput any](arr []TItem, keySelector SliceFuncType[TItem, TKey], valueTransformer SliceFuncType[TItem, TOutput]) map[TKey]TOutput {
	res := make(map[TKey]TOutput)
	for _, elm := range arr {
		res[keySelector(elm)] = valueTransformer(elm)
	}
	return res
}

// SliceToMapIgnore converts the 'arr' slice into a map[TKey]TValue, using 'keySelector' to define the key, and 'valueTransformer' to transform the value.
// If 'ignoreKeyPredicate' returns true, then that key is ignored, and not added to the map.
// If the same key exists more than once, then the last one will be the one that exists in the output map.
//
//goland:noinspection GoUnusedExportedFunction
func SliceToMapIgnore[TItem any, TKey int | int8 | int32 | int16 | int64 | string, TOutput any](arr []TItem, keySelector SliceFuncType[TItem, TKey], valueTransformer SliceFuncType[TItem, TOutput], ignoreKeyPredicate SliceFuncType[TItem, bool]) map[TKey]TOutput {
	res := make(map[TKey]TOutput)
	for _, elm := range arr {
		if ignoreKeyPredicate(elm) {
			continue
		}
		res[keySelector(elm)] = valueTransformer(elm)
	}
	return res
}

// SliceSplit splits the slice into a map[int][]TItem where the map will contain the number of groups needed based on the 'max' per group
//
//goland:noinspection GoUnusedExportedFunction
func SliceSplit[TItem any](list []TItem, max int) map[int][]TItem {
	res := make(map[int][]TItem)
	pos := 0
	index := 0
	for index < len(list) {
		indexMax := index + max
		if indexMax > len(list) {
			indexMax = len(list)
		}
		res[pos] = list[index:indexMax]
		pos++
		index += max
	}
	return res
}

// SliceWhere returns a new slice, where the 'predicate' function determines if the value should be included or not
//
//goland:noinspection GoUnusedExportedFunction
func SliceWhere[TInput any, TResult any](items []TInput, predicate func(TInput) (TResult, bool)) []TResult {
	result := make([]TResult, 0)
	for _, item := range items {
		if val, ok := predicate(item); ok {
			result = append(result, val)
		}
	}
	return result
}

// SliceRemove removes the index from the array and keeps the order of elements
//
//goland:noinspection GoUnusedExportedFunction
func SliceRemove[TValue any](arr []TValue, index int) []TValue {
	return append(arr[:index], arr[index+1:]...)
}

// SliceRemoveFast removes the index from the array, but the order of elements are not kept
//
//goland:noinspection GoUnusedExportedFunction
func SliceRemoveFast[TValue any](s []TValue, index int) []TValue {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

// SliceCut returns a "part" of the original slice based on the 'startIndex' and 'endIndex'.
// Out-Of-Bounds are not checked.
//
//goland:noinspection GoUnusedExportedFunction
func SliceCut[TValue any](arr []TValue, startIndex int, endIndex int) []TValue {
	return append(arr[:startIndex], arr[endIndex:]...)
}

// SlicePop removes the first element of the slice, and returns both the new slice and the value of the first element
//
//goland:noinspection GoUnusedExportedFunction
func SlicePop[TValue any](arr []TValue) ([]TValue, TValue) {
	x, a := arr[0], arr[1:]
	return a, x
}

// SliceReverse reverses the order of all the elements
//
//goland:noinspection GoUnusedExportedFunction
func SliceReverse[TValue any](arr []TValue) []TValue {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}
	return arr
}

// SliceExists checks if any of the elements matches the 'predicate', then returns true
//
//goland:noinspection GoUnusedExportedFunction
func SliceExists[TValue any](arr []TValue, predicate func(TValue) bool) bool {
	for _, i := range arr {
		if predicate(i) {
			return true
		}
	}
	return false
}

// SliceOfStringToInt converts all elements in the slice to a new slice of int
//
//goland:noinspection GoUnusedExportedFunction
func SliceOfStringToInt(sArr []string) []int {
	var iArr = make([]int, 0)

	for _, item := range sArr {
		res, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		iArr = append(iArr, res)
	}

	return iArr
}

// SliceSurroundWith returns a new slice where all the elements have been surrounded with the 'char` value
//
//goland:noinspection GoUnusedExportedFunction
func SliceSurroundWith(arr []string, char string) []string {
	var resArr = make([]string, 0)
	for _, i := range arr {
		resArr = append(resArr, fmt.Sprintf("%s%s%s", char, i, char))
	}
	return resArr
}

// SliceToCommaString returns a string containing all the elements of 'list', 'approxLen' is used for estimation of the []byte slice used for storing the values while generating
//
//goland:noinspection GoUnusedExportedFunction
func SliceToCommaString[TItem int | int32 | int64](list []TItem, approxLen int) string {
	if len(list) == 0 {
		return ""
	}

	// Approx 3 chars per num plus the comma.
	estimate := len(list) * (approxLen + 1)
	b := make([]byte, 0, estimate)

	for _, n := range list {
		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ',')
	}

	b = b[:len(b)-1]
	return string(b)
}
