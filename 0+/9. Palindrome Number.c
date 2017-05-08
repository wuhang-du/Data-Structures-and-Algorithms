//
//data: 2017-5-8
//desc: 反向的方法。刚好除了之后是反向的。
//
//


bool isPalindrome(int x) {
    if(x < 0) return false;
    
    int num = 0;
    int n = x;
    while(n){
        num = num * 10 + n % 10;
        n /= 10;
    }
    
   return num == x;
}
