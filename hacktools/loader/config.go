package loader

// It is temporarily recommended to copy the parts you need form this library into you own code
const (
	MEM_COMMIT     = 0x1000
	MEM_RESERVE    = 0x2000
	MEM_RESET      = 0x80000
	MEM_RESET_UNDO = 0x1000000

	MEM_TOP_DOWN    = 0x100000
	MEM_WRITE_WATCH = 0x200000
	MEM_PHYSICAL    = 0x400000
	MEM_LARGE_PAGES = 0x20000000

	PAGE_EXECUTE_READWRITE = 0x40
)

/*
LPVOID VirtualAlloc(
	[in, optional] LPVOID lpAddress
	[in]           SIZE_T dwSize
	[in]           DWORD  flAllocationType
	[in]           DWORD  flProtect
	);
*/

/*
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
*/
