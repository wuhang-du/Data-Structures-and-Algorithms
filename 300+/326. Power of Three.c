//
//date:2017-4-26
//desc: 好些流氓的办法
//
//
//
bool isPowerOfThree(int n) {
    if(n <= 0) return false;
    if(n == 1) return true;
    if(n % 3) return false;
    int temp = 1;
    while(temp < n){
        temp = temp * 3;
    }
    return temp == n;
}  
    //return (n > 0) ? (1162261467 % n) == 0 : 0;
