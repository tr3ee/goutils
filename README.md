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

this project is released under [the MIT license](https://github.com/tr3ee/goutils/blob/master/LICENSE)