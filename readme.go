byte  别名init8类型
代表utf8字符单个字节的值
rune  
代表unicode字符值

两个相当于 "sdfdfs"   和   "汉字sd省道"

字符切割用


不定参数,向某个方法传递的参数个数不固定
func myfunc(args ...int) {
	for _, arg := range args {
	fmt.Println(arg)
	}
}