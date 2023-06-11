package main

import (
    "fmt"

    "github.com/piao11piao/go_tutor_module/greetings"
    //这里是引入greetings这个module，我们之前创建的。但是由于greetings这个module不在标准库里，所以暂时是找不到的
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
