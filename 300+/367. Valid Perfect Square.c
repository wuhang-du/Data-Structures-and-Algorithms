//
//data：2017-5-24
//desc: 牛顿法
//
//
//
bool isPerfectSquare(int num) {
    double k = (double)num / 2;
    double diff = 2.0;
    
    while(diff - 1.0 > 0.0){
        double temp = k;
        k = (k + num/k) / 2;
        diff = abs(temp - k);
    }
    int answer = floor(k);
    return (answer*answer) == num;
    
}
