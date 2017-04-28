//
//
//date:2017-4-28
//desc:递归在效率上一般，但是写起来比较优雅
//
//

struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {
    if(l1 == NULL && l2 == NULL) return NULL;
    if(l1 != NULL && l2 == NULL) return l1;
    if(l1 == NULL && l2 != NULL) return l2;
    if(l1->val <= l2->val){
        l1->next = mergeTwoLists(l1->next,l2);
        return l1;
    }else{
        l2->next = mergeTwoLists(l1,l2->next);
        return l2;
    }
}
