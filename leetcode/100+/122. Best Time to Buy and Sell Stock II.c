//
//date: 2017-5-27
//
//
int maxProfit(int* prices, int pricesSize) {
    if(prices == NULL || pricesSize < 2)  return 0;
    
    int min = prices[0];
    int max = prices[0];
    
    int answer = 0;
    
    for(int i = 1; i < pricesSize; i++){
        if(prices[i] >= max){
            max = prices[i];
        }else{
            answer += (max - min);
            min = max = prices[i];
            printf("%d %d\n",i,answer);
        }
    }
    
    return answer += (max - min) ;
}
