package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main(){
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.w3cschool.cn"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.w3cschool.cn"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   fmt.Println(Book1.title)
   fmt.Println(Book1.author)
   fmt.Println(Book1.subject)
   fmt.Println(Book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Println(Book2.title)
   fmt.Println(Book2.author)
   fmt.Println(Book2.subject)
   fmt.Println(Book2.book_id)

}
