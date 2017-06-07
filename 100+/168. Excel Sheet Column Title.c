//
//date: 2017-6-7
//desc: boring
//

char* convertToTitle(int n) {
    char array[100];
    char count = 0;
    while (n > 0) {
        array[count++] = (n - 1) % 26;
        n = (n - 1) / 26;
    }
    char * answer = malloc(sizeof(char) * (count + 1));
    answer[count] = '\0';
    int temp = --count;
    
    for(; count >= 0; count--){
        answer[temp -count] = array[count] + 65;
    }
    return answer;
}

