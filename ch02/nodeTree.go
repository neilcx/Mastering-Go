package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there!")
}

//➜  ch02 git:(master) ✗ go tool compile -W nodeTree.go
//'
//before walk os.(*File).close
//.   AS l(1) tc(1)
//.   .   NAME-main..this a(true) g(2) l(1) x(0) class(PPARAM) esc(no) tc(1) assigned used PTR-*os.File
//.   .   CONVNOP l(1) tc(1) PTR-*os.File
//.   .   .   DOTPTR l(1) x(0) tc(1) implicit(true) os.file PTR-*os.file
//.   .   .   .   NAME-main..this a(true) g(2) l(1) x(0) class(PPARAM) esc(no) tc(1) assigned used PTR-*os.File
//
//.   RETJMP l(1) tc(1) os.(*file).close
//after walk os.(*File).close
//.   AS l(1) tc(1) hascall
//.   .   NAME-main..this a(true) g(2) l(1) x(0) class(PPARAM) esc(no) tc(1) assigned used PTR-*os.File
//.   .   CONVNOP l(1) tc(1) hascall PTR-*os.File
//.   .   .   DOTPTR l(1) x(0) tc(1) implicit(true) hascall os.file PTR-*os.file
//.   .   .   .   NAME-main..this a(true) g(2) l(1) x(0) class(PPARAM) esc(no) tc(1) assigned used PTR-*os.File
//
//.   RETJMP l(1) tc(1) os.(*file).close
//enter os.(*File).close
//.   AS l(1)
//.   .   NAME-main.~r0 a(true) g(1) l(236) x(8) class(PPARAMOUT) esc(no) error
//
//before walk main
//.   DCL l(8)
//.   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//
//.   DCL l(8)
//.   .   NAME-fmt.n a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   DCL l(8)
//.   .   NAME-fmt.err a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   AS2 l(8) tc(1)
//.   AS2-list
//.   .   NAME-main.~arg0 a(true) l(8) x(0) class(PAUTO) esc(no) tc(1) assigned used INTER-interface {}
//.   AS2-rlist
//.   .   CONVIFACE l(8) esc(h) tc(1) implicit(true) INTER-interface {}
//.   .   .   NAME-main..stmp_0 a(true) l(8) x(0) class(PEXTERN) tc(1) used string
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//.   .   SLICELIT l(8) esc(no) tc(1) SLICE-[]interface {}
//.   .   .   LITERAL-1 l(8) untyped number
//.   .   SLICELIT-list
//.   .   .   NAME-main.~arg0 a(true) l(8) x(0) class(PAUTO) esc(no) tc(1) assigned used INTER-interface {}
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt.n a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt.err a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   INLMARK l(8) x(0)
//
//.   DCL l(8) tc(1)
//.   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   DCL l(8) tc(1)
//.   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   AS l(274) tc(1)
//.   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//
//.   AS2FUNC l(274) tc(1)
//.   AS2FUNC-list
//.   .   NAME-main..autotmp_6 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) used int
//
//.   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//.   AS2FUNC-rlist
//.   .   CALLFUNC l(274) tc(1) isddd(true) STRUCT-(int, error)
//.   .   .   NAME-fmt.Fprintln a(true) l(262) x(0) class(PFUNC) tc(1) used FUNC-func(io.Writer, ...interface {}) (int, error)
//.   .   CALLFUNC-list
//.   .   .   CONVIFACE l(274) esc(h) tc(1) io.Writer
//.   .   .   .   NAME-os.Stdout a(true) l(64) x(0) class(PEXTERN) tc(1) used PTR-*os.File
//
//.   .   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//
//.   AS2 l(274) tc(1)
//.   AS2-list
//.   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//.   AS2-rlist
//.   .   NAME-main..autotmp_6 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) used int
//
//.   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//
//.   VARKILL l(274) tc(1)
//.   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//
//.   VARKILL l(274) tc(1)
//.   .   NAME-main..autotmp_6 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) used int
//
//.   AS2 l(8) tc(1)
//.   AS2-list
//.   .   NAME-fmt.n a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   .   NAME-fmt.err a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//.   AS2-rlist
//.   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   GOTO l(8) tc(1) main..i0
//
//.   LABEL l(8) tc(1) main..i0
//after walk main
//.   DCL l(8)
//.   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//
//.   DCL l(8)
//.   .   NAME-fmt.n a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   DCL l(8)
//.   .   NAME-fmt.err a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   BLOCK-init
//.   .   AS l(8) tc(1)
//.   .   .   NAME-main..autotmp_8 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-interface {}
//.   .   .   EFACE l(8) tc(1) INTER-interface {}
//.   .   .   .   ADDR a(true) l(8) tc(1) PTR-*uint8
//.   .   .   .   .   NAME-type.string a(true) x(0) class(PEXTERN) tc(1) uint8
//.   .   .   .   ADDR l(8) tc(1) PTR-*string
//.   .   .   .   .   NAME-main..stmp_0 a(true) l(8) x(0) class(PEXTERN) tc(1) addrtaken used string
//.   BLOCK l(8) hascall
//.   BLOCK-list
//.   .   AS l(8) tc(1)
//.   .   .   NAME-main.~arg0 a(true) l(8) x(0) class(PAUTO) esc(no) tc(1) assigned used INTER-interface {}
//.   .   .   NAME-main..autotmp_8 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) assigned used INTER-interface {}
//
//.   EMPTY-init
//.   .   AS l(8) tc(1)
//.   .   .   NAME-main..autotmp_11 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-[1]interface {}
//
//.   .   AS l(8) tc(1)
//.   .   .   NAME-main..autotmp_9 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-*[1]interface {}
//.   .   .   ADDR l(8) tc(1) PTR-*[1]interface {}
//.   .   .   .   NAME-main..autotmp_11 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) addrtaken assigned used ARRAY-[1]interface {}
//
//.   .   BLOCK l(8)
//.   .   BLOCK-list
//.   .   .   AS l(8) tc(1) hascall
//.   .   .   .   INDEX l(8) tc(1) assigned bounded hascall INTER-interface {}
//.   .   .   .   .   DEREF l(8) tc(1) implicit(true) assigned hascall ARRAY-[1]interface {}
//.   .   .   .   .   .   NAME-main..autotmp_9 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-*[1]interface {}
//.   .   .   .   .   LITERAL-0 l(8) tc(1) int
//.   .   .   .   NAME-main.~arg0 a(true) l(8) x(0) class(PAUTO) esc(no) tc(1) assigned used INTER-interface {}
//
//.   .   BLOCK l(8)
//.   .   BLOCK-list
//.   .   .   AS l(8) tc(1) hascall
//.   .   .   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//.   .   .   .   SLICEARR l(8) tc(1) hascall SLICE-[]interface {}
//.   .   .   .   .   NAME-main..autotmp_9 a(true) l(8) x(0) class(PAUTO) esc(N) tc(1) assigned used PTR-*[1]interface {}
//.   EMPTY l(8) tc(1) hascall
//.   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt.n a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt.err a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   INLMARK l(8) x(0)
//
//.   DCL l(8) tc(1)
//.   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   DCL l(8) tc(1)
//.   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   AS l(8) tc(1)
//.   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   AS l(274) tc(1)
//.   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//
//.   BLOCK-init
//.   .   CALLFUNC l(274) tc(1) isddd(true) hascall STRUCT-(int, error)
//.   .   .   NAME-fmt.Fprintln a(true) l(262) x(0) class(PFUNC) tc(1) used FUNC-func(io.Writer, ...interface {}) (int, error)
//.   .   CALLFUNC-rlist
//.   .   .   EFACE l(274) tc(1) io.Writer
//.   .   .   .   ADDR a(true) l(274) tc(1) PTR-*uint8
//.   .   .   .   .   NAME-go.itab.*os.File,io.Writer a(true) l(274) x(0) class(PEXTERN) tc(1) uint8
//.   .   .   .   NAME-os.Stdout a(true) l(64) x(0) class(PEXTERN) tc(1) used PTR-*os.File
//
//.   .   .   NAME-fmt.a a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used SLICE-[]interface {}
//.   BLOCK l(274) hascall
//.   BLOCK-list
//.   .   AS l(274) tc(1)
//.   .   .   NAME-main..autotmp_6 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) used int
//.   .   .   RESULT l(274) x(40) tc(1) int
//
//.   .   AS l(274) tc(1)
//.   .   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//.   .   .   RESULT l(274) x(48) tc(1) error
//
//.   BLOCK l(274)
//.   BLOCK-list
//.   .   AS l(274) tc(1)
//.   .   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//.   .   .   NAME-main..autotmp_6 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) used int
//
//.   .   AS l(274) tc(1)
//.   .   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//.   .   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//
//.   VARKILL l(274) tc(1)
//.   .   NAME-main..autotmp_7 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) assigned used error
//
//.   VARKILL l(274) tc(1)
//.   .   NAME-main..autotmp_6 a(true) l(274) x(0) class(PAUTO) esc(N) tc(1) used int
//
//.   BLOCK l(8)
//.   BLOCK-list
//.   .   AS l(8) tc(1)
//.   .   .   NAME-fmt.n a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//.   .   .   NAME-fmt..autotmp_3 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used int
//
//.   .   AS l(8) tc(1)
//.   .   .   NAME-fmt.err a(true) l(273) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//.   .   .   NAME-fmt..autotmp_4 a(true) l(274) x(0) class(PAUTO) esc(no) tc(1) assigned used error
//
//.   GOTO l(8) tc(1) main..i0
//
//.   LABEL l(8) tc(1) main..i0
//➜  ch02 git:(master) ✗ go tool compile -W nodeTree.go | grep before
//before walk os.(*File).close
//before walk main
//➜  ch02 git:(master) ✗ go tool compile -W nodeTree.go | grep after
//after walk os.(*File).close
//after walk main