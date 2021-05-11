package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once!")
	}

//	go run funWithFlag.go -k=100 -o=chenxi -names=haha,lala,kaka,kaka diyi dier disan disan
//v: haha,lala,kaka,kaka
//	item haha
//	item lala
//	item kaka
//	item kaka

	fmt.Println("v:", v)
	names := strings.Split(v, ",")
	for _, item := range names {
		fmt.Println("item",item)
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	//0 %c 91
	//1 %c 104
	//2 %c 97
	//3 %c 104
	//4 %c 97

	//5 %c 32

	//6 %c 108
	//7 %c 97
	//8 %c 108
	//9 %c 97

	//10 %c 32

	//11 %c 107
	//12 %c 97
	//13 %c 107
	//14 %c 97

	//15 %c 32

	//16 %c 107
	//17 %c 97
	//18 %c 107
	//19 %c 97

	//20 %c 93

	for _, item := range manyNames.String() {
		fmt.Printf("%c", item) //[haha lala kaka kaka]
	}

	fmt.Println("\nRemaining command line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}

//go run funWithFlag.go -k=100 -o=neil -names=haha -names=heihei
//invalid value "heihei" for flag -names: Cannot use names flag more than once!
//Usage of /var/folders/n1/c53wz14x6rdfgyv467bmhpt8m5hmgl/T/go-build052387376/b001/exe/funWithFlag:
//-k int
//An int
//-names value
//Comma-separated list
//-o string
//The name (default "Mihalis")
//exit status 2