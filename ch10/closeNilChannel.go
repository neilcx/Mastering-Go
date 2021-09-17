package main

func main() {
	var c chan string
	close(c)
}


//go run closeNilChannel.go
//panic: close of nil channel
//
//goroutine 1 [running]:
//main.main()
///Users/***/workspace/github/Mastering-Go/ch10/closeNilChannel.go:5 +0x2a
//exit status 2
//cn0214006386m:ch10