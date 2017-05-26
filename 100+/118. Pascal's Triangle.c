//
//data:2017-5-27
//desc:boring
//
//





int** generate(int numRows, int** columnSizes) {
    
    if(numRows < 1) return NULL;
    int ** answer = (int**)malloc(sizeof(void*) * numRows);
    *columnSizes =  (int*) malloc(sizeof(int) * numRows);
    int * temp = malloc(sizeof(int) * ((numRows + 1)* numRows / 2));
    
    int pos = 0;
    temp[pos++] = 1;
    answer[0] = temp;
    (*columnSizes)[0] = 1;
    int prev = 0;
    
    for(int i = 2; i <= numRows; i++){
        int a = prev + 1;
        prev = pos;
        
        (*columnSizes)[i - 1] = i;
        answer[i - 1] = temp + prev;
        temp[pos++] = 1;
        for(int j = 1; j < i - 1; j++,a++){
            temp[pos++] = temp[a] + temp[a + 1];
        }
        temp[pos++] = 1;
    }
    return answer;
}
/*
1: *temp@55 = {1, 1, 1, 1, 2, 1, 1, 3, 3, 1, 1, 4, 6, 4, 1, 1, 5, 10, 10, 5, 1, 1, 6, 
  15, 20, 15, 6, 1, 1, 7, 21, 35, 35, 21, 7, 1, 1, 8, 28, 56, 70, 56, 28, 8, 1, 1, 9, 
  36, 84, 126, 126, 84, 36, 9, 1}
(gdb) display **answer@10
2: **answer@10 = {1, 1, 1, 1, 2, 1, 1, 3, 3, 1}
(gdb) display **answer
3: **answer = 1
(gdb) display *answer@10
4: *answer@10 = {0x6020a0, 0x6020a4, 0x6020ac, 0x6020b8, 0x6020c8, 0x6020dc, 0x6020f4, 
  0x602110, 0x602130, 0x602154}
(gdb) display *0x6020a0
5: *0x6020a0 = 1
(gdb) display *0x6020a4
6: *0x6020a4 = 1
(gdb) display *0x6020a4@2
7: *0x6020a4@2 = {1, 1}
(gdb) display *0x6020ac@3
8: *0x6020ac@3 = {1, 2, 1}
(gdb) display *0x6020b8@4
9: *0x6020b8@4 = {1, 3, 3, 1}
(gdb) display *0x6020c8@5
10: *0x6020c8@5 = {1, 4, 6, 4, 1}
