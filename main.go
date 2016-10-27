package main;

import (
    "github.com/patmooney/ansiiart/ansiiart"
    "fmt"
    "os"
);

func main () {
    var out string = ansiiart.TransformFile( os.Args[1], 50 );
    fmt.Println( out );
}
