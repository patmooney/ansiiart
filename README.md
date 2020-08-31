我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# ansiiart

Package ansiiart is here to convert images of all popular types into the far
superious format of bash ansii art. Just think of the possibilities!? have
your favourite kitten picture print out when you start a shell or even
use this package to make some sort of awesome shell based magazine or something.
Endless possibilities.

*Image goes in - Ansii art comes out ( linux shell art! )*

![Example image](doc/example.png)

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

