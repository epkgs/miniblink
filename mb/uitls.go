package mb

import (
	"math"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
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
	p, err := windows.BytePtrFromString(s)
	if err != nil {
		return 0
	}

	return (uintptr)(unsafe.Pointer(p))
}

func StrFromPtr(p uintptr) string {
	return windows.BytePtrToString(AssertType[uint8](p))
}

func StrToChar(s string) []Char {
	bytes, err := windows.ByteSliceFromString(s)
	if err != nil {
		return nil
	}
	return *((*[]Char)(unsafe.Pointer(&bytes)))
}

func StrToWPtr(s string) uintptr {
	p, err := windows.UTF16PtrFromString(s)
	if err != nil {
		return 0
	}
	return uintptr(unsafe.Pointer(p))
}

func StrToWChar(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

func StrFromWPtr(p uintptr) string {
	return windows.UTF16PtrToString(AssertType[uint16](p))
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

func AssertType[T interface{}](ptr uintptr) *T {
	return (*T)(unsafe.Pointer(ptr))
}

func LOWORD(dwValue uint32) uint16 {
	return uint16(dwValue & 0xFFFF)
}

func StrCopy(src uintptr, n int) string {
	return string(BytesCopy(src, n))
}

func BytesCopy(src uintptr, n int) []byte {
	if n == 0 {
		return make([]byte, 0)
	}

	byts := make([]uint8, n)
	for i := 0; i < n; i++ {
		byts[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}

	return byts
}

func CallbackToPtr(callback interface{}) uintptr {
	return windows.NewCallback(callback)
}
