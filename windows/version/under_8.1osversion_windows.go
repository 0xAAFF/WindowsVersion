package version

import (
	"syscall"
	"unsafe"

	. "gowinver/windows/winapi_type"
)

// 获取windows 版本指令
// 参考
// https://www.cnblogs.com/micro-chen/p/5939662.html
//
// https://www.cnblogs.com/lizhigang/p/7161132.html
//
// GetVersionExA 函数 (sysinfoapi.h)
// Kernel32.lib
// 最低支持 Windows 2000 Professional/Server [desktop apps | UWP apps]
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getversionexa

// #region GetVersionExW 函数 (sysinfoapi.h)
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getversionexw
/*
Kernel32.dll
GetVersionExW function (sysinfoapi.h)
NOT_BUILD_WINDOWS_DEPRECATE BOOL GetVersionExW(
  [in, out] LPOSVERSIONINFOW lpVersionInformation
);
*/
func GetVersionExW() (OSVersionInfoExA, error) {

	//csdVer := make([]byte, 128)
	var osVersionInfoExA OSVersionInfoExA
	// osVersionInfoExA := OSVersionInfoExA{
	// 	//CSDVersion: csdVer, //uintptr(unsafe.Pointer(&csdVer[0])), //StrPtr(csdVer), //csdVer[:128], //uintptr(unsafe.Pointer(&csdVer[0])),
	// }
	osVersionInfoExA.OSVersionInfoSize = DWORD(unsafe.Sizeof(osVersionInfoExA))

	kernel32 := syscall.NewLazyDLL("Kernel32.dll")
	procGetVersionExW := kernel32.NewProc("GetVersionExA")
	_, _, err := procGetVersionExW.Call(uintptr(unsafe.Pointer(&osVersionInfoExA)))
	//v, vv, err := procGetVersionExW.Call(uintptr(unsafe.Pointer(&osVersionInfoExA)))
	// fmt.Println(v)
	// fmt.Println(vv)
	// fmt.Println(err)
	if err != nil && err.Error() == "The operation completed successfully." {
		err = nil
	}

	// fmt.Println("-----------------GetVersionExW-----------------")
	// fmt.Printf("%#v \n", osVersionInfoExA)
	// fmt.Println("主版本        :", osVersionInfoExA.MajorVersion)
	// fmt.Println("次版本        :", osVersionInfoExA.MinorVersion)
	// fmt.Println("编译版本      :", osVersionInfoExA.BuildNumber)
	// fmt.Println("操作系统平台   :", osVersionInfoExA.PlatformId)
	// fmt.Println("SP(主)       :", osVersionInfoExA.MajorVersion)
	// fmt.Println("SP(次)       :", osVersionInfoExA.ServicePackMinor)

	return osVersionInfoExA, err
}

// # endregion

// #region OSVERSIONINFOEXA structure
// https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexa
/*
typedef struct _OSVERSIONINFOEXA {
  DWORD dwOSVersionInfoSize;
  DWORD dwMajorVersion;
  DWORD dwMinorVersion;
  DWORD dwBuildNumber;
  DWORD dwPlatformId;
  CHAR  szCSDVersion[128];
  WORD  wServicePackMajor;
  WORD  wServicePackMinor;
  WORD  wSuiteMask;
  BYTE  wProductType;
  BYTE  wReserved;
} OSVERSIONINFOEXA, *POSVERSIONINFOEXA, *LPOSVERSIONINFOEXA;
*/

type OSVersionInfoExA struct {
	OSVersionInfoSize DWORD     // 结构体大小, in bytes. Set this member to sizeof(OSVERSIONINFOEX).
	MajorVersion      DWORD     // 主版本号
	MinorVersion      DWORD     // 次版本号
	BuildNumber       DWORD     // 编译版本号
	PlatformId        DWORD     // 系统支持的平台
	CSDVersion        [128]byte // 系统补丁包的名称  CSDVersion[128]// 这个128是必须的 不然就会出现"The data area passed to a system call is too small."
	ServicePackMajor  WORD      // 系统补丁包的主版本
	ServicePackMinor  WORD      // 系统补丁包的次版本
	SuiteMask         WORD      // 标识系统上的程序组
	ProductType       byte      // 标识系统类型
	Reserved          byte      // 保留,未使用
} // 这个结构在Windows 2000后出现，老版本的OSVERSIONINFO结构没有wServicePackMajor、wServicePackMinor、wSuiteMask、wProductType和wReserved这几个成员。

// #endregion

// #region  OSVERSIONINFOA
// OSVERSIONINFOA structure (winnt.h)
// https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoa
/*
typedef struct _OSVERSIONINFOA {
  DWORD dwOSVersionInfoSize;
  DWORD dwMajorVersion;
  DWORD dwMinorVersion;
  DWORD dwBuildNumber;
  DWORD dwPlatformId;
  CHAR  szCSDVersion[128];
} OSVERSIONINFOA, *POSVERSIONINFOA, *LPOSVERSIONINFOA;
*/
type OSVersionInfoA struct {
	OSVersionInfoSize DWORD     // 结构体大小, in bytes. Set this member to sizeof(OSVERSIONINFOEX).
	MajorVersion      DWORD     // 主版本号
	MinorVersion      DWORD     // 次版本号
	BuildNumber       DWORD     // 编译版本号
	PlatformId        DWORD     // 系统支持的平台
	CSDVersion        [128]byte // 系统补丁包的名称  CSDVersion[128]
}

// #endregion

// Version Helper functions
// Windwos 系统版本助手(微软官方推荐此函数列表)
// https://docs.microsoft.com/en-us/windows/win32/sysinfo/version-helper-apis

// IsWow64Process函数
// GetNativeSystemInfo function (sysinfoapi.h)
// https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getnativesysteminfo

// Win8一下的系统版本获取
// https://www.cnblogs.com/lizhigang/p/7161132.html
// https://blog.csdn.net/aidem_brown/article/details/52060140?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~aggregatepage~first_rank_ecpm_v1~rank_aggregation-19-52060140.pc_agg_rank_aggregation&utm_term=go+%E8%8E%B7%E5%8F%96%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F%E7%89%88%E6%9C%AC&spm=1000.2123.3001.4430
//
