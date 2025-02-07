package ext

import (
	"fmt"
	"strconv"
)

type FuncSelector[TItem any, TOutput any] func(TItem) TOutput

//goland:noinspection GoUnusedExportedFunction
func ConvertToArray[TItem any, TOutput any](arr []TItem, selector FuncSelector[TItem, TOutput]) []TOutput {
	res := make([]TOutput, len(arr))
	for index, elm := range arr {
		val := selector(elm)
		res[index] = val
	}
	return res
}

//goland:noinspection GoUnusedExportedFunction
func ArrayToArrayGroups[TItem any, TOutput any](arr []TItem, count int, selector FuncSelector[TItem, TOutput]) [][]TOutput {
	groups := make([][]TOutput, 0)
	group := make([]TOutput, count)
	groupCount := 0

	for _, elm := range arr {
		if groupCount == count {
			groups = append(groups, group)
			group = make([]TOutput, count)
			groupCount = 0
		}
		val := selector(elm)
		group[groupCount] = val
		groupCount++
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups
}

//goland:noinspection GoUnusedExportedFunction
func ArrayToMapList[TItem any, TKey int | int8 | int32 | int16 | int64 | string, TOutput any](arr []TItem, keySelector FuncSelector[TItem, TKey], valueSelector FuncSelector[TItem, TOutput]) map[TKey][]TOutput {
	res := make(map[TKey][]TOutput)

	for _, elm := range arr {
		selKey := keySelector(elm)
		selVal := valueSelector(elm)

		_, ok := res[selKey]
		if !ok {
			res[selKey] = make([]TOutput, 0)
		}
		res[selKey] = append(res[selKey], selVal)
	}

	return res
}

//goland:noinspection GoUnusedExportedFunction
func ArrayToMap[TItem any, TKey int | int8 | int32 | int16 | int64 | string, TOutput any](arr []TItem, keySelector FuncSelector[TItem, TKey], valueSelector FuncSelector[TItem, TOutput]) map[TKey]TOutput {
	res := make(map[TKey]TOutput)
	for _, elm := range arr {
		res[keySelector(elm)] = valueSelector(elm)
	}
	return res
}

//goland:noinspection GoUnusedExportedFunction
func ArrayToMapIgnore[TItem any, TKey int | int8 | int32 | int16 | int64 | string, TOutput any](arr []TItem, keySelector FuncSelector[TItem, TKey], valueSelector FuncSelector[TItem, TOutput], ignoreKeyMatcher FuncSelector[TItem, bool]) map[TKey]TOutput {
	res := make(map[TKey]TOutput)
	for _, elm := range arr {
		if ignoreKeyMatcher(elm) {
			continue
		}
		res[keySelector(elm)] = valueSelector(elm)
	}
	return res
}

//goland:noinspection GoUnusedExportedFunction
func ArrayAddRange[TItem any](list []TItem, items []TItem) []TItem {
	for _, i := range items {
		list = append(list, i)
	}
	return list
}

//goland:noinspection GoUnusedExportedFunction
func ArraySplit[TItem any](list []TItem, max int) map[int][]TItem {
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

//goland:noinspection GoUnusedExportedFunction
func ArrayUnique[TKey int | int8 | int32 | int16 | int64 | string, TValue any](arr []TValue, keySelector func(TValue) TKey) []TValue {
	occurred := make(map[TKey]bool)
	result := make([]TValue, 0)
	for e := range arr {
		key := keySelector(arr[e])
		// check if already the mapped and variable are set to true or not
		if _, ok := occurred[key]; !ok {
			occurred[key] = true
			// Append to result slice.
			result = append(result, arr[e])
		}
	}
	return result
}

//goland:noinspection GoUnusedExportedFunction
func ArrayFilter[TInput any, TResult any](items []TInput, test func(TInput) (TResult, bool)) []TResult {
	result := make([]TResult, 0)
	for _, item := range items {
		if val, ok := test(item); ok {
			result = append(result, val)
		}
	}
	return result
}

//goland:noinspection GoUnusedExportedFunction
func ArrayAdd[TValue any](arr []TValue, val TValue) []TValue {
	return append(arr, val)
}

// ArrayRemove removes the index from the array and keeps the order of elements
//
//goland:noinspection GoUnusedExportedFunction
func ArrayRemove[TValue any](arr []TValue, index int) []TValue {
	return append(arr[:index], arr[index+1:]...)
}

// ArrayRemoveFast removes the index from the array, but the order of elements are not kept
//
//goland:noinspection GoUnusedExportedFunction
func ArrayRemoveFast[TValue any](s []TValue, index int) []TValue {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

//goland:noinspection GoUnusedExportedFunction
func ArrayCut[TValue any](arr []TValue, startIndex int, endIndex int) []TValue {
	return append(arr[:startIndex], arr[endIndex:]...)
}

//goland:noinspection GoUnusedExportedFunction
func ArrayRemovePointer[TValue *any](arr []TValue, index int) []TValue {
	copy(arr[index:], arr[index+1:])
	arr[len(arr)-1] = nil // or the zero value of T
	arr = arr[:len(arr)-1]
	return arr
}

//goland:noinspection GoUnusedExportedFunction
func ArrayCutPointer[TValue *any](arr []TValue, startIndex int, endIndex int) []TValue {
	copy(arr[startIndex:], arr[endIndex:])
	for k, n := len(arr)-endIndex+startIndex, len(arr); k < n; k++ {
		arr[k] = nil // or the zero value of T
	}
	arr = arr[:len(arr)-endIndex+startIndex]
	return arr
}

//goland:noinspection GoUnusedExportedFunction
func ArrayPop[TValue any](arr []TValue) ([]TValue, TValue) {
	x, a := arr[0], arr[1:]
	return a, x
}

//goland:noinspection GoUnusedExportedFunction
func ArrayReverse[TValue any](arr []TValue) []TValue {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}
	return arr
}

//goland:noinspection GoUnusedExportedFunction
func ArrayExists[TValue any](arr []TValue, matcher func(TValue) bool) bool {
	for _, i := range arr {
		if matcher(i) {
			return true
		}
	}
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ArrayOfStringToInt(sArr []string) []int {
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

//goland:noinspection GoUnusedExportedFunction
func ArraySurroundWith(arr []string, char string) []string {
	var resArr = make([]string, 0)
	for _, i := range arr {
		resArr = append(resArr, fmt.Sprintf("%s%s%s", char, i, char))
	}
	return resArr
}
