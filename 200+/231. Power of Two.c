//
//
//date:2017-4-26
//desc: 10000 01111
//
//
bool isPowerOfTwo(int n) {
    if(n <= 0) return false;
    return (n & (n-1)) == 0;
}
