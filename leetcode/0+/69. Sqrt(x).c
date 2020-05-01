//
//date:2017-5-24
//desc: 牛顿法
//
//
int mySqrt(int x) {
    if(x <= 0) return 0;
    long k = x ;
    
    while(k * k > x){
        k = (k + x / k)/2;
    }
    return k;
    /*
    double k = (double)x / 2;
    double diff = 2.0;
    
    while(diff - 0.9 > 0.0){
        double temp = k;
        k = (k + x/k) / 2;
        diff = abs(temp - k);
    }
    printf("%f \n",k);
    return floor(k);
    */
}
