package safe_map

//goland:noinspection GoUnusedExportedFunction
func NewSafeMapInt() SafeMapInt {
	return SafeMapInt{Store: make(map[int]interface{})}
}

type SafeMapInt struct {
	Store map[int]any
}

func (re SafeMapInt) Get(key int) any {
	return re.Store[key]
}

func (re SafeMapInt) Set(key int, val any) {
	re.Store[key] = val
}

func (re SafeMapInt) Remove(key int) {
	delete(re.Store, key)
}

func (re SafeMapInt) Contains(key int) bool {
	_, ok := re.Store[key]
	return ok
}
