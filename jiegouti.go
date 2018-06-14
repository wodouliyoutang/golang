package main

import "fmt"
// 结构体
// type后跟结构名称  struct 后跟声明的类型
type Lanague struct {
	num int 
	name string
	time int
}

func main(){
   /*
   var lenmu Lanague
   lenmu.num = 13
   lenmu.name = "channel"
   lenmu.time = 2018
   fmt.Println(lenmu.num)
   fmt.Println(lenmu.name)
   fmt.Println(lenmu.time)
   
   
   var lenmu Lanague
   lenmu.num = 13
   lenmu.name = "channel"
   lenmu.time = 2018
   //var_dump(lenmu)

   //结构体指针作为参数
   var_dump(&lenmu)
   
   s :=[3] int {1,2,3 }  //数组容量=长度 切不可改变
   fmt.Println(cap(s))
  
   numbers := make([]int,3)
   numbers[2] = 4
   fmt.Println(numbers)
    
   
   //空切片为  nill
   
   arrsli := []int {2,5,8,3,1,6}    //切片截图
   var_dump(arrsli)
   arr1 := arrsli[:2]
   var_dump(arr1)
   arr2 := arrsli[3:]
   var_dump(arr2)
   

   var numbers = []int {1,2,3,4}       //切片追加值,容量的变化规律
   fmt.Println("初始容量=",cap(numbers))
   var i int 
   for i=0; i<7; i++ {
      numbers  = append(numbers,2)
      fmt.Println("追加一",i+1,"次","容量=",cap(numbers))
   }

   var numbers []int          //切片的拷贝
   numbers = append(numbers, 2,3,4)

   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
 
   copy(numbers1,numbers)
   var_dump(numbers)   
   var_dump(numbers1)   


    nums := [][]int{{2, 3, 4},{6,9}}   //使用range遍历数组 key不用要用占位符替代
    for _, value := range nums {
      for keys, values := range value {
         fmt.Println("key=",keys)
         fmt.Println("value=",values)
      }
    }

    for i, c := range "go" {     //遍历unicode值
        fmt.Println(i, c)
    }
 
  
    // 创建集合       
   var countryCapitalMap = make(map[string]string)

   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"
   countryCapitalMap["China"] = "beijing" 
   value,_ := countryCapitalMap["China"]     //golang提倡写法
   fmt.Println(value)       //通过索引取值,若取没定义的,将返回所有的值,返回值是两个

   for key,country := range countryCapitalMap {
      fmt.Println("key:",key,"value:",country)
   }
   
   _, res := countryCapitalMap["s"]    // 查看元素在集合中是否存在  第二个参数为bool
   if res {
      fmt.Println("存在")  
   }else {
      fmt.Println("不存在") 
   }
  
   
   result := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
   
   fmt.Println("-------")   
   for key,value := range result {
      fmt.Println(key,":",value)
   }

   // 删除元素 
   delete(result,"France");
   fmt.Println("-------")   
   for key,value := range result {
      fmt.Println(key,":",value)
   }  
   
  var sum int = 10   //数据类型转换
  var son int = 3
  mean := float64(sum) /float64(son)
  fmt.Println(mean)
   */

}

/*
func var_dump(data Lanague) {    //结构体作为参数
   fmt.Println(data.num)
   fmt.Println(data.name)
   fmt.Println(data.time)
}

func var_dump(data *Lanague) {    //结构体作为参数
   fmt.Println(&data.num)
   fmt.Println(data.name)
   fmt.Println(data.time)
}

func var_dump(x []int){
   fmt.Println(len(x),cap(x),"value=",x)
}

*/