//
//date:2017-4-26
//desc:
//
//
bool isPowerOfFour(int num) {
    if(num <= 0) return false;
    if( (num & (num - 1)) == 0  && ((num-1)%3) == 0 ) return true;
    else return false;
}
