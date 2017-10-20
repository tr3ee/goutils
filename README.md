# goutils

## Introduction

[goutils]() is the Utilities I often use when writing the Go language code

## Using this package

```sh
$ git clone https://github.com/tr3ee/goutils.git tr3e/utils
```

## Quick Test

create a go file named `hello-world.go`
```GO
package main

import (
    "fmt"
    "tr3e/utils/cli-banner"
)

func main() {
    banner.Autoload("hello-world","1.0","this is an quick test file")
    fmt.Println("hello world")
}
```

run hello-world.go

```sh
$ go run hello-world.go
```
if it works properly, Congrats! you can now do something awesome.

## License

MIT License

Copyright (c) 2017 tr3e

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
