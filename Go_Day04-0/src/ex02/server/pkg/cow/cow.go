package cow

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
unsigned int i;
unsigned int argscharcount = 0;
char* ask_cow(char phrase[]) {
	int phrase_len = strlen(phrase);
	char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
	strcpy(buf, " ");
	for (i = 0; i < phrase_len + 2; ++i) {
		strcat(buf, "_");
	}
	strcat(buf, "\n< ");
	strcat(buf, phrase);
	strcat(buf, " ");
	strcat(buf, ">\n ");
	for (i = 0; i < phrase_len + 2; ++i) {
		strcat(buf, "-");
	}
	strcat(buf, "\n");
	strcat(buf, "        \\   ^__^\n");
	strcat(buf, "         \\  (oo)\\_______\n");
	strcat(buf, "            (__)\\       )\\/\\\n");
	strcat(buf, "                ||----w |\n");
	strcat(buf, "                ||     ||\n");
	return buf;
}
*/
import "C"
import (
	"bytes"
	"unsafe"
)

func AskCow() string {
	phrase := C.CString("Thank you!")

	defer C.free(unsafe.Pointer(phrase))
	phraseLen := C.int(C.strlen(phrase))

	ptr := C.ask_cow(phrase)
	defer C.free(unsafe.Pointer(ptr))

	phraseLen += C.int(C.strlen(ptr))
	p := C.GoBytes(unsafe.Pointer(ptr), phraseLen)
	p = bytes.TrimRight(p, "\x00")

	return string(p)
}
