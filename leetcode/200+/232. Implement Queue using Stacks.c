//
//date:2017-5-4
//desc:有一些理想化的处理。 我自己的版本，关于内存有一些很难处理的东西，就是因为将一个大小，放在两个地方处理了。
//		实际就是放在一个处理，另外申请足够的空间就行。空间够用，空间够用。 
//
//




typedef struct {
    int * out;
    int * in;
    int max;
    int out_pos;
    int in_pos;
} MyQueue;

/** Initialize your data structure here. */
MyQueue* myQueueCreate(int maxSize) {
    MyQueue * temp = malloc(sizeof(MyQueue));
    temp->out_pos = temp->in_pos = 0;
    temp->max = maxSize;
    temp->in = malloc(sizeof(int) * maxSize);
    temp->out = malloc(sizeof(int) * maxSize);
    return temp;
}

/** Push element x to the back of queue. */
void myQueuePush(MyQueue* obj, int x) {
    obj->in[obj->in_pos] = x;
    obj->in_pos++;
}

/** Removes the element from in front of queue and returns that element. */
int myQueuePop(MyQueue* obj) {
    myQueuePeek(obj);
    (obj->out_pos)--;
    return obj->out[obj->out_pos];
}

/** Get the front element. */
int myQueuePeek(MyQueue* obj) {
    if(obj->out_pos <= 0){
        while(obj->in_pos > 0){
            obj->out[(obj->out_pos)++] = obj->in[--(obj->in_pos)];
        }
    }
    return obj->out[obj->out_pos - 1];
}

/** Returns whether the queue is empty. */
bool myQueueEmpty(MyQueue* obj) {
    return (obj->in_pos == 0) && (obj->out_pos == 0);
    
}

void myQueueFree(MyQueue* obj) {
    free(obj->in);
    free(obj->out);
    free(obj);
}

/**
 * Your MyQueue struct will be instantiated and called as such:
 * struct MyQueue* obj = myQueueCreate(maxSize);
 * myQueuePush(obj, x);
 * int param_2 = myQueuePop(obj);
 * int param_3 = myQueuePeek(obj);
 * bool param_4 = myQueueEmpty(obj);
 * myQueueFree(obj);
 */
