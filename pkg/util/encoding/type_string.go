// Code generated by "stringer -type=Type"; DO NOT EDIT.

package encoding

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Null-1]
	_ = x[NotNull-2]
	_ = x[Int-3]
	_ = x[Float-4]
	_ = x[Decimal-5]
	_ = x[Bytes-6]
	_ = x[BytesDesc-7]
	_ = x[Time-8]
	_ = x[Duration-9]
	_ = x[True-10]
	_ = x[False-11]
	_ = x[UUID-12]
	_ = x[Array-13]
	_ = x[IPAddr-14]
	_ = x[JSON-15]
	_ = x[Tuple-16]
	_ = x[BitArray-17]
	_ = x[BitArrayDesc-18]
	_ = x[TimeTZ-19]
}

const _Type_name = "UnknownNullNotNullIntFloatDecimalBytesBytesDescTimeDurationTrueFalseUUIDArrayIPAddrJSONTupleBitArrayBitArrayDescTimeTZ"

var _Type_index = [...]uint8{0, 7, 11, 18, 21, 26, 33, 38, 47, 51, 59, 63, 68, 72, 77, 83, 87, 92, 100, 112, 118}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
