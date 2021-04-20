package main

import (
	"bufio"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../sources/gamestate")
	conv_file, err_conv := os.Open("../sources/eu4bin.csv")
	if err_conv != nil {
		panic(err_conv)
	}
	defer conv_file.Close()

	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvr := csv.NewReader(conv_file)
	decoders := make(map[string]string)
	index_string := 0
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("broken row is", row)
			}
		}
		if index_string > 0 {
			splitted_string := strings.Split(row[0], ";")
			decoders[strings.ToLower(splitted_string[0][2:])] = splitted_string[1]
		}

		index_string++
	}

	stats, statsErr := file.Stat()
	fmt.Println("stats :", stats.Size())
	if statsErr != nil {
		fmt.Println("error")
	}

	//start decode
	var size int64 = stats.Size()

	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	dst := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(dst, bytes)

	var dst_dec []byte
	fmt.Println("starting decode")
	for key, val := range decoders {
		dst_dec = []byte(strings.ReplaceAll(string(dst), key, hex.EncodeToString([]byte(val))))
	}
	fmt.Println("finished decode")

	decoded := make([]byte, hex.DecodedLen(len(dst_dec)))
	n, err_dec := hex.Decode(decoded, dst_dec)
	if err_dec != nil {
		fmt.Println("error in decoding")
	}
	fmt.Println(n)
	f, err := os.Create("/output.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.Write(decoded)
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
}
