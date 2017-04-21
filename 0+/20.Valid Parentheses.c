//
//date:2017-4-22
//desc: 刚开始用的 数组索引，后来觉得不好看，
//		又使用了直接写的办法，逻辑上也更统一。

bool isValid(char* s) {  
    int length = strlen(s);
    if(length % 2) return false;
    int count = length/2;
    char* stack = malloc(sizeof(char)*(length/2));
    int pos = -1;


    /*
    
    for(int i=1;i<length;i++){
        if(s[i] == temp){
            pos--;
            if(pos > -1 && (stack[pos] != 93 || stack[pos] != 125
                            || stack[pos] != 41))
                temp = check[stack[pos]/50];
            else temp = 0;
        }else{
            if(++pos > count) return false;
            stack[pos] = s[i];
            if(s[i] == 93 || s[i] == 125 || s[i] == 41)
                temp = 0;
            else temp = check[s[i]/50];
        }
        
    }
    */
    
    for(int i=0;i<length;i++){
        if(pos < 0){
            stack[++pos] = s[i];
        }else{
            if((stack[pos] == 40 && s[i] == 41) ||
               (stack[pos] == 91 && s[i] == 93) ||
               (stack[pos] == 123 && s[i] == 125)){
                pos--;
            } else{
               if(++pos > count) return false;
                    stack[pos] = s[i]; 
            }
        }
    }
     return pos == -1 ? true : false;
}
