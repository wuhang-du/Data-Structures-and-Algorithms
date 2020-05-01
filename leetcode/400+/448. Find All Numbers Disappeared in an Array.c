//
//date: 2017-4-23
//
//desc: 不能动用新的内存，因此，采用了下标移位的原理，即对应的下标在对应的位置上。
//		第一版，采用了新的状态变量，因此可以在赋值阶段有效的统计计数，减少了一次遍历。
//		第二版，在不改变原有数组的基础上，增加了一次遍历。以下是第二版。
//


int* findDisappearedNumbers(int* nums, int numsSize, int* returnSize) {
    
    int count = 0;
    
    for(int i = 0;i < numsSize; i++){
        if(nums[i] == i+1)  continue;
        int temp = nums[i] - 1;  //temp是下标
        while(nums[temp] != nums[i]){
            nums[i] = nums[temp];
            nums[temp] = temp + 1;
            if(nums[i] != i+1) temp = nums[i] - 1;
            else break;
        }
    }
    
    for(int j = 0; j < numsSize; j++){
        if(nums[j] != j + 1)
            count++;
    }
    *returnSize = count;
    int* answer = (int*) malloc(sizeof(int) * count);
    count = 0;
    for(int j = 0; j< numsSize; j++){
        if(nums[j] != j + 1){
            answer[count++] = j + 1;
        }
    }
    return answer;
}
