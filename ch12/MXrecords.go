package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := arguments[1]
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}

//➜  ch12 git:(master) ✗ go run MXrecords.go GOOGLE.COM
//aspmx.l.GOOGLE.COM.
//alt1.aspmx.l.GOOGLE.COM.
//alt2.aspmx.l.GOOGLE.COM.
//alt3.aspmx.l.GOOGLE.COM.
//alt4.aspmx.l.GOOGLE.COM.
//➜  ch12 git:(master) ✗ go run MXrecords.go golang.com
//aspmx.l.google.com.
//alt1.aspmx.l.google.com.
//alt3.aspmx.l.google.com.
//alt2.aspmx.l.google.com.
//➜  ch12 git:(master) ✗ go run MXrecords.go baidu.com
//mx.maillb.baidu.com.
//mx.n.shifen.com.
//usmx01.baidu.com.
//mx1.baidu.com.
//jpmx.baidu.com.
//mx50.baidu.com.
//➜  ch12 git:(master) ✗ host -t mx golang.com
//golang.com mail is handled by 1 aspmx.l.google.com.
//golang.com mail is handled by 2 alt2.aspmx.l.google.com.
//golang.com mail is handled by 2 alt3.aspmx.l.google.com.
//golang.com mail is handled by 2 alt1.aspmx.l.google.com.