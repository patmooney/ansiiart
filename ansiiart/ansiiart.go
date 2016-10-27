// Package ansiiart is here to convert images of all popular types into the far
// superious format of bash ansii art. Just think of the possibilities, have
// your favourite kitten picture print out when you start a shell or even
// use this package to make some sort of awesome shell based browser with images!
package ansiiart;

import (
    "image"
    "os"
    "io"
    "fmt"
    "strings"
    "strconv"
    "math"

    _ "image/png"
    _ "image/gif"
    _ "image/jpeg"
);

var _source map[int][]int;

// TransformFile accepts a filename (str) and resolution (int, 50 recommended)
// The resolution dictates the accuracy of the render but also the size of the
// outputted image.
// Example
//      ansiiart.TransformFile( "myImage.jpg", 49 );
// The output is a string of unix bash text colour codes
func TransformFile ( filename string, resolution int ) ( string, error ) {
    reader, err := os.Open(filename);
    defer reader.Close();
    if err != nil {
        return "", err;
    }

    return Transform( reader, resolution );
}

// Transform is similar to TransformFile but accepts a io.Reader
func Transform ( reader io.Reader, resolution int ) ( string, error ) {
    img, _, err := image.Decode(reader)
    if err != nil {
        return "", err;
    }

    bounds := img.Bounds();

    var spacing float64 = float64(bounds.Max.X) / float64(resolution);
    var output string = "";
    var previousAnsii string = "";
    // take samples from the source image and then map
    // to the nearest RGB ansii colour
    for y := 0; y < bounds.Max.Y; y = y + int(spacing) {
        for x := 0; x < bounds.Max.X; x = x + int(spacing/2) {
            r, g, b, _ := img.At( x, y ).RGBA();
            var ansii string = toAnsii([]int{ int( r>>8 ), int( g>>8 ), int( b>>8 ) });
            if ( ansii == previousAnsii ){
                output = output + " ";
            } else {
                output = output + fmt.Sprintf( "\033[0m%s ", ansii );
                previousAnsii = ansii;
            }
        }
        output = output + "\033[0m\n";
        previousAnsii = "";
    }

    return output, nil;
}

func toAnsii ( colour []int ) string {
    var source map[int][]int = loadSource();

    var closest int;
    var chosen int;

    for n := 0; n < len(source); n++ {
        var check int = getDist( colour, source[n] );
        if closest == 0 {
            closest = check;
            chosen = n;
        } else if check < closest {
            closest = check;
            chosen = n;
        }
    }

    return fmt.Sprintf( "\033[48;5;%dm\033[38;5;%dm", chosen, chosen );
}

func loadSource () map[int][]int {

    if len(_source) > 0 {
        return _source;
    }

    var codes []string = strings.Split( _ansiiSource, ":" );
    source := make(map[int][]int, len(codes));

    for _, code := range codes {
        var parts []string = strings.Split( code, "," );

        ansii, _ := strconv.Atoi(parts[0]);
        r, _ := strconv.Atoi(parts[1]);
        g, _ := strconv.Atoi(parts[2]);
        b, _ := strconv.Atoi(parts[3]);

        source[ansii] = []int{r, g, b};
    }

    _source = source;

    return _source;
}

// #pythag
func getDist( a []int, b []int ) int {
    dist := math.Pow(
        math.Pow( float64(a[0] - b[0]), 2 ) +
        math.Pow( float64(a[1] - b[1]), 2 ) +
        math.Pow( float64(a[2] - b[2]), 2 ),
        0.5,
    );
    return int(dist);
}

