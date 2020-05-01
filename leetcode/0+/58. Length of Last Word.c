//
//date:2017-4-17
//desc:题不难，和别人的思路一样，时间却没怎么提高，很神奇。
//		数组寻址，和指针寻址速度差很多吗。
//		再试试。
//
//
int lengthOfLastWord(char* s) {
    int i = strlen(s) -1;   
    while((i > -1) && (s[i] ==' '))
        i--;
    int count = 0;
    while((i > -1) && (s[i] != ' ')){
        count++;
        i--;
    }
    return count;
}
