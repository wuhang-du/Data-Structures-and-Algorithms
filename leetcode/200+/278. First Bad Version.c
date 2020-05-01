
//date: 2017-4-16
//desc: 我不认为这道题很好。
唯一要注意的点：middle的计算。
//


// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

int firstBadVersion(int n) {
    //if(n < 1 || !isBadVersion(n) )
    //   return 0;
    
    int left = 1;
    int right = n;
    while(left < right){
        //int middle = (left+right)/2;
        int middle = left + (right -left) /2 ;
        if(isBadVersion(middle)){
            right = middle;
        } else{
            left = middle + 1 ;
        }
    }
    
    return right;
}
