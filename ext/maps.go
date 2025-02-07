package ext

// MapKeyExists checks if the map contains the 'key' and returns true plus the value in map, else it returns false plus the 'def' value
//
//goland:noinspection GoUnusedExportedFunction
func MapKeyExists[TKey int | int8 | int32 | int16 | int64 | string, TValue any](this map[TKey]TValue, key TKey, def TValue) (bool, TValue) {
	if val, ok := this[key]; ok {
		return true, val
	}
	return false, def
}
