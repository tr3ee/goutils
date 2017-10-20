package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"tr3e/utils/cli-banner"
	"tr3e/utils/trans/layers"
)

var input, output string
var dFlag bool

func init() {
	flag.StringVar(&input, "i", "", "input data file path")
	flag.StringVar(&output, "o", "", "path to output the file")
	flag.BoolVar(&dFlag, "d", false, "decompress the input file if set to true")
}

func main() {
	flag.Parse()
	banner.Autoload("LZ4 compress", "example", "using this program to compress/decompress data")
	if input == "" || output == "" {
		flag.Usage()
		return
	}
	fmt.Println("[*] reading data from", input)
	buf, err := ioutil.ReadFile(input)
	fmt.Println("[*] reading data --done-- ")

	fmt.Println("[*] processing data with lz4 algorithm")
	var data []byte
	if dFlag {
		data, err = layers.LZ4Decode(buf)
	} else {
		data, err = layers.LZ4Encode(buf)
	}
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
