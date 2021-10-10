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
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}

//
//➜  ch12 git:(master) ✗ go run NSrecords.go google.com
//ns2.google.com.
//ns1.google.com.
//ns3.google.com.
//ns4.google.com.
//➜  ch12 git:(master) ✗ go run NSrecords.go www.google.com
//lookup www.google.com on 10.4.192.27:53: no such host
//➜  ch12 git:(master) ✗ go run NSrecords.go baidu.com
//ns2.baidu.com.
//ns3.baidu.com.
//ns4.baidu.com.
//ns1.baidu.com.
//ns7.baidu.com.
//➜  ch12 git:(master) ✗ go run NSrecords.go baidu.com
//ns3.baidu.com.
//ns4.baidu.com.
//ns1.baidu.com.
//ns7.baidu.com.
//ns2.baidu.com.
//➜  ch12 git:(master) ✗ host -t ns google.com
//google.com name server ns2.google.com.
//google.com name server ns1.google.com.
//google.com name server ns3.google.com.
//google.com name server ns4.google.com.
//➜  ch12 git:(master) ✗ host -t ns www.google.com
//www.google.com has no NS record