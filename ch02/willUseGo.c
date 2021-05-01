#include <stdio.h>
#include "usedByC.h"

int main(int argc, char **argv) {
	GoInt x = 12;
	GoInt y = 23;

	printf("About to call a Go function!\n");
	PrintMessage();
	GoInt p = Multiply(x,y);
	printf("Product: %d\n",(int)p);
	printf("It worked!\n");
	return 0;
}

//➜  ch02 git:(master) ✗ gcc -o willUseGo willUseGo.c ./usedByC.o
//➜  ch02 git:(master) ✗ ./willUseGo
//About to call a Go function!
//A Go function!
//Product: 276
//It worked!