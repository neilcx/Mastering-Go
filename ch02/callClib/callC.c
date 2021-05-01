#include <stdio.h>
#include "callC.h"

void cHello() {
    printf("Hello from C!\n");
}


// message is a pointer of char
void printMessage(char* message) {
	printf("Go send me %s\n", message);
}