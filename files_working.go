// 1. create file 

package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Create("myfile.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

$go run main.go

11 bytes written successfully

Program exited.


// 2. file operations

package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    mydata := []byte("I write all the data to a file\n")

    // the WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.data", mydata, 0777)
    // handle this error
    if err != nil {
        // print it out
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

    f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
        panic(err)
    }

    data, err = ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}

$go run main.go

I write all the data to a file
I write all the data to a file
new data that wasn't there originally


