package version

import (
	"syscall"
)

// GetSystemMetrics function (winuser.h)
// User32.dll
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
/*
int GetSystemMetrics(
  [in] int nIndex
);
*/

func GetSystemMetrics(index int) (int, error) {
	user32 := syscall.NewLazyDLL("User32.dll")
	procGetSystemMetrics := user32.NewProc("GetSystemMetrics")
	v, _, err := procGetSystemMetrics.Call(uintptr(index))
	// fmt.Printf("%#v \n", v)
	// fmt.Println(err)
	return int(v), err

}
