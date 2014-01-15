// sorter.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sortting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err_ := br.ReadLine()
		if err_ != nil {
			if err_ != io.EOF {
				err = err_
			}
			break;
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
		}

		// 转换字符数组到字符串
		str := string(line)
		value, err_ := strconv.Atoi(str)
		if err_ != nil {
			err = err_
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main () {
	flag.Parse()
	if infile == nil {
		return
	}
	fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
}

