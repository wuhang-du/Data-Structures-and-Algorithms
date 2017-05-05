//
//
//date: 2017-5-5
//desc: 加油！
//
//
//
//
struct{
    int * data;
    int pos;
    int max;
}queue;

void push(queue* obj, int x){
    obj->data[(obj->pos)--] = x;
}

void init(queue ** objp,int maxSize){
    (*objp)->data = malloc(sizeof(int)* (maxSize+1));
    (*objp)->max = maxSize;
    (*objp)->pos = maxSize;
}

typedef struct {
    queue * first; //定义为正在使用的。
    queue * second; //空的。
} MyStack;

/** Initialize your data structure here. */
MyStack* myStackCreate(int maxSize) {
    MyStack * temp = malloc(sizeof(MyStack));
    temp->first = malloc(sizeof(queue));
    temp->second = malloc(sizeof(queue));
    init(&(temp->first),maxSize);
    init(&(temp->second),maxSize);
    /*
    temp->first->data = malloc(sizeof(int)* (maxSize+1));
    temp->first->max = maxSize;
    temp->first->pos = maxSize;
    temp->second->data = malloc(sizeof(int)* (maxSize+1));
    temp->second->max = maxSize;
    temp->second->pos = maxSize;
    */
    return temp;
}

/** Push element x onto stack. */
void myStackPush(MyStack* obj, int x) {
    push(obj->first,x);
}

/** Removes the element on top of the stack and returns that element. */
int myStackPop(MyStack* obj) {
    //把数字从first中挪出来。
    int num = obj->first->max;
    int last = obj->first->pos + 1;
    for(;num > last;num--){
        push(obj->second,obj->first->data[num]);
    }
    
    num = obj->first->data[last];
    obj->first->pos = obj->first->max;
    
    //始终保证first指向存数据的那个队列。
    queue * temp = obj->second;
    obj->second = obj->first;
    obj->first = temp;
    
    return num;
}

/** Get the top element. */
int myStackTop(MyStack* obj) {
    int temp = myStackPop(obj);
    push(obj->first,temp);
    return temp;
}

/** Returns whether the stack is empty. */
bool myStackEmpty(MyStack* obj) {
    return obj->first->max - obj->first->pos ? false : true;
}

void myStackFree(MyStack* obj) {
    free(obj->first->data);
    free(obj->second->data);
    free(obj->first);
    free(obj->second);
    free(obj);
}

/**
 * Your MyStack struct will be instantiated and called as such:
 * struct MyStack* obj = myStackCreate(maxSize);
 * myStackPush(obj, x);
 * int param_2 = myStackPop(obj);
 * int param_3 = myStackTop(obj);
 * bool param_4 = myStackEmpty(obj);
 * myStackFree(obj);
 */
