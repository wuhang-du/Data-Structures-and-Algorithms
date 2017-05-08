//
//date: 2017-5-8
//desc: 简单的左右相近，
//		如果不是有序的，先排序？没想到更好的办法。
//
//

int* twoSum(int* numbers, int numbersSize, int target, int* returnSize) {
    if(numbersSize < 0 || numbers == NULL) return NULL;
    *returnSize = 2;
    int * answer = malloc(sizeof(int) * 2);
    int left = 0,right = numbersSize - 1;
    int temp = numbers[left] + numbers[right];
    for(;temp != target;temp = numbers[left] + numbers[right]) {
        if(temp < target)   left++;
        else    right--;
    }
    
    answer[0] = left + 1;
    answer[1] = right + 1;
    return answer;
}
