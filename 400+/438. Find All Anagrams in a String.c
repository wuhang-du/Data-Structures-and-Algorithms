//
//date:2017-4-21
//desc:这道题墨迹了好久。想用窗口来做，有一些东西想的不透彻，中间改来改去。
//      说到底还是一个状态的保持的记录与还原，
//      这种东西不是花空间去提前记录，就是花时间
//      再实时计算，这是一个大的方向。
//      状态机的转换
//      本来还想造vector的轮子。忍了。


struct vector{
    int pos;
    int num;
    int* data;
};


int* findAnagrams(char* s, char* p, int* returnSize) {
    if(p == NULL) return NULL;
    int array[26] = {0};
    int temparray[26] = {0};
    
    struct vector answer = {
    .pos = 0,
    .num = 20100,
    .data = NULL,
    };
    
    
    int count = strlen(p);
    char * tempory = p;
    while(*tempory != '\0'){
        temparray[*tempory - 'a']++;
        tempory++;
    }
    
    memcpy(array,temparray,sizeof(int)*26); 
    int counttemp = count;

    answer.data = malloc(sizeof(int) * 20100);
    
    int length = strlen(s);
    
    for(int i = 0; i < length; i++){

        if(array[s[i] - 'a'] > 0){
            array[s[i] - 'a']--;
            count--;
            if(count == 0){
                int start = i - counttemp + 1;
                array[*(s+start) - 'a']  = 1;
                count++;
                answer.data[answer.pos] = start;
                answer.pos++;
            }
        }else{
            memcpy(array,temparray,sizeof(int)*26);
            count = counttemp;  //还原状态
            
            if(array[s[i] - 'a'] > 0){
                int temp = counttemp-1;
                for(int j = 0;j < temp ;j++){
                    if(array[s[i-j] - 'a' > 0]){
                        array[s[i-j] - 'a']--;
                        count--;
                    }
                }
            }
        }
    }
    *returnSize = answer.pos;
    return answer.data;
}
