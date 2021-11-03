package msgp

// The sizes provided
// are the worst-case
// encoded sizes for
// each type. For variable-
// length types ([]byte, string),
// the total encoded size is
// the prefix size plus the
// length of the object.
const (
	Int64Size      = 9
	IntSize        = Int64Size
	Uint64Size     = Int64Size
	Float64Size    = 9
	Float32Size    = 5
	Complex64Size  = 10
	Complex128Size = 18

	TimeSize = 15

	BytesPrefixSize = 5
)
