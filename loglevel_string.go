// Code generated by "stringer -type=LogLevel"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SILENT-1]
	_ = x[FATAL-2]
	_ = x[ERROR-4]
	_ = x[WARN-8]
	_ = x[INFO-16]
	_ = x[DEBUG-32]
	_ = x[TRACE-64]
}

const (
	_LogLevel_name_0 = "SILENTFATAL"
	_LogLevel_name_1 = "ERROR"
	_LogLevel_name_2 = "WARN"
	_LogLevel_name_3 = "INFO"
	_LogLevel_name_4 = "DEBUG"
	_LogLevel_name_5 = "TRACE"
)

var (
	_LogLevel_index_0 = [...]uint8{0, 6, 11}
)

func (i LogLevel) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _LogLevel_name_0[_LogLevel_index_0[i]:_LogLevel_index_0[i+1]]
	case i == 4:
		return _LogLevel_name_1
	case i == 8:
		return _LogLevel_name_2
	case i == 16:
		return _LogLevel_name_3
	case i == 32:
		return _LogLevel_name_4
	case i == 64:
		return _LogLevel_name_5
	default:
		return "LogLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
