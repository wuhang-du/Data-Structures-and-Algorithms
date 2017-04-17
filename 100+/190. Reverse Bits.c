void setbit(uint32_t * pointer,int pos)
{
    uint32_t temp = 1;
    temp = temp << pos ;
    *pointer = *pointer | temp ;
}


int isbit(uint32_t * pointer,int pos)
{
    uint32_t temp = 1;
    temp = temp << (pos - 1);
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
