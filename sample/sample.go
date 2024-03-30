package main

import "fmt"

const globalConst = "foo"

func main() {
    var localVar = "foo"
    const localConst = "foo"

    fmt.Println("Literal String: foo") // Literal string in function call

    var arr = [...]string{"foo", "bar"} // Literal string in array initialization
    slice := []string{"foo", "bar"}     // Literal string in slice initialization
    m := map[string]int{"foo": 1, "bar": 2}

    // Literal string as map key
    m["foo"] = 1
    m["bar"] = 2

    // Literal string in struct field
    type Example struct {
        Name string
    }
    e := Example{Name: "foo"}

    // Literal string in composite literal
    s := struct {
        Name string
    }{Name: "foo"}

    // Literal string in type assertion
    var iface interface{} = "foo"
    str := iface.(string)

    // Literal string in switch case
    switch str {
    case "foo":
        fmt.Println("Switch case matched: foo")
    }

    // Literal string in channel send
    ch := make(chan string)
    go func() {
        ch <- "foo"
    }()

    // Literal string in channel receive
    select {
    case <-ch:
        fmt.Println("Received from channel: foo")
    }

    // Literal string in channel buffer
    chBuf := make(chan string, 1)
    chBuf <- "foo"

    // Literal string in map literal value
    mVal := map[int]string{1: "foo", 2: "bar"}

    // Literal string in function return
    returnedStr := returnString()

    // Literal string in function parameter
    funcParam("foo")

    // Literal string in method call
    var example Example
    example.method("foo")
}

func returnString() string {
    return "foo"
}

func funcParam(s string) {
    fmt.Println("Function parameter:", s)
}

func (e Example) method(s string) {
    fmt.Println("Method parameter:", s)
}
