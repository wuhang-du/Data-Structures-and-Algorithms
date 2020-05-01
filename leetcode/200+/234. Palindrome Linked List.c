//
//date: 2017-5-2
//desc: 噩梦一样的一道题。
//desc: 调整了一下，重新学习了一下，竟然还有递归这种想法，太厉害了。
//		递归就是个天然栈，先入后出，刚好把整个链表转过来了。
//
//
//

int result = true;
struct ListNode* temp = NULL;
void check(struct ListNode*);

bool isPalindrome(struct ListNode* head) {
    if(head == NULL) return true;
    temp = head;
    check(head);
    return result;
}

void check(struct ListNode* node){
    if(node->next)
        check(node->next);
    if(result){
        if(temp->val == node->val){
            temp = temp->next;
        }else{
            result = false;
        }
    }
}







//==============================================

bool isPalindrome(struct ListNode* head) {
    if(!head || !head->next) return true;
    
    struct ListNode* tmp = head;
    int i = 0;
    for(;tmp ; tmp = tmp ->next)  i++;
    
    tmp = head->next;
    head->next = NULL;
    int a = i;
    i >>= 1;
    for(;i > 1;i--){
        struct ListNode* p = tmp;
        tmp = tmp -> next;
        p -> next = head;
        head = p;
    }
    
    if(a % 2) tmp = tmp->next;
    struct ListNode* right = tmp;
    while(head != NULL && right != NULL){
        if(head->val == right->val){
            head = head -> next;
            right = right -> next;
        }else return false;
    }
    
    if(head == NULL && right == NULL)
        return true;
    return false;
}
