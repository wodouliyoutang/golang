package main

 import "fmt"	
import _"reflect"
import _"math"
// import "unsafe"
 //引入的包必须要使用 ,不使用用 _声明

// var name = "string-type"
// var email string = "3983@qq.com"
/*
func init() {  //包的构造方法
	fmt.Println("init")
}
*/
// type myint int //自定义数据类型
func main() {
	// phone := 123	//这种不带声明格式的只能在函数体中出现
	//phone := 123 		//报错
   //var adr int		
   //println(&adr)		//取地址
    //println(name,email,phone)
	// println(phone)
	//var c, d int = 1, 2
	//println(c,d)	
	 //const c = "s"	//字符串类型，单引号asis码 双引号string
	 // const NUMS int = 123
	 // const SEX int =12
	 // const A, b = 'cona', 'conb'
	// const (
	// 	a="abc" 
	// 	b=a
	// 	)
 //    println(a,b) 
 // 	a := unsafe.Sizeof("abc")
	// fmt.Println(a)
	// const (
	// 	a = iota
	// 	b = iota
	// 	)
	// if a <1 {
	// 	if b != 1 {

	// 		fmt.Println(a,b)
	// 	}
	// 		fmt.Println('e')
	// }
	/*
	grode := "B"
	score := 3
	scores := 8
	
	switch score {
		case 80 : grode = "B"
		case 90 : grode = "A"
		case 50,60,70 : grode = "C"
		default: grode = "D"
	}
	fmt.Println(grode)
	
	switch  {
		case score == 80 : 
			grode = "B"
		case scores == 8 : 
			grode = "A"
		case score == 70 :
			grode = "C"
	}
	fmt.Println(grode)
	*/
   // var sum1 int = 14
   // var sum2 int = 13
   // num3 := max(sum1,sum2)
   // var num3 int
   // var num4 int	
   // num3,num4 = test(sum1,sum2)	//定义的变量一定要使用到,编译时检测到定义的变量没使用会报错
   // fmt.Println(num3, num4)
   //num1 := 12
   //fmt.Println(reflect.TypeOf(num1))	//判断数据类型    xx.(type)只能存在于switch中
   // num1 := 12
   // num2 := 14
   // fmt.Println(num1,num2)
   // fmt.Println(&num1,&num2)

   // swap(&num1, &num2)
   // fmt.Println(num1,num2)
   // fmt.Println(&num1,&num2)
   /*
   getNumSqr := func(x float64) float64{ //函数体作为值
   		return math.Sqrt(x)
   }
   fmt.Println(getNumSqr(15))
   */
  /*
  var myint myint = 3	//函数与方法
  myint.mycla()
  fmt.Println(myint)
  */
  // var array = [5] int {1,3,5,7,9}
  // fmt.Println(array[2])
    // var a int = 21
   // var b int = 10
   // var c int

   // c = a + b
   // fmt.Printf("第一行 - c 的值为 [%d]\n", c )	//占位符 %d 10进制数字  %v 相应值的默认格式
   // fmt.Printf("test[%v]","sdf")
 
  // var n [10]int //数组中数据必须是同一数据类型
  // var i,j int

  //  for i = 0; i < 10; i++ {
  //     n[i] = i + 100     
  //  }         

  //  for j = 0; j < 10; j++ { 
  //     fmt.Printf("Element[%d] = %d\n", j, n[j] ) 
  //  } 
  // var a [3]int  
  // a[2] = 201  
  // fmt.Println(a[2]) 
  // var ns [10]int
  // ns[0] = 100
  // fmt.Printf("haha%d",ns[0])
  // Printf 只能输出字符串,需要占位符结合才能输出整型，Println则不用
  /*
   var a = [5][2]int{ //多维数组声明换行报错
   	  {0,0},
   	  {1,2},
   	  {2,4}, 
   	  {3,6},
   	  {4,8}
   	}

   fmt.Println(a[1][1])
   
  	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
  	var i,j int 
  	for i=0; i<len(a); i++ {
  		for j=0; j<2; j++ {
  			fmt.Printf("a[%d] = %d\n", i, a[i][j])
  		}
  	}
   
   var  balance = [5]int {1000, 2, 3, 17, 50}  // []int 切片数据类型, [5]数组数据类型  作为函数形参时要一致
   var avg float32

   avg = AvgNnm( balance ) ;

   fmt.Printf( "%f ", avg );
	
   	
  // 定义指针变量。
  // 为指针变量赋值。 赋值要赋予地址
  // 访问指针变量中指向地址的值    访问指针的值要 加 *
   var val = 20	
   var ip *int
   ip = &val
   fmt.Println(*ip)
   
   //空指针为 nil
   var ip *int
   fmt.Println(ip)
	

   //指针数组   可用于存储数组数据
   const MAX = 3
   a := []int{10,100,200}
   var i int
   var ptr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptr[i] = &a[i] 	// 整数地址赋值给指针数组 
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )   //通过指针访问值
      fmt.Printf("a[%d] = %d\n", i,a[i] )     //直接访问值
   }

   //指向指针的指针变量      
   var a int
   var ptr *int 
   var pptr **int   //指向另一个指针的地址

   a = 3000

   // 指针 ptr 地址
   ptr = &a

   // 指向指针 ptr 地址 
   pptr = &ptr

   // 获取 pptr 的值 
   fmt.Printf("变量 a = %d\n", a )
   fmt.Printf("指针变量 *ptr = %d\n", *ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
   

   var a = 100		
   var b = 200
   adrswp(&a,&b)	//向函数传递指针达到交换值目的
   fmt.Println(a,b)
   */


}
/*
func max (num1 int, num2 int) int {
	var num3 int
	if num1 > num2 {
		num3 = num1
	} else {
		num3 = num2
	}

	return num3
}

func test(num1 int, num2 int) (int, int){
	return num1,num2
}
*/
/*
func swap(x *int, y *int) { //地址交换实参需加  *号
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func (p *myint) mycla () int {  //函数与方法 此为方法
	*p = *p * 2
	return 1
}

func AvgNnm (arr [5]int) float32 {  //求平均数
	var i,sum int

	for i=0; i<len(arr); i++ {
		sum += arr[i]
	}
	avg := float32(sum / len(arr))
	return  avg
}


func adrswp (x *int,y *int) {	//交换指针值
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp

} 
*/
