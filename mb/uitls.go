package mb

import (
	"math"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type (
	// char
	Char byte
	// singed char
	SChar int8
	// unsigned char
	UChar uint8
	// short
	Short int16
	// unsigned short
	UShort uint16
	// int
	Int int32
	// unsigned int
	UInt uint32
	// long
	Long int32
	// unsigned long
	ULong uint32
	// long int
	LongInt int64
	// unsigned long int
	ULongInt uint64
	// long long
	LongLong int64
	// unsigned long long
	ULongLong uint64
	// float
	Float float32
	// double
	Double float64
	// long double
	LongDouble float64
	// wchar_t
	Wchar_t uint16
	// utf8
	Utf8 byte
)

// alias
type (
	any = interface{}
)

func StrToPtr(s string) uintptr {
	p, err := syscall.BytePtrFromString(s)
	if err != nil {
		return 0
	}

	return (uintptr)(unsafe.Pointer(p))
}

func StrFromPtr(p uintptr) string {
	if p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := p; *(*byte)(unsafe.Pointer(ptr)) != 0; n++ {
		ptr++
	}

	return string(unsafe.Slice((*byte)(unsafe.Pointer(p)), n))
}

func StrToChar(s string) []Char {
	bytes, err := syscall.ByteSliceFromString(s)
	if err != nil {
		return nil
	}
	return *((*[]Char)(unsafe.Pointer(&bytes)))
}

func StrToWPtr(s string) uintptr {
	p, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return 0
	}
	return uintptr(unsafe.Pointer(p))
}

func StrToWChar(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

func StrFromWPtr(p uintptr) string {

	if p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := p; *(*uint16)(unsafe.Pointer(ptr)) != 0; n++ {
		ptr += unsafe.Sizeof(uint16(0))
	}

	return syscall.UTF16ToString(
		unsafe.Slice(
			(*uint16)(unsafe.Pointer(p)),
			n,
		),
	)
}

func F32ToPtr(f float32) uintptr {
	return uintptr(math.Float32bits(f))
}

func F32FromPtr(p uintptr) float32 {
	return math.Float32frombits(uint32(p))
}

func F64ToPtr(f float64) uintptr {
	return uintptr(math.Float64bits(f))
}

func F64FromPtr(p uintptr) float64 {
	return math.Float64frombits(uint64(p))
}

func BoolToPtr(b bool) uintptr {

	if b {
		return 1
	}

	return 0
}

func BoolFromPtr(p uintptr) bool {
	return p != 0
}

func LOWORD(dwValue uint32) uint16 {
	return uint16(dwValue & 0xFFFF)
}

func CallbackToPtr(callback interface{}) uintptr {
	return syscall.NewCallback(callback)
}
