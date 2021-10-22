package winapi_type

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type HANDLE uintptr
type HMODULE uintptr
type LPVOID uintptr
type WORD uint16
type DWORD uint32
type DWORD_PTR uintptr // 不确定
type ULong uint32
type HRESULT uint32 // https://blog.csdn.net/ixsea/article/details/7272909
type LPBYTE *byte
type LPDWORD *uint32
type LPWSTR *uint16
type LMSTR *uint16

// type WCHAR = wchar_t
// type wchar_t = uint16

// UTF16toString converts a pointer to a UTF16 string into a Go string.
func UTF16toString(p *uint16) string {
	return syscall.UTF16ToString((*[4096]uint16)(unsafe.Pointer(p))[:])
}

func StrPtr(s string) uintptr {
	systemNameUTF16String, err := syscall.UTF16PtrFromString(s)
	if err == nil {
		// 这里转换的时候出错,则不继续执行 OR 赋值用本地的
		tmp := utf16.Encode([]rune("\x00"))
		systemNameUTF16String = &tmp[0]
	}

	return uintptr(unsafe.Pointer(systemNameUTF16String))
}

func CharsPtr(s string) uintptr {
	bPtr, err := syscall.BytePtrFromString(s)
	if err != nil {
		return uintptr(0) // 这么写肯定不太对 @TODO
	}
	return uintptr(unsafe.Pointer(bPtr))
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}
