//
//date:2017-4-15
//descibe: 字符串的处理，顺序或者逆序都很重要。
//			另外就是及时退出。
//
//
//


bool detectCapitalUse(char* word) {
    int size = strlen(word);
    bool haveUpper = false;
    bool haveLower = false;
    int count = 0;
    
    if(word[size-1] < 97)
        haveUpper = true;
    else
        haveLower = true;
    for(int i = size-2;i > 0;i--)
    {
        if(haveUpper && word[i] >=97) 
            return false;
        if(haveLower && word[i] <97)
            return false;            
    }
    
    if(haveUpper && word[0] >= 97)
        return false;
    
    return true;
}
