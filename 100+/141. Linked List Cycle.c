//
//date: 2017-4-26
//desc: 想不出更好的办法了。
//
//
//


bool hasCycle(struct ListNode *head) {
    if(head == NULL) return false;
    struct ListNode* fast = head;
    //while(fast->next && fast->next->next){
    while(fast->next && (fast = fast->next->next)){
        //fast = fast->next->next;
        head = head->next;
        if(head == fast)
        return true;
    }
    return false;
}
