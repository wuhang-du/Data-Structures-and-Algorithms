//
//date:2017-4-30
//desc:如果有重复的代码，则一定有优化的空间。
//		手写了一个快排，比心。刚开始的判定条件有问题，其他的都还ok.
//
//

void quicksort(int* nums, int left, int right){
    if(left < 0 || left >= right) return;
    int end = right;
    int start = left;
    int num = nums[right--];
    while(left <= right){
        while(nums[right] >= num) right--;
        while(nums[left] < num) left++;
        if(left < right){
            int temp = nums[right];
            nums[right--] = nums[left];
            nums[left++] = temp;
        }else break;
    }
    
    int temp = nums[end];
    nums[end] = nums[left];
    nums[left] = temp;
    
    quicksort(nums,start,left-1);
    quicksort(nums,left+1,end);
    
}


int findPairs(int* nums, int numsSize, int k) {
    if(nums == NULL || numsSize < 0) return 0;
    quicksort(nums,0,numsSize - 1);
    int answer = 0;
    for(int i = 0; i < numsSize - 1; i++){
        int target = nums[i] + k;
        int temp = i + 1;
        while(temp < numsSize && nums[temp] <= target){
            if(nums[temp++] == target){
                answer++;
                break;
            }
        }
        while(nums[i] == nums[i+1]) i++;
        
        /*
        if(k != 0){
            while(nums[i] == nums[i+1]) i++;
            int target = nums[i] + k;
            int temp = i + 1;
            while(temp < numsSize && nums[temp] <= target){
                if(nums[temp++] == target){
                    answer++;
                    break;
                }
            }
        } else{
            if(nums[i] == nums[i+1]) answer++;
            while(nums[i] == nums[i+1]) i++;
        }
        */
    }
    return answer;
}
