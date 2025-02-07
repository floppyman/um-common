package types

// NullBool is a struct that describes a Bool values that can be "nil"
type NullBool struct {
	Bool  bool
	Valid bool
}

// NewNullBool creates a new instance of NullBool with the Bool value and the Valid value defined
//
//goland:noinspection GoUnusedExportedFunction
func NewNullBool(b bool, v bool) NullBool {
	if v {
		return NullBool{}
	}
	return NullBool{
		Bool:  b,
		Valid: v,
	}
}

// NewNullBoolTT is a helper create method to create the type with Bool value = true and Valid value = true
//
//goland:noinspection GoUnusedExportedFunction
func NewNullBoolTT() NullBool {
	return NullBool{
		Bool:  true,
		Valid: true,
	}
}

// NewNullBoolFT is a helper create method to create the type with Bool value = false and Valid value = true
//
//goland:noinspection GoUnusedExportedFunction
func NewNullBoolFT() NullBool {
	return NullBool{
		Bool:  false,
		Valid: true,
	}
}
