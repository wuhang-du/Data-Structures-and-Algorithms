//
//date: 2017-5-8
//desc: 如果修改了其中的一部分，一定要关注其他位置有没有使用该变量。
//		另外就是数组指针。
//		还有从头插入，后来者居上的问题。
//


struct node{
    int value; 
    struct node* next;
};

int distributeCandies(int* candies, int candiesSize) {
    if(candies == NULL || candiesSize <= 0) return 0;
    
    int half = candiesSize >> 1;
    struct node* array = malloc(sizeof(struct node) * half);
    struct node* arraytemp = array;
    
    struct node* pointer[100] = {0};
    
    int count = 0;
    for(int i = 0; i < candiesSize; i++){
        int a = (candies[i] >= 0) ? candies[i]: -candies[i];
        
        int index = a % 100;
        
        struct node* temp = pointer[index];
        
        while(temp != NULL){
            if(temp->value == candies[i])
                break;
            else temp = (temp)->next;
        }
        
        if(temp == NULL) {
            struct node * head = pointer[index];
            pointer[index] = array + (count++); 
            if(count == half) {break;}
            pointer[index]->next = head;
            pointer[index]->value = candies[i];
        }
    }
    free(array);
    return count;
}
