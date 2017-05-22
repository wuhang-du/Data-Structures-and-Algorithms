// date:2017-5-22
// desc:建立的观点错了，以为只有把所有的数都算出来，才可以确定最大的和。
// 		实际上是一个状态机。对这种题，可以试着一个数，二个数，三个数，逐渐找规律。


#define max(a,b)  (a) > (b) ? (a) : (b)
int rob(int* nums, int numsSize) {
    if(nums == NULL || numsSize <= 0) return 0;
    //对一个新的数而言，需要保存它的前一个，与前前一个状态的最大值。
    
    int left = 0; //前前一个 
    int right = 0; //前一个数
    bool sign = true;
    
    for(int i = 0; i < numsSize; sign = !sign, i++){
        if(sign)    left = max(left + nums[i], right);
        else    right = max(left, right + nums[i]);
    }
    return max(left,right);
}


int rob(int* nums, int numsSize) {
    if(nums == NULL || numsSize <= 0) return 0;
    int left = worker(nums, 0, numsSize, 0);
    int right = worker(nums, 1, numsSize, 0);
    return left > right ? left : right;
}

int worker(int* nums, int pos, int numsSize, int sum){
    if(pos >= numsSize)  return sum;
    sum = sum + nums[pos];
    int left = sum;
    int right = sum;
    pos += 2;
    if(pos < numsSize) left = worker(nums, pos, numsSize,left);
    if( ++pos < numsSize) right = worker(nums, pos, numsSize,right);
    return left > right ? left : right;
}
