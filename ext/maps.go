package ext

// CloneMap returns a new map with the same values as the original
//
//goland:noinspection GoUnusedExportedFunction
func CloneMap[Key int | string, Value any](this map[Key]Value) map[Key]Value {
	nm := make(map[Key]Value)
	for i, v := range this {
		nm[i] = v
	}
	return nm
}

// MapToArray returns an array of each of the values in the map
//
//goland:noinspection GoUnusedExportedFunction
func MapToArray[TKey int | int8 | int32 | int16 | int64 | string, TValue any](this map[TKey]TValue) []TValue {
	res := make([]TValue, 0)
	for _, v := range this {
		res = append(res, v)
	}
	return res
}

// MapKeyExists returns true and the value if the 'key' exists in the map
//
//goland:noinspection GoUnusedExportedFunction
func MapKeyExists[TKey int | int8 | int32 | int16 | int64 | string, TValue any](this map[TKey]TValue, key TKey, def TValue) (bool, TValue) {
	if _, ok := this[key]; ok {
		// fmt.Printf("MapKeyExists, %v\n", ok)
		return true, this[key]
	}
	return false, def
}

// MapAddOrUpdate updates the 'key' with the 'value' if it already exists, else it will add the value to the map
//
//goland:noinspection GoUnusedExportedFunction
func MapAddOrUpdate[TKey int | int8 | int32 | int16 | int64 | string, TValue any](this map[TKey]TValue, key TKey, val TValue) map[TKey]TValue {
	this[key] = val
	return this
}

// MapArrayAdd adds 'value' to the array of the 'key' in the map
//
//goland:noinspection GoUnusedExportedFunction
func MapArrayAdd[TKey int | int8 | int32 | int16 | int64 | string, TValue any](this map[TKey][]TValue, key TKey, val TValue) map[TKey][]TValue {
	if _, ok := this[key]; !ok {
		this[key] = []TValue{val}
		return this
	}
	this[key] = append(this[key], val)
	return this
}
