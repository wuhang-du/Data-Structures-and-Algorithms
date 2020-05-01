//
//
//date:2017-5-2
//desc: 1.规律找出来了，但是内存那块不知道怎么处理。
//			没想到呀，心累。
//			2倍的长度+1
//
//
char* countAndSay(int n) {
    if(n <= 0) return NULL;
    char * cur = malloc(sizeof(char) * 2);
    cur[0] = '1';
    cur[1] = '\0';
    
    for(int i = 1; i < n; i++){
        char * temp = malloc(sizeof(char) * 2 * strlen(cur) + 1);
        char * old = cur;
        int pos = 0;
        while(*cur != '\0'){
            char num = *cur;
            int count = 49;
            while(*(cur+1) != '\0' && *cur == *(cur+1) ){
                    count++;
                    cur++;
            }
            cur++;
            temp[pos++] = count;
            temp[pos++] = num;
        }
        temp[pos] = '\0';
        free(old);
        cur = temp;
    }
    return cur;
}
