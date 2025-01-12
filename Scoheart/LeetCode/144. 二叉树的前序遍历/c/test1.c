/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

void preorder(struct TreeNode* root, int* result, int* index) {
    if (root == NULL) return;
    
    // 记录当前节点的值
    result[(*index)++] = root->val;
    
    // 递归遍历左子树
    preorder(root->left, result, index);
    
    // 递归遍历右子树
    preorder(root->right, result, index);
}

int* preorderTraversal(struct TreeNode* root, int* returnSize) {
    // 处理空树的情况
    if (root == NULL) {
        *returnSize = 0;
        return NULL;
    }
    
    // 分配内存空间，最多存储10000个节点
    int* result = (int*)malloc(10000 * sizeof(int));
    *returnSize = 0;
    
    // 开始前序遍历
    preorder(root, result, returnSize);
    
    return result;
}