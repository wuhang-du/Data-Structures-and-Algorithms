//
//date: 2017-5-7
//desc:	很经典的解法 
//
//
int hammingWeight(uint32_t n) {
    int count = 0;
    while(n){
        count++;
        n &= (n-1);
    }
    return count;
}
