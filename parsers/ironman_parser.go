package main

import (
	"bufio"
	"bytes"
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
			if len(splitted_string[1]) > 0 && splitted_string[1] != "String" && splitted_string[1] != "Integer" && splitted_string[1] != "Float" && splitted_string[1] != "Boolean" {
				decoders[strings.ToLower(splitted_string[0][2:])] = splitted_string[1]
			} else if splitted_string[1] == "String" || splitted_string[1] == "Integer" || splitted_string[1] == "Float" || splitted_string[1] == "Boolean" {
				decoders[strings.ToLower(splitted_string[0][2:])] = " "
			} else if splitted_string[0] == "0x0600" {
				decoders[strings.ToLower(splitted_string[0][2:])] = string("\n")
			}
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

	file_bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(file_bytes)

	dst := make([]byte, hex.EncodedLen(len(file_bytes)))
	hex.Encode(dst, file_bytes)
	byteReader := bytes.NewReader(dst)
	hex_reader := make([]byte, 2)
	dst_decoded := make([]byte, hex.EncodedLen(len(file_bytes)))
	hex_buf := bytes.Buffer{}

	fmt.Println("starting decode", len(dst))
	tmp := ""
	for {
		_, err := byteReader.Read(hex_reader)
		if err == io.EOF {
			dst_decoded = hex_buf.Bytes()
			break
		}

		tmp += string(hex_reader)
		if len(tmp) <= 2 {
			_, err = byteReader.Read(hex_reader)
			if err == io.EOF {
				dst_decoded = hex_buf.Bytes()
				break
			}
			tmp += string(hex_reader)
		}

		if val, ok := decoders[tmp]; ok {
			hex_buf.Write([]byte(hex.EncodeToString([]byte(val))))
			tmp = ""
		} else {
			hex_buf.Write([]byte(tmp[:2]))
			tmp = tmp[2:]
		}
	}

	fmt.Println("finished decode", len(dst))

	decoded := make([]byte, hex.DecodedLen(len(dst_decoded)))
	n, err_dec := hex.Decode(decoded, dst_decoded)
	if err_dec != nil {
		fmt.Println("error in decoding")
	}
	fmt.Println(n)
	f, err := os.Create("./output.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.Write(decoded)
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
}
