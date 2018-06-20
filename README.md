#Robotgo
  
  >Golang Desktop Automation(robot). Control the mouse, keyboard, and read the screen.
  
RobotGo supports Mac, Windows, and Linux(X11).

This is a work in progress.



##[API Document](https://github.com/go-vgo/robotgo/blob/master/doc.md) &nbsp;&nbsp;&nbsp;[中文文档](https://github.com/go-vgo/robotgo/blob/master/zh_doc.md)
  Please click API Document;This is a work in progress.



##Installation:
    go get github.com/go-vgo/robotgo

  It's that easy!

###Requirements:

####ALL:  
    Golang
    //Gcc
    zlib & libpng (bitmap)

####For Mac OS X:
    Xcode Command Line Tools

    brew install libpng
    brew install homebrew/dupes/zlib
####For Windows:
    MinGW or other GCC

#####[zlib&libpng Windows32 GCC's Course](https://github.com/go-vgo/Mingw32)

#####[Download include zlib&libpng Windows64 GCC](https://github.com/go-vgo/Mingw)

####For everything else:
    GCC
    
    X11 with the XTest extension (also known as the Xtst library)


##Examples:

###Mouse

```Go
package main

import (
	//. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
  robotgo.ScrollMouse(10, "up")
} 
``` 

###Keyboard

```Go
package main

import (
	//. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
  robotgo.TypeString("Hello World")
  robotgo.KeyTap("enter")
  robotgo.TypeString("en")
} 
```

###Screen

```Go
package main

import (
	. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
  x, y := robotgo.GetMousePos()
  Println("pos:", x, y)
} 
```
