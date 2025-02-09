package safe_map

// NewSafeMapInt create a new instance of SafeMapInt with int key
//
//goland:noinspection GoUnusedExportedFunction
func NewSafeMapInt() SafeMapInt {
	return SafeMapInt{Store: make(map[int]interface{})}
}

type SafeMapInt struct {
	Store map[int]any
}

// Get a key from the SafeMapInt
func (re SafeMapInt) Get(key int) any {
	return re.Store[key]
}

// Set a key and value into the SafeMapInt
func (re SafeMapInt) Set(key int, val any) {
	re.Store[key] = val
}

// Remove a key from the SafeMapInt
func (re SafeMapInt) Remove(key int) {
	delete(re.Store, key)
}

// Contains check if the key exists in the SafeMapInt
func (re SafeMapInt) Contains(key int) bool {
	_, ok := re.Store[key]
	return ok
}
