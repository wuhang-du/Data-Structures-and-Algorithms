//
//date: 2017-4-25
//desc: 第一个点： 没想到strlen,和直接求值 时间差这么多，值得测试一下。
//		第二个点：这个代码虽然不好看，但是性能和原理是没问题的。可以将两个for循环整合在一起。及时跳出来。
//      第三个点：如何精简代码，一个是整合逻辑，另外一个就是运算。
//                  比如：取余。
//
//		看到配置，想起来 配置项，比如16进制，48。类似的东西应该单独出来


char* addStrings(char* num1, char* num2) {
    int lengthNum1 = 0;
    for(;num1[lengthNum1] != 0;lengthNum1++);
    int lengthNum2 = 0;
    for(;num2[lengthNum2] != 0;lengthNum2++);
    
    char * more = num1;
    char * less = num2;
    int lengthMore = lengthNum1;
    int lengthLess = lengthNum2;
    if(lengthNum1 < lengthNum2){
        more = num2;
        less = num1;
        lengthMore = lengthNum2;
        lengthLess = lengthNum1;
    }
    
    
    //char* answer = calloc(sizeof(char) * (lengthMore + 2));
    char* answer = (char*)calloc(lengthMore + 2, sizeof(char));
    //memset(answer,0,(lengthMore + 2));
    int pos = lengthMore + 1;
    answer[pos--] = '\0';

    char add = 0;
    
    int i = lengthLess - 1;
    int j = lengthMore - 1;
    
    for(; i > -1 ; i--,j--){
        char a = less[i] - 48;
        char b = more[j] - 48;
        a = a + b + add;
        //add = 0;
        //两个符号，减少了7，8行代码。
        add = a / 10;
        a = a % 10;
        /*
        if(a >= 10){
            add = 1;
            a = a - 10;
            answer[pos--] = a + 48;
        }else{
            answer[pos--] = a + 48;
        }*/
    }
    
    for(j = lengthMore - lengthLess - 1; j > -1; j--){
        char a = more[j] - 48;
        a = a + add;
        //add = 0;
        //两个符号，减少了7，8行代码。
        add = a / 10;
        a = a % 10;
        /*
        if(a >= 10){
            add = 1;
            a = a - 10;
            answer[pos--] = a + 48;
        }else{
            answer[pos--] = a + 48;
        }*/
    }
    
    if(add > 0){
        answer[0] = '1';
        return answer;
    }
    else
        return answer + 1;
    
    //return answer[0] == 0 ? answer + 1: answer;
    //return answer;
    
}
