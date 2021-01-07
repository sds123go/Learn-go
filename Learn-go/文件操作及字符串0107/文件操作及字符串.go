package main

import (
	//"bufio"
	"fmt"
	//"strings"
	"strconv"
	// "io"
	// "os"
)

/*
//CreateFile 创建一个文件,并使用WriteString()写入数据
func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	for i := 0; i < 5; i++ {
		f.WriteString("Hello World\n")
	}

	defer f.Close()
}
*/
/*
//CreateFile 创建一个文件,并用Write()写入数据
func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	str := "hello world"
	buf := []byte(str)
	for i := 0; i < 5; i++ {
		n, err := f.Write(buf)
		if err != nil {
			fmt.Println("err = %s", err)

		}
		fmt.Println(n)
	}

	defer f.Close()
}
*/
/*
//CreateFile 创建一个文件,并用WriteAt()写入数据
func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	str := "hello world"
	buf := []byte(str)
	for i := 0; i < 5; i++ {
		n, _ := f.Seek(0, os.SEEK_END) //查找文件末尾的偏移量
		a, err := f.WriteAt(buf, n)    // 从偏移量开始写入文件
		if err != nil {
			fmt.Println("err = %s", err)
		}
		fmt.Println(a)
	}

	defer f.Close()
}*/
/*
func WriteFile(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND, 6)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	n, err := f.WriteString("这是追加的内容。。。\n")
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	fmt.Println(n)
	defer f.Close()

}*/
/*
//ReadFile 读取文件数据,全部读取
func ReadFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	buf := make([]byte, 1024*2) //2K大小内存
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错同时未到文件结尾
		fmt.Println("err = %s", err1)
		return
	}
	fmt.Println(string(buf[:n]))
	defer f.Close()

}
*/
/*
// ReadFile 按行读取数据
func ReadFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	//新建一个缓冲区
	r := bufio.NewReader(f)
	for {
		//遇到换行符读取，换行符也读取在内
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Println(string(buf))
	}
	defer f.Close()

}
*/
/*
//CopyFile 复制文件
func CopyFile(srcfilename string, dstfilename string) {
	if srcfilename == dstfilename {
		fmt.Println("不能同名！！！")
		return
	}
	sf, err := os.Open(srcfilename)
	if err != nil {
		fmt.Println("err = %s", err)
		return
	}
	buf := make([]byte, 1024*4) //4K大小内存
	n, err1 := sf.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错同时未到文件结尾
		fmt.Println("err = %s", err1)
		return
	}
	df, err2 := os.Create(dstfilename)
	if err2 != nil {
		fmt.Println("err = %s", err2)
		return
	}
	df.Write(buf[:n])
	defer sf.Close()
	defer df.Close()
}
func main() {
	var srcfilename string
	var dstfilename string
	fmt.Println("请输入要复制的文件名:")
	fmt.Scan(&srcfilename)
	fmt.Println("请输入目的文件名:")
	fmt.Scan(&dstfilename)
	CopyFile(srcfilename, dstfilename)

}*/
func main() {
	// str:="2021-1-7"
	// s:=strings.Split(str,"-")
	// fmt.Printf("%s年%s月%s日",s[0],s[1],s[2])
	// var str string
	// //浮点型转换成字符串，'f'表示打印格式以小数表示，3是小数点位数，64表示以float64表示
	// str=strconv.FormatFloat(3.14, 'f', 3, 64)
	// fmt.Println(str)  //3.140

	// var flag bool
	// var err error
	// //字符串转换成bool型
	// flag, err = strconv.ParseBool("true")
	// if err == nil {
	// 	fmt.Print(flag)
	// } else {
	// 	fmt.Print(err)
	// }

	// //字符串转整型
	// a, _ := strconv.Atoi("123")
	// fmt.Print(a)
	//字符串转换成浮点型
	f, err := strconv.ParseFloat("3.14", 64)
	if err == nil {
		fmt.Print(f)
	} else {
		fmt.Print(err)
	}
}
