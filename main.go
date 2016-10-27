package main;

import (
    "github.com/patmooney/ansiiart/ansiiart"
    "fmt"
);

func main () {
    var out string = ansiiart.TransformFile( "/tmp/img.png", 10 );
    fmt.Println( out );
}
