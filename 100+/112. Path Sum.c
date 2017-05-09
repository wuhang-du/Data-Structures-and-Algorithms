//
//data: 2017-5-9
//desc: 想清楚再动手。
//
//正负不知道，所以，必须遍历到每一个叶子节点
//一堆0的问题
bool pathSum(struct TreeNode* node,int temp, int sum){
    if(node == NULL) return false;
    temp += node->val;
    
    if(node->left == NULL && node->right == NULL) 
        return (sum == temp) ? true : false;
    return pathSum(node->left,temp,sum) || pathSum(node->right,temp,sum);
    
}
bool hasPathSum(struct TreeNode* root, int sum) {
    return pathSum(root,0,sum);
}
