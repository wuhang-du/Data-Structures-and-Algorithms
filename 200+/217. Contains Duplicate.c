//
//date: 2017-4-16
//describe: 纸上得来终觉浅，绝知此事要躬行。
//	简答的开链方式的哈希表，昨天想到这个，但是不知道怎么去实现。看了讨论，最后实现了。
//			
//
struct Node{
    int val;
    struct Node* next;
};

struct Set{
    int bucketSize;
    struct Node** table;
};

void initSet(struct Set* myset, int size){
    myset->bucketSize = size;
    myset->table = malloc(sizeof(struct Node*) * size);
    memset(myset->table,0,sizeof(struct Node*) * size);
}

bool addValue(struct Set* myset, int value){
    int idx = value > 0 ? value: -value;
    idx = idx % (myset->bucketSize);

    struct Node* temp = myset->table[idx];
    
    
    while(temp!=NULL){
        if(temp->val == value)
            return false;
        temp = temp->next;
    }

    temp = malloc(sizeof(struct Node));
    memset(temp, 0, sizeof(struct Node));
    
    temp->next = myset->table[idx];
    temp->val = value;
    myset->table[idx] = temp;
    return true;
}

void releaseSet(struct Set* myset, int size){

    for(int i = 0; i < size; i++){
        struct Node* temp = myset->table[i];
        if(temp == NULL)
            continue;
        while(temp){
            struct Node* a = temp->next;
            free(temp);
            temp = a;
        }
    }
    free(myset->table);
}

bool containsDuplicate(int* nums, int numsSize) {

    if(nums == NULL || numsSize < 1)
        return false;
    struct Set my_set;
    initSet(&my_set,numsSize);
    for(int i = 0; i < numsSize; i++ ){
        if(!addValue(&my_set,nums[i]))
            return true;
    }
    releaseSet(&my_set,numsSize);
    return false;
}
