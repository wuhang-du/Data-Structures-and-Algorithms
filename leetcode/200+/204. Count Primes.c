//
//date:2017-4-21
//desc: 这个题先去网上查了相关的算法，只不过自己实现了一下罢了。
//
//
//
int countPrimes(int n) {
    int* array = (int *)malloc(sizeof(int) * (n+1));
    memset(array,0,sizeof(int));
    int all = 0;
    for(int i = 2 ; i < n ; i++){
        if(array[i] == 0){
            all++;
            int count = 2;
            int temp = i * 2;
            while(temp < n){
                if(array[temp] == 0)   
                    array[temp] = 1;
                temp = i * (++count);
            }
        }
    }
    return all;
}
