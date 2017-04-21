//
//date:2017-4-21
//desc:简单的异或就可以。
//	   刚开始不能相信这么简单，有测试了几个例子，真的就是这样。
//
//
//



int singleNumber(int* nums, int numsSize) {
    
    int answer = 0;
    if(numsSize < 0 || nums == NULL)
        return 0;
    if( numsSize > 0){
        if(numsSize == 1)  return nums[0];
        answer = nums[0];
    }
    for(int i = 1;i < numsSize;i++){
        answer = answer ^ nums[i];
    }
    
    return answer;   
}
