//
//date: 2017-5-2
//desc: 噩梦一样的一道题。
//
//
//
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