var _ansiiSource string = "0,0,0,0:1,128,0,0:2,0,128,0:3,128,128,0:4,0,0,128:5,128,0,128:6,0,128,128:7,192,192,192:8,128,128,128:9,255,0,0:10,0,255,0:11,255,255,0:12,0,0,255:13,255,0,255:14,0,255,255:15,255,255,255:16,0,0,0:17,0,0,95:18,0,0,135:19,0,0,175:20,0,0,215:21,0,0,255:22,0,95,0:23,0,95,95:24,0,95,135:25,0,95,175:26,0,95,215:27,0,95,255:28,0,135,0:29,0,135,95:30,0,135,135:31,0,135,175:32,0,135,215:33,0,135,255:34,0,175,0:35,0,175,95:36,0,175,135:37,0,175,175:38,0,175,215:39,0,175,255:40,0,215,0:41,0,215,95:42,0,215,135:43,0,215,175:44,0,215,215:45,0,215,255:46,0,255,0:47,0,255,95:48,0,255,135:49,0,255,175:50,0,255,215:51,0,255,255:52,95,0,0:53,95,0,95:54,95,0,135:55,95,0,175:56,95,0,215:57,95,0,255:58,95,95,0:59,95,95,95:60,95,95,135:61,95,95,175:62,95,95,215:63,95,95,255:64,95,135,0:65,95,135,95:66,95,135,135:67,95,135,175:68,95,135,215:69,95,135,255:70,95,175,0:71,95,175,95:72,95,175,135:73,95,175,175:74,95,175,215:75,95,175,255:76,95,215,0:77,95,215,95:78,95,215,135:79,95,215,175:80,95,215,215:81,95,215,255:82,95,255,0:83,95,255,95:84,95,255,135:85,95,255,175:86,95,255,215:87,95,255,255:88,135,0,0:89,135,0,95:90,135,0,135:91,135,0,175:92,135,0,215:93,135,0,255:94,135,95,0:95,135,95,95:96,135,95,135:97,135,95,175:98,135,95,215:99,135,95,255:100,135,135,0:101,135,135,95:102,135,135,135:103,135,135,175:104,135,135,215:105,135,135,255:106,135,175,0:107,135,175,95:108,135,175,135:109,135,175,175:110,135,175,215:111,135,175,255:112,135,215,0:113,135,215,95:114,135,215,135:115,135,215,175:116,135,215,215:117,135,215,255:118,135,255,0:119,135,255,95:120,135,255,135:121,135,255,175:122,135,255,215:123,135,255,255:124,175,0,0:125,175,0,95:126,175,0,135:127,175,0,175:128,175,0,215:129,175,0,255:130,175,95,0:131,175,95,95:132,175,95,135:133,175,95,175:134,175,95,215:135,175,95,255:136,175,135,0:137,175,135,95:138,175,135,135:139,175,135,175:140,175,135,215:141,175,135,255:142,175,175,0:143,175,175,95:144,175,175,135:145,175,175,175:146,175,175,215:147,175,175,255:148,175,215,0:149,175,215,95:150,175,215,135:151,175,215,175:152,175,215,215:153,175,215,255:154,175,255,0:155,175,255,95:156,175,255,135:157,175,255,175:158,175,255,215:159,175,255,255:160,215,0,0:161,215,0,95:162,215,0,135:163,215,0,175:164,215,0,215:165,215,0,255:166,215,95,0:167,215,95,95:168,215,95,135:169,215,95,175:170,215,95,215:171,215,95,255:172,215,135,0:173,215,135,95:174,215,135,135:175,215,135,175:176,215,135,215:177,215,135,255:178,215,175,0:179,215,175,95:180,215,175,135:181,215,175,175:182,215,175,215:183,215,175,255:184,215,215,0:185,215,215,95:186,215,215,135:187,215,215,175:188,215,215,215:189,215,215,255:190,215,255,0:191,215,255,95:192,215,255,135:193,215,255,175:194,215,255,215:195,215,255,255:196,255,0,0:197,255,0,95:198,255,0,135:199,255,0,175:200,255,0,215:201,255,0,255:202,255,95,0:203,255,95,95:204,255,95,135:205,255,95,175:206,255,95,215:207,255,95,255:208,255,135,0:209,255,135,95:210,255,135,135:211,255,135,175:212,255,135,215:213,255,135,255:214,255,175,0:215,255,175,95:216,255,175,135:217,255,175,175:218,255,175,215:219,255,175,255:220,255,215,0:221,255,215,95:222,255,215,135:223,255,215,175:224,255,215,215:225,255,215,255:226,255,255,0:227,255,255,95:228,255,255,135:229,255,255,175:230,255,255,215:231,255,255,255:232,8,8,8:233,18,18,18:234,28,28,28:235,38,38,38:236,48,48,48:237,58,58,58:238,68,68,68:239,78,78,78:240,88,88,88:241,98,98,98:242,108,108,108:243,118,118,118:244,128,128,128:245,138,138,138:246,148,148,148:247,158,158,158:248,168,168,168:249,178,178,178:250,188,188,188:251,198,198,198:252,208,208,208:253,218,218,218:254,228,228,228:255,238,238,238";

