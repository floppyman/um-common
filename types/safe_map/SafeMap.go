package safe_map

//goland:noinspection GoUnusedExportedFunction
func NewSafeMap() SafeMap {
	return SafeMap{Store: make(map[string]interface{})}
}

type SafeMap struct {
	Store map[string]any
}

func (re SafeMap) Get(key string) any {
	return re.Store[key]
}

func (re SafeMap) Set(key string, val any) {
	re.Store[key] = val
}

func (re SafeMap) Remove(key string) {
	delete(re.Store, key)
}

func (re SafeMap) Contains(key string) bool {
	_, ok := re.Store[key]
	return ok
}
