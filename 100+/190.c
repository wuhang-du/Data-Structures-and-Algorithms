#include <stdio.h>

typedef unsigned int uint32_t;

enum {false = 0,true};

void setbit(uint32_t * pointer,int pos)
{
    uint32_t temp = 1;
    temp = (pos)? temp << (pos) : 1 ;
    *pointer = *pointer | temp ;
}


int isbit(uint32_t * pointer,int pos)
{
    uint32_t temp = 1;
    temp = (pos - 1)? temp << (pos - 1) : 1 ;
    if(*pointer & temp) return true;
    else return false;
}


uint32_t reverseBits(uint32_t n) {
    uint32_t answer = 0;
    for(int i = 32; i > 0; i--){
        if(isbit(&n,i)) setbit(&answer,32-i);
    }
    return answer;
}

void main(void){
	unsigned int a = 000000010100101000001111010011100;
	unsigned int answer = reverseBits(a);
	printf("%u",answer);
}
