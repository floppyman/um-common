package safe_map

// NewSafeMap create a new instance of SafeMap with String key
//
//goland:noinspection GoUnusedExportedFunction
func NewSafeMap() SafeMap {
	return SafeMap{Store: make(map[string]interface{})}
}

type SafeMap struct {
	Store map[string]any
}

// Get a key from the SafeMap
func (re SafeMap) Get(key string) any {
	return re.Store[key]
}

// Set a key and value into the SafeMap
func (re SafeMap) Set(key string, val any) {
	re.Store[key] = val
}

// Remove a key from the SafeMap
func (re SafeMap) Remove(key string) {
	delete(re.Store, key)
}

// Contains check if the key exists in the SafeMap
func (re SafeMap) Contains(key string) bool {
	_, ok := re.Store[key]
	return ok
}
