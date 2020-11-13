
// 1. java script object notation(JSON)

package main  
import "encoding/json"  
import "fmt"  
  
func main() {  
    bolType, _ := json.Marshal(false) //boolean Value  
    fmt.Println(string(bolType))  
    intType, _ := json.Marshal(10) // integer value  
    fmt.Println(string(intType))  
    fltType, _ := json.Marshal(3.14) //float value  
    fmt.Println(string(fltType))  
    strType, _ := json.Marshal("go") // string value  
    fmt.Println(string(strType))  
    slcA := []string{"sunday", "monday", "tuesday"} //slice value  
    slcB, _ := json.Marshal(slcA)  
    fmt.Println(string(slcB))  
    mapA := map[string]int{"sunday": 1, "monday": 2} //map value  
    mapB, _ := json.Marshal(mapA)  
    fmt.Println(string(mapB))  
}  

$go run main.go

false
10
3.14
"go"
["sunday","monday","tuesday"]
{"monday":2,"sunday":1}

