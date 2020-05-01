//
//date: 2017-6-28
//desc: 变相的先序遍历。排序二叉树的性质。
//


struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {

    if(!root) { return NULL;}
    
    if(p->val < root->val && q->val < root->val) {
        return lowestCommonAncestor(root->left,p,q);
    }
    if(p->val > root->val && q->val > root->val) {
        return lowestCommonAncestor(root->right,p,q);
    }
    return root;
}
