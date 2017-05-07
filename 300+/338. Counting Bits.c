//
//
//date: 2017-5-7
//desc: 减1的迭代解法
//
//
int* countBits(int num, int* returnSize) {
    int * answer = malloc(sizeof(int) * (num+1));
    int pos = 0;
    answer[pos++] = 0;
    while(pos <= num){
        int index = pos & (pos-1);
        answer[pos++] = answer[index] + 1;
    }
    *returnSize = pos;
    return answer;
}
