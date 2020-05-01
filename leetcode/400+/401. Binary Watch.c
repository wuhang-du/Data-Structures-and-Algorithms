//
//date: 2017-5-4
//desc: c搞字符串真是太痛苦了，心累。
//		两种办法：1.是图示的下面的这种。2.是12*60 循环，看每一次是不是满足字节数是 n.
//
//
//

public class Solution {
		String[][] hour = {{"0"}, 
                                   {"1", "2", "4", "8"},
				   {"3", "5", "6", "9", "10"},
				   {"7", "11"}};
		String[][] minute = {{"00"}, //1
			             {"01", "02", "04", "08", "16", "32"}, //6
			             {"03", "05", "06", "09", "10", "12", "17", "18", "20", "24", "33", "34", "36", "40", "48"}, //15
			             {"07", "11", "13", "14", "19", "21", "22", "25", "26", "28", "35", "37", "38", "41", "42", "44", "49", "50", "52", "56"}, //20
			             {"15", "23", "27", "29", "30", "39", "43", "45", "46", "51", "53", "54", "57", "58"}, //14
			             {"31", "47", "55", "59"}}; //4
    public List<String> readBinaryWatch(int num) {
		List<String> ret = new ArrayList();
		for (int i = 0; i <= 3 && i <= num; i++) {
			if (num - i <= 5) {
				for (String str1 : hour[i]) {
					for (String str2 : minute[num - i]) {
						ret.add(str1 + ":" + str2);
					}
				}
			}
		}
		return ret;     
    }
}


 
/* 
char numofBitOne(char num){
    if(num == 0) return 0;
    char count = 0;
    while(num & (num -1)){
        count++;
        num &= (num - 1); 
    }
    return count;
} 
 
char** readBinaryWatch(int num, int* returnSize) {
    char hour[4][10] = {0};
    for(int i = 0; i < 12; i++){
        int temp = numofBitOne(i);
        hour[temp][0]++;
        hour[temp][hour[temp][0]] = i;
    }
    
    char minute[6][30] = {0};
    for(int i = 0; i < 60; i++){
        char temp = numofBitOne(i);
        minute[temp][0]++;
        minute[temp][minute[temp][0]] = i;
    }
    
    char answer[300][6] = {0};
    int pos = 0;
    
    int i = 0;
    int j = num - i;
    for(; j > -1 && i <= 3; i++){
        for(int k = 1; k <= hour[i][0]; k++){
            for(int m = 1; m <= minute[j][0];m++){
                char index = 0;
                //hour[i][k]
                if(hour[i][k] > 9){
                    answer[pos][index++] = '1';
                    answer[pos][index++] = hour[i][k] % 10 + 48;
                }else answer[pos][index++] = hour[i][k] % 10 + 48;
                answer[pos][index++] = ':';
                answer[pos][index++] = minute[j][m] / 10 + 48;      
                answer[pos][index++] = minute[j][m] % 10 + 48;
                answer[pos++][index++] = '\0';
            }
        }
        j = num - i;
    }
    *returnSize = pos;
    return answer;
}

*/
