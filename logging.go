package main

import (
   "pkg"
   "fmt"
)

const (
   VERSION = "0.13"
 )

func main() {
    bar := "test log"
    message := fmt.Sprintf("foo: %s", bar)
    logger.Debug(message)
}

