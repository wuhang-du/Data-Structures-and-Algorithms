//
//
//date: 2017-4-27
//desc: 双层循环，第一个while是空白，第二个while是循环非空白。
//
//
int countSegments(char* s) {
    if(s == NULL) return 0;
    int length = strlen(s);
    char * end = s + length;
    int count = 0;
    while(s < end){
        if(*s == ' ')   s++;
        else{
            while(s < end && *s != ' ') s++;
            count++;
        }
    }
    return count;
}
