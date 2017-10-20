package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"
	"tr3e/trans/layers"
	banner "tr3e/utils/cli-banner"
)

var input, output string
var dFlag bool

func init() {
	flag.StringVar(&input, "i", "", "input data file path")
	flag.StringVar(&output, "o", "", "path to output the file")
	flag.BoolVar(&dFlag, "d", false, "decode the input file if set to true")
}

func main() {
	flag.Parse()
	banner.Autoload("PNG encode/decode", "example", "using this program to encode/decode data")
	if input == "" || output == "" {
		flag.Usage()
		return
	}
	fmt.Println("[*] reading data")
	buf, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] reading data --done-- ")

	// if len(buf) > 1024*1024*1024 {
	// 	panic("to long")
	// }

	fmt.Println("[*] processing data, size=", len(buf))
	t1 := time.Now()
	l := layers.NewPNG(0)
	var data []byte
	if dFlag {
		data, err = l.Decode(buf, true)
	} else {
		data, err = l.Encode(buf, true)
	}
	runtime := time.Now().Sub(t1).Seconds()
	fmt.Println("Time:", runtime, " Speed:", float64(len(buf))/runtime/1024/1024, "MB/S")
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] writing data to", output)
	err = ioutil.WriteFile(output, data, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("[*] --- done! ---")
}
