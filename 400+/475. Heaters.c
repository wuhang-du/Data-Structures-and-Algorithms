//
//date:2017-5-15
//desc: 排序，逻辑的精简与清晰
//		换了大屏幕，看来需要练习指法了。
//
// 假设是排好序的
// 应该只是告诉大小

void quicksort(int* num, int start, int end){
    if(start >= end ) return;
    int key = num[end];
    int left = start;
    int right = end - 1;
    
    while(left <= right){
        while(num[left] < key) left++;
        while(num[right] >= key) right--;
        if(left < right){
            int temp = num[right];
            num[right++] = num[left];
            num[left++] = temp;
        }
    }
    
    num[end] = num[left];
    num[left]  = key;
    
    quicksort(num,start,left-1);
    quicksort(num,left+1,end);
    
}

int findRadius(int* houses, int housesSize, int* heaters, int heatersSize) {
    quicksort(heaters, 0, heatersSize - 1);  //排好序的蜡烛数组

    int min = 0;
    for(int i = 0; i < housesSize; i++){
        int j = 0;
        while(j < heatersSize && houses[i] >= heaters[j])   j++;
        int temp = 0;
        if(j == 0)  temp = heaters[j] - houses[i];
        else if(j == heatersSize)   temp = houses[i] - heaters[j - 1];
        else{
            temp = houses[i] - heaters[j - 1];
            int temp1 = heaters[j] - houses[i];
            temp = (temp >= temp1) ? temp1 : temp;
        }
        min = min > temp ? min : temp;
    }
    return min;
}
