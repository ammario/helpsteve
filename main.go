package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

/*
Func main asks for a number, then asks sum, or pro.
Product multiplies all numbers 1-Selector and prints product(FIX).
Sum adds all numbers 1-selector(WORKING).
*/
func main() {
    //SSelector gets number
    SSelector := GetInput()
    //Iselector converts to string
    Iselector, err := strconv.Atoi(SSelector)
    if err != nil {
        panic(err)
    }
    fmt.Println("Sum, or product?")
    totalsum := 0
    for i := 0; i < 1; {
        //sumpro asks if you want to get the product or sum of the numbers 1-Iselector
        sumpro := GetInput()
        switch sumpro {
        case "sum":
            for j := 0; j <= Iselector; j++ {
                totalsum += j
            }
            fmt.Println(totalsum)
            i++
        case "pro":
            var totalpro int
            for k := 1; k <= Iselector; k++ {
                totalpro *= k
            }
            fmt.Println(totalpro)
        case "default":
            continue
        }

    }
    /*    for i := 0; i <= Iselector; i++ {
            fmt.Println(i)
        }
    */fmt.Println("Task Complete.")
}

func GetInput() string {
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    return strings.TrimSuffix(text, "\r\n")
}