package robotgo

/*
//#if defined(IS_MACOSX)
	#cgo darwin CFLAGS: -x objective-c  -Wno-deprecated-declarations -I/usr/local/opt/libpng/include -I/usr/local/opt/zlib/include
	#cgo darwin LDFLAGS: -framework Cocoa -framework OpenGL -framework IOKit -framework Carbon -framework CoreFoundation -L/usr/local/opt/libpng/lib -lpng -L/usr/local/opt/zlib/lib -lz
//#elif defined(USE_X11)
	#cgo linux CFLAGS:-I/usr/src
	#cgo linux LDFLAGS:-L/usr/src -lpng -lz -lX11 -lXtst -lm
//#endif
	#cgo windows LDFLAGS: -lgdi32 -luser32 -lpng -lz
//#include <AppKit/NSEvent.h>
#include "screen/goScreen.h"
#include "mouse/goMouse.h"
#include "key/goKey.h"
#include "bitmap/goBitmap.h"
//#include "window/goWindow.h"
//#include "event/goEvent.h"
*/
import "C"

import (
	. "fmt"
	"unsafe"
	// "runtime"
	// "syscall"
)

/*
      _______.  ______ .______       _______  _______ .__   __.
    /       | /      ||   _  \     |   ____||   ____||  \ |  |
   |   (----`|  ,----'|  |_)  |    |  |__   |  |__   |   \|  |
    \   \    |  |     |      /     |   __|  |   __|  |  . `  |
.----)   |   |  `----.|  |\  \----.|  |____ |  |____ |  |\   |
|_______/     \______|| _| `._____||_______||_______||__| \__|
*/

type Bit_map struct {
	ImageBuffer   *C.uint8_t
	Width         C.size_t
	Height        C.size_t
	Bytewidth     C.size_t
	BitsPerPixel  C.uint8_t
	BytesPerPixel C.uint8_t
}

func GetPixelColor(x, y C.size_t) string {
	color := C.aGetPixelColor(x, y)
	gcolor := C.GoString(color)
	defer C.free(unsafe.Pointer(color))
	return gcolor
}

func GetScreenSize() (C.size_t, C.size_t) {
	size := C.aGetScreenSize()
	// Println("...", size, size.width)
	return size.width, size.height
}

func GetXDisplayName() string {
	name := C.aGetXDisplayName()
	gname := C.GoString(name)
	defer C.free(unsafe.Pointer(name))
	return gname
}

func SetXDisplayName(name string) string {
	cname := C.CString(name)
	str := C.aSetXDisplayName(cname)
	gstr := C.GoString(str)
	return gstr
}

func CaptureScreen(x, y, w, h C.int) Bit_map {
	bit := C.aCaptureScreen(x, y, w, h)
	// Println("...", bit)
	bit_map := Bit_map{
		ImageBuffer:   bit.imageBuffer,
		Width:         bit.width,
		Height:        bit.height,
		Bytewidth:     bit.bytewidth,
		BitsPerPixel:  bit.bitsPerPixel,
		BytesPerPixel: bit.bytesPerPixel,
	}

	return bit_map
}

/*
 __  __
|  \/  | ___  _   _ ___  ___
| |\/| |/ _ \| | | / __|/ _ \
| |  | | (_) | |_| \__ \  __/
|_|  |_|\___/ \__,_|___/\___|

*/

type MPoint struct {
	x int
	y int
}

//C.size_t  int
func MoveMouse(x, y C.size_t) {
	C.aMoveMouse(x, y)
}

func DragMouse(x, y C.size_t) {
	C.aDragMouse(x, y)
}

func MoveMouseSmooth(x, y C.size_t) {
	C.aMoveMouseSmooth(x, y)
}

func GetMousePos() (C.size_t, C.size_t) {
	pos := C.aGetMousePos()
	// Println("pos:###", pos, pos.x, pos.y)
	return pos.x, pos.y
}

func MouseClick() {
	C.aMouseClick()
}

func MouseToggle() {
	C.aMouseToggle()
}

func SetMouseDelay(x C.size_t) {
	C.aSetMouseDelay(x)
}

func ScrollMouse(x C.size_t, y string) {
	z := C.CString(y)
	C.aScrollMouse(x, z)
	defer C.free(unsafe.Pointer(z))
}

