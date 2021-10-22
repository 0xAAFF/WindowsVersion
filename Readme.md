# gowinver

## 支持版本
```go
/*
	Windows NT 4
	Windows 95
	Windows 98
	Windows Me
	Windows 2000
	Windows XP
	Windows XP 64
	Windows Server 2003
	Windows Server 2003 R2
	Windwos Vista
	Windows Server 2008
	Windwos 7
	Windows Server 2008 R2
	Windwos 8
	Windows Server 2012
	Windows 8.1
	Windows Server 2012 R2
	Windows 10
	Windows Server 2016
	Windows Server 2019
	Windows 11
	Windows 11 + ()
*/
```

## 原理
使用windows API得到windwos的主版本号,次版本号,编译版本号,这样大概区分出windows的版本
然后根据各个版本号下的具体的属性再次细分得到是否为Server或者WorkStation

Windows8.1以下版本的API:
GetVersionExW()

Windwos8.1及以上的版本API:
RtlGetNtVersionNumbers() 和 GetVersionExW() 配合使用
注:为什么这样使用:
因为GetVersionExW()在windows8.1及以上获取的主版本号,次版本号和编译版本号都是一样的 :6.2.9200
但是获取的其他信息是对的.
所以,使用RtlGetNtVersionNumbers()这个API会得到正确的版本号.

当然,中间为了获取winxp64的,使用了GetNativeSystemInfo这个API来获取cpu的架构:ARCH=386可以执行 amd64获取的数据不对.请考虑使用

## Use API

Under Windows 8.1 Use API `GetVersionExW()`

Windows 8 + Use API `RtlGetNtVersionNumbers()` and `GetVersionExW()`

and Other APIs

## UseAge

OSVersion()

##  注意

 `GetNativeSystemInfo()` 
 May not Working When `go env -w GOARCH=amd64`
 But Working When `go env -w GOARCH=386`

 ```go
    /*

    When `go env -w GOARCH=amd64`: NumberOfProcessors:0x10000,ProcessorType:0x7e050006,(Error.不是太清楚这个什么原因):
	&version.SystemInfo{DummyUnionName:0x100000000009, PageSize:0x10000, MinimumApplicationAddress:0x7ffffffeffff, MaximumApplicationAddress:0xff, ActiveProcessorMask:0x21d800000008, NumberOfProcessors:0x10000, ProcessorType:0x7e050006, AllocationGranularity:0x0, ProcessorLevel:0x0, ProcessorRevision:0x0}


	When `go env -w GOARCH=386`  : NumberOfProcessors:0x8, ProcessorType:0x21d8 (correct 这个才是正确的):
	&version.SystemInfo{DummyUnionName:0x9, PageSize:0x1000, MinimumApplicationAddress:0x10000, MaximumApplicationAddress:0xfffeffff, ActiveProcessorMask:0xff, NumberOfProcessors:0x8, ProcessorType:0x21d8, AllocationGranularity:0x10000, ProcessorLevel:0x6, ProcessorRevision:0x7e05}

    */

 ```

## End

如果有疑问或者代码出现失误,请您及时联系我修改


