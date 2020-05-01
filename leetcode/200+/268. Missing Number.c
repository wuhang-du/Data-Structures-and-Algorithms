//
//date: 2017-5-7
//desc: 快排可以确定位置
//		异或法
//		求和会溢出。
//
//已知的信息： 一个数组，数组个数，不知道是不是缺数
//              正常的就是下标；
//              给定的数组并没有排序
//                二分 log(n)
int quicksort(int* nums,int left, int right){
    if(left > right) return left;
    
    int lefttemp = left;
    int righttemp = right--;
    int key = nums[righttemp];
    
    while(left <= right){
        while(nums[left] < key) left++;
        while(nums[right] >= key) right--;
        if(left < right){
            int temp = nums[right];
            nums[right++] = nums[left];
            nums[left++] = temp;
        }
    }
    
    nums[righttemp] = nums[left];
    if(left ==  key){
        return quicksort(nums,left+1,righttemp);
    }else
        return quicksort(nums,lefttemp,left-1);
    
}

int missingNumber(int* nums, int numsSize) {
    return quicksort(nums,0,numsSize-1);
}
