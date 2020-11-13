

//1. Accessing Structure Members


package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    
   var Book2 Books    
 
   
   Book1.title = "Go Programming"
   Book1.author = "Prasanthi"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   
   Book2.title = "java programming"
   Book2.author = " Geetha"
   Book2.subject = "java Programming Tutorial"
   Book2.book_id = 6495700
 
   
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

$go run main.go

Book 1 title : Go Programming
Book 1 author : Prasanthi
Book 1 subject : Go Programming Tutorial
Book 1 book_id : 6495407
Book 2 title : java programming
Book 2 author :  Geetha
Book 2 subject : java Programming Tutorial
Book 2 book_id : 6495700

// 2.structure (user)


package main

import "fmt"

type Address struct {
    city    string
    country string
}

type User struct {
    name    string
    age     int
    address Address
}

func main() {

    p := User{
        name: "prasanthi",
        age:  23,
        address: Address{
            city:    "New York",
            country: "USA",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("Country:", p.address.country)
}

$go run main.go

Name: prasanthi
Age: 23
City: New York
Country: USA