package astra

// AcceptedTypes is a list of all accepted types for the astra package.
// Everything else is considered a type that has to be processed.
var AcceptedTypes = []string{
	"nil",
	"string",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float",
	"float32",
	"float64",
	"bool",
	"byte",
	"rune",
	"struct",
	"map",
	"slice",
	"any",
	"file", // not an official type, but we use it to replace for the binary format
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return Index(s, v) >= 0
}

func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func IsAcceptedType(t string) bool {
	return Contains(AcceptedTypes, t)
}
