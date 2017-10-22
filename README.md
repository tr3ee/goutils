# goutils

## Introduction

[goutils]() is the Utilities I often use when writing the Go language code

## Using this package

```sh
$ git clone https://github.com/tr3ee/goutils.git tr3e/utils
```

Once it's done,run the following command to fix dependencies

```sh
$ go get tr3e/utils/trans/layers
$ go get tr3e/utils/
```

## Quick Test

create a go file named `hello-world.go`
```GO
package main

import (
    "fmt"
    "tr3e/utils/cli-banner"
    _ "tr3e/utils/semaphore"
    _ "tr3e/utils/cipher"
    _ "tr3e/utils/trans/layers"
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

Otherwise,run the command `go get` under the folder of `hello-world.go`

## License

this project is released under [the MIT license](https://github.com/tr3ee/goutils/blob/master/LICENSE)