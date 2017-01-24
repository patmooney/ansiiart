package main;

import (
    "./ansiiart"
    "fmt"
    "os"
    "log"
    "strconv"
);

func main () {
    var err error;
    var filename string;
    var resolution int = 40;

    switch nArgs := len(os.Args); nArgs {
        case 1:
            log.Fatal("Filename is required as first argument");
        case 2:
            filename = os.Args[1];
        default:
            filename = os.Args[1];
            resolution, err = strconv.Atoi(os.Args[2]);
            if err != nil {
                log.Fatal("Invalid value for resolution");
            }
    }

    ansii, err := ansiiart.TransformFile( filename, resolution );

    if err != nil {
        log.Fatal( err );
    }

    fmt.Println( ansii );
}
