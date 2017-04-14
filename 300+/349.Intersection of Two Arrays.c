
int* intersection(int* nums1, int nums1Size, int* nums2, int nums2Size, int* returnSize) {
    
    if(nums1Size < nums2Size)
    {
        int * temp = nums2;
        nums2 = nums1;
        nums1 = temp;
        
        int tempSize = nums2Size;
        nums2Size = nums1Size;
        nums1Size = tempSize;
    }
    
    int* record = malloc(sizeof(int) * nums2Size);
    memset(record,0,sizeof(int) * nums2Size); 
    
    
    if(record == NULL)
    return false;
    
    
    for(int i = 0;i < nums1Size; i++)   {
        for(int j = 0;j < nums2Size; j++)   {
            if(nums2[j] - nums1[i] == 0)    {
                record[j] = 1;
                break;
            }
        }
    }
    
    int count = 0;
    for(int j = 0;j < nums2Size; j++)
    {
        if(record[j] == 1)
        count++;
    }
    
    * returnSize = count;
    int* answer = malloc(sizeof(int) * count);
    count = 0;
    for(int j = 0;j < nums2Size; j++)   {  
        if(record[j] == 1){
            answer[count] = nums2[j];
            count++;
        }
    }
    free(record);
    return answer;
}
