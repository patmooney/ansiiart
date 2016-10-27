# ansiiart

Package ansiiart is here to convert images of all popular types into the far
superious format of bash ansii art. Just think of the possibilities!? have
your favourite kitten picture print out when you start a shell or even
use this package to make some sort of awesome shell based magazine or something.
Endless possibilities.

*Image goes in - Ansii art comes out ( linux shell art! )*

### Usage: as an app

    # resolution is the amount of "pixels" rendered
    go run main.go <image_file> <resolution>

    # e.g.

    go run main.go my-picture.png 50 > my-picture.ansii
    cat my-picture.ansii

### Usage: as a package

    package main;

    import (
        "fmt"
        "github.com/skinnyjeans/ansiiart"
    );

    func main () {
        ansii, err := ansiiart.TransformFile( "kitten.png" );
        fmt.Println( ansii );
    }

