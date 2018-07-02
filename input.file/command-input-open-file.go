package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"

var filea *string = flag.String("a", "file-a-is-empty", "File contains values for sorting")
var fileb *string = flag.String("b", "file-b-is-empty", "File to receive sorted values")
var filec *string = flag.String("c", "file-c-is-empty", "Sort algorithm")

/**
 * 通过命令行参数传入文件,并读出文件内容
 * @return {[type]} [description]
 */
func main() {
	flag.Parse()

	values, err := readValues(*filea)
	if err == nil {
		fmt.Println("文件内容:", values)
		err := writeValues(values, "d.ini")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(err)
	}
}

/**
 * 通过文件指针读数据
 * @param  {[type]} infile string)       (values []string, err error [description]
 * @return {[type]}        [description]
 */
func readValues(infile string) (values []string, err error) {
	//打开文件
	file, err := os.Open(infile)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	//创建文件缓存区
	br := bufio.NewReader(file)
	values = make([]string, 0)
	// br.ReadLine()返回值 [49] false <nil>
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line) // 转换字符数组为字符串
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, str)
	}
	return
}

/**
 * 新建文件并存入内容
 * @param values 文件内容
 * @param outfile 文件名称
 */
func writeValues(values []string, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {

		file.WriteString(value + "\n")
	}
	return nil
}
