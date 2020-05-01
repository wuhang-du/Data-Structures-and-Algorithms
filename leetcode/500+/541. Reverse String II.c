//
//date: 2017-4-29
//desc: easy problem
//
//

void reverse(char* s, int left, int right){
    while(left < right){
        char temp = s[right];
        s[right--] = s[left];
        s[left++] = temp;
    }
}


char* reverseStr(char* s, int k) {
    if(k == 0) return NULL;
    int length = strlen(s);
    
    for(int i = 0; i < length; i += 2*k){
        if( i + k >= length) {
            reverse(s,i,length-1);
        }else   reverse(s,i,i+k-1);
    }
    return s;
}
