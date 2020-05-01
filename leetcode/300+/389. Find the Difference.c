//
//date: 2017-4-23
//desc: hasn
//
//
//

char findTheDifference(char* s, char* t) {
    int length = strlen(s);
    int i = 0;
    int array[26] = {0};
    for(; i < length; i++){
        array[s[i] - 'a']--;
        array[t[i] - 'a']++;
    }
    array[t[i] - 'a']++;
    for(i = 0; i < 26; i++){
        if(array[i] == 1)
            return i+'a';
    }
    return 0;
}
