//
//data: 2017-5-25
//
//
//
int* plusOne(int* digits, int digitsSize, int* returnSize) {
    
    int * answer = (int *) malloc(sizeof(int) * (digitsSize + 1));
    int carry = 1;
    int i = digitsSize;
    for(; i > 0; i--){
        if((digits[i - 1] + carry) == 10){
            answer[i] = 0;
            carry = 1;
        }else {
            answer[i] = digits[i - 1] + carry;
            carry = 0;
        } 
    }
    *returnSize = digitsSize + carry;
    if(carry) { answer[0] = 1; return answer;}
    return answer + 1;
}
