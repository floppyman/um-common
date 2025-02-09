package ext

// Ifi is an "if" that is inline, since go does not provide one in its syntax
//
//goland:noinspection GoUnusedExportedFunction
func Ifi[TVal any](x bool, t TVal, f TVal) TVal {
	if x {
		return t
	}
	return f
}

// Iif is an inline "if", since go does not provide one in its syntax
//
//goland:noinspection GoUnusedExportedFunction
func Iif[TVal any](x bool, t TVal, f TVal) TVal {
	return Ifi(x, t, f)
}