/*
 _  __          _                         _
| |/ /___ _   _| |__   ___   __ _ _ __ __| |
| ' // _ \ | | | '_ \ / _ \ / _` | '__/ _` |
| . \  __/ |_| | |_) | (_) | (_| | | | (_| |
|_|\_\___|\__, |_.__/ \___/ \__,_|_|  \__,_|
		  |___/
*/
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func KeyTap(args ...string) {
	var apara string
	Try(func() {
		apara = args[1]
	}, func(e interface{}) {
		// Println("err:::", e)
		apara = "null"
	})

	zkey := C.CString(args[0])
	amod := C.CString(apara)
	// defer func() {
	C.aKeyTap(zkey, amod)
	// }()

	defer C.free(unsafe.Pointer(zkey))
	defer C.free(unsafe.Pointer(amod))
}

func KeyToggle(args ...string) {
	var apara string
	Try(func() {
		apara = args[1]
	}, func(e interface{}) {
		// Println("err:::", e)
		apara = "null"
	})

	zkey := C.CString(args[0])
	amod := C.CString(apara)
	// defer func() {
	str := C.aKeyToggle(zkey, amod)
	Println(str)
	// }()
	defer C.free(unsafe.Pointer(zkey))
	defer C.free(unsafe.Pointer(amod))
}

func TypeString(x string) {
	cx := C.CString(x)
	C.aTypeString(cx)
	defer C.free(unsafe.Pointer(cx))
}

func TypeStringDelayed(x string, y C.size_t) {
	cx := C.CString(x)
	C.aTypeStringDelayed(cx, y)
	defer C.free(unsafe.Pointer(cx))
}

func SetKeyboardDelay(x C.size_t) {
	C.aSetKeyboardDelay(x)
}

/*
.______    __  .___________..___  ___.      ___      .______
|   _  \  |  | |           ||   \/   |     /   \     |   _  \
|  |_)  | |  | `---|  |----`|  \  /  |    /  ^  \    |  |_)  |
|   _  <  |  |     |  |     |  |\/|  |   /  /_\  \   |   ___/
|  |_)  | |  |     |  |     |  |  |  |  /  _____  \  |  |
|______/  |__|     |__|     |__|  |__| /__/     \__\ | _|
*/

func OpenBitmap(gpath string) C.MMBitmapRef {
	path := C.CString(gpath)
	bit := C.aOpenBitmap(path)
	Println("opening...", bit)
	return bit
	// defer C.free(unsafe.Pointer(path))
}

func SaveBitmap(args ...interface{}) {
	var mtype C.MMImageType
	Try(func() {
		mtype = args[2].(C.MMImageType)
	}, func(e interface{}) {
		Println("err:::", e)
		mtype = 1
	})

	path := C.CString(args[1].(string))
	savebit := C.aSaveBitmap(args[0].(C.MMBitmapRef), path, mtype)
	Println("opening...", savebit)
	// return bit
	// defer C.free(unsafe.Pointer(path))
}

func TostringBitmap(bit C.MMBitmapRef) *C.char {
	str_bit := C.aTostringBitmap(bit)
	// Println("...", str_bit)
	return str_bit
}

/*
____    __    ____  __  .__   __.  _______   ______   ____    __    ____
\   \  /  \  /   / |  | |  \ |  | |       \ /  __  \  \   \  /  \  /   /
 \   \/    \/   /  |  | |   \|  | |  .--.  |  |  |  |  \   \/    \/   /
  \            /   |  | |  . `  | |  |  |  |  |  |  |   \            /
   \    /\    /    |  | |  |\   | |  '--'  |  `--'  |    \    /\    /
    \__/  \__/     |__| |__| \__| |_______/ \______/      \__/  \__/

*/

/*
------------ ---    ---  ------------ ----    ---- ------------
************ ***    ***  ************ *****   **** ************
----         ---    ---  ----         ------  ---- ------------
************ ***    ***  ************ ************     ****
------------ ---    ---  ------------ ------------     ----
****          ********   ****         ****  ******     ****
------------   ------    ------------ ----   -----     ----
************    ****     ************ ****    ****     ****

*/
