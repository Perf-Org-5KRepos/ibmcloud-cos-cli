// Code generated by "stringer -type=ObjectGetErrorCause"; DO NOT EDIT.

package errors

import "strconv"

const _ObjectGetErrorCause_name = "IsDirOpening"

var _ObjectGetErrorCause_index = [...]uint8{0, 5, 12}

func (i ObjectGetErrorCause) String() string {
	i -= 1
	if i < 0 || i >= ObjectGetErrorCause(len(_ObjectGetErrorCause_index)-1) {
		return "ObjectGetErrorCause(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ObjectGetErrorCause_name[_ObjectGetErrorCause_index[i]:_ObjectGetErrorCause_index[i+1]]
}
