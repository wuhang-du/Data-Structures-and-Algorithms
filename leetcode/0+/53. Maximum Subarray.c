//
//date:2017-5-23
//desc:状态变化类的题目
//		纸质分析。
//
//

#define max(a,b)  (a) > (b) ? (a) : (b)
int maxSubArray(int* nums, int numsSize) {
    int sum = nums[0];
    int tmp = sum;
    
    for(int i = 1; i < numsSize; i++){
        if(nums[i] < 0) sum = max(sum,tmp);
        if(tmp < 0)  tmp = 0;
        tmp += nums[i];
    }

	/*
        if(tmp + nums[i] > 0){
            if(tmp < 0)  tmp = 0;
            if(nums[i] < 0) sum = max(sum,tmp);
            tmp += nums[i];
        }else{
            sum = max(sum, tmp);
            tmp = nums[i];
        }
	
	*/
    return sum > tmp ? sum : tmp;
}
