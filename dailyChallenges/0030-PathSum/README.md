# Day 30: Path Sum
## [Question](https://leetcode.com/problems/path-sum/description/?envType=study-plan-v2&envId=top-interview-150)

Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

**Example:**
![Alt text](image.png)

```
Given root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.
```

## Problem-Solving Ideas
### DFS approach
1. check for root node, if empty return False
2. else, check for it's left node, same for checking it's right node
3. add the node value to currSum
4. if reaching either left leaf node or right leaf node, compare currSum with targetSum
5. return true if both are same, else false


## Code
## in Go 

``` Go

```

## in Python
``` python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def hasPathSum(self, root, targetSum):
        """
        :type root: TreeNode
        :type targetSum: int
        :rtype: bool
        """
        def dfs(node, currSum):
            if not node:
                return False
            currSum += node.val
            if not node.left and not node.right:
                return currSum == targetSum
            return (dfs(node.left, currSum) or dfs(node.right, currSum))
        return dfs(root, 0)
```

## in C++
``` C++

```



