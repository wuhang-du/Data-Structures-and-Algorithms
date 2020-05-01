//
//date: 2017-4-24
//desc: 注意审题
//


bool checkRecord(char* s) {
    int a = 0, l = 0;
    while(*s != '\0'){
        while(*s == 'L'){
            l++;
            s++;
        }
        if( l > 2) return false;
        l = 0;
        if(*s++ == 'A') a++;
    }
    if(a < 2)  return true;
    else return false;
}
