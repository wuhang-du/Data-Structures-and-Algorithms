//
//date: 2017-4-30
//desc: 几个问题：
//	1：for和while,for好看，看起来代码整洁，while循环速度快一些。
//	2：不要把心思放在for和while上，没有意义，这也是刷题需要注意的。
//	3：关于这个题，可以把标志位放在最后一位处理，这样逻辑就统一了。
//	   现在的代码在接口外面不友好。内存也不好释放。
//
//
// 7^9 = 40353607 > 10000000
// 最少申请10位char,再加上符号位共11位,再加上结束符 共12位。

char* convertToBase7(int num) {
    
    char flag = 1; //正数
    if(num < 0){
        flag = 0;
        num = -num;
    }
    
    char* answer = (char*)calloc(12, sizeof(char));
    
    if(num == 0){
        answer[0] = '0';
        answer[1] = '\0';
        return answer;
    }
    
    answer[0] = '-';
    char pos = 1;
    
    for(;num > 0;num /= 7)
        answer[pos++] = num % 7 + 48;
    
    answer[pos] = '\0';
    for(int i = 1;  i < pos - i; i++){
        char temp = answer[i];
        answer[i]  = answer[pos - i];
        answer[pos - i] = temp;
    }
    
    return flag ? answer + 1: answer;
    
}
