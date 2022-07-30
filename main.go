// cspell:words okcancel yesno yesnocancel systemmodal idcancel idyes idno

package Popup

import (
	"syscall"
	"unsafe"
)

const (
	MB_OK          = 0x00000000
	MB_OKCANCEL    = 0x00000001
	MB_YESNO       = 0x00000004
	MB_YESNOCANCEL = 0x00000003
	// this allows to focus the created popup instead another window
	MB_SYSTEMMODAL = 0x00001000
)

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
var user32DLL = syscall.NewLazyDLL("user32.dll")
var procMessageBox = user32DLL.NewProc("MessageBoxW") // Return value: Type int

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw#return-value
const (
	IDCANCEL = 2
	IDYES    = 6
	IDNO     = 7
)

// Create an alert with title and message. Closes it with close button or Ok press. Returns void.
func Alert(title string, message string) {
	lpCaption, _ := syscall.UTF16PtrFromString(title)
	lpText, _ := syscall.UTF16PtrFromString(message)
	procMessageBox.Call(uintptr(0x00),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(MB_OK|MB_SYSTEMMODAL))
}

// Create a "lazy" alert (without default focus) with title and message. Closes it with close button or Ok press. Returns void.
func LazyAlert(title string, message string) {
	lpCaption, _ := syscall.UTF16PtrFromString(title)
	lpText, _ := syscall.UTF16PtrFromString(message)
	procMessageBox.Call(uintptr(0x00),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(MB_OK))
}

// Create a dialog that closes choosing Ok or No, returns true if Ok and false if No. Can't be closed via close button.
func Dialog(title string, message string) bool {
	lpCaption, _ := syscall.UTF16PtrFromString(title)
	lpText, _ := syscall.UTF16PtrFromString(message)
	responseValue, _, _ := procMessageBox.Call(uintptr(0x00),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(MB_YESNO|MB_SYSTEMMODAL))

	// true if Yes, false if No
	return responseValue == IDYES
}

// Create a "lazy" dialog (without default focus) that closes choosing Ok or No, returns true if Ok and false if No. Can't be closed via close button.
func LazyDialog(title string, message string) bool {
	lpCaption, _ := syscall.UTF16PtrFromString(title)
	lpText, _ := syscall.UTF16PtrFromString(message)
	responseValue, _, _ := procMessageBox.Call(uintptr(0x00),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(MB_YESNO))

	// true if Yes, false if No
	return responseValue == IDYES
}
