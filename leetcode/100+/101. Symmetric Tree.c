//
//date:2017-4-26
//desc: 递归
//    
//   看到101，想到101 空降师了，就选了这道题。
//
//
//
bool check(struct TreeNode* leftNode,struct TreeNode* rightNode){
    if(leftNode == NULL && rightNode == NULL) return true;
    if(leftNode == NULL || rightNode == NULL) return false;
    if(leftNode->val != rightNode->val) return false;
    return check(leftNode->left,rightNode->right) &&
        check(leftNode->right,rightNode->left);
}
bool isSymmetric(struct TreeNode* root) {
    if(root == NULL) return true;
    return check(root->left,root->right);
}
