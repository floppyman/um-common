package ext

const (
	TrueStr  string = "true"
	FalseStr string = "false"
)

// Iif is an inline "if", since go does not provide one in its syntax
//
//goland:noinspection GoUnusedExportedFunction
func Iif[TVal any](x bool, t TVal, f TVal) TVal {
	if x {
		return t
	}
	return f
}

// BoolToStr will always just return the True / False string value
//
//goland:noinspection GoUnusedExportedFunction
func BoolToStr(x bool) string {
	return Iif(x, TrueStr, FalseStr)
}
