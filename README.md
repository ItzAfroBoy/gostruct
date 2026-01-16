<div align="center">
<pre>
 ____            ____    __                           __      
/\  _`\         /\  _`\ /\ \__                       /\ \__   
\ \ \L\_\    ___\ \,\L\_\ \ ,_\  _ __   __  __    ___\ \ ,_\  
 \ \ \L_L   / __`\/_\__ \\ \ \/ /\`'__\/\ \/\ \  /'___\ \ \/  
  \ \ \/, \/\ \L\ \/\ \L\ \ \ \_\ \ \/ \ \ \_\ \/\ \__/\ \ \_ 
   \ \____/\ \____/\ `\____\ \__\\ \_\  \ \____/\ \____\\ \__\
    \/___/  \/___/  \/_____/\/__/ \/_/   \/___/  \/____/ \/__/
<br>
A go implementation of Python's Struct module
<br>
<img alt="GitHub License" src="https://img.shields.io/github/license/ItzAfroBoy/gostruct"> <img alt="GitHub tag (with filter)" src="https://img.shields.io/github/v/tag/ItzAfroBoy/gostruct?label=version"> <a href="https://www.codefactor.io/repository/github/itzafroboy/gostruct"><img src="https://www.codefactor.io/repository/github/itzafroboy/gostruct/badge" alt="CodeFactor" /></a> <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/ItzAfroBoy/gostruct">
</pre>
</div>

## Installation

```shell
go get github.com/ItzAfroBoy/gostruct@latest
```

## [Example](https://github.com/ItzAfroBoy/gostruct/blob/main/main_test.go)

```go
package main

import (
    "fmt"
    "github.com/ItzAfroBoy/gostruct"
)

type Data struct {
    A int32
    B int32
    C float32
    D float32
    E uint8
    F uint32
}

func main() {
    byteStream := receive() // Data to unpack
    packet := Data{}
    gostruct.UnpackToStruct("<iiffBI", byteStream, &packet)
    fmt.Println(packet)
}
```
