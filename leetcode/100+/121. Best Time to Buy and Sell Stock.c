//
//date： 2017-5-27
//desc:  想的从后往前，都一样。
//
//
int maxProfit(int* prices, int pricesSize) {
  if(prices == NULL || pricesSize < 2) return 0;
  int maxNum = prices[pricesSize - 1];
  int answer = 0;
  for(int i = pricesSize - 2; i > -1; i--){
      int temp = maxNum - prices[i];
      if(temp <= 0) maxNum = prices[i];
      else  answer = answer > temp ? answer : temp;
  }
  return answer;
}
