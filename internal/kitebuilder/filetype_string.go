// Code generated by "stringer -type=FileType"; DO NOT EDIT.

package kitebuilder

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UNKNOWN-0]
	_ = x[TXT-1]
	_ = x[JSON-2]
	_ = x[KITE-3]
}

const _FileType_name = "UNKNOWNTXTJSONKITE"

var _FileType_index = [...]uint8{0, 7, 10, 14, 18}

func (i FileType) String() string {
	if i < 0 || i >= FileType(len(_FileType_index)-1) {
		return "FileType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FileType_name[_FileType_index[i]:_FileType_index[i+1]]
}