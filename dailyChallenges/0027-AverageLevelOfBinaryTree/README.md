# Day 27: Average Level of Binary Tree

## [Question](https://leetcode.com/problems/average-of-levels-in-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150)

Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.

**Example:**
![Alt text](image.png)

```
Given root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
Hence return [3, 14.5, 11].
```

## Problem-Solving Ideas
A problem discuss about levels in binary tree --> BFS approach
BFS approach:
1. use a queue and push each children onto the end of the queue
2. queue will run to the end of level before move to the next level
3. since problem requires us to isolate the level, we take the length of the queue at start of level
4. process the nodes from queue
5. as long as queue exists, we take each row (level of values) divide by length of row to find the average
6. then push the average to our result variable

## Code
## in Go 

``` Go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    var res []float64

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sum := 0.0
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			sum += float64(curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		avg := sum / float64(size)
		res = append(res, avg)
	}

	return res
}
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
    def averageOfLevels(self, root):
        """
        :type root: TreeNode
        :rtype: List[float]
        """
        q = [root]
        result = []

        while(q):
            avg = 0
            num_nodes = len(q)

            for i in range(num_nodes):
                curr = q.pop(0)
                avg += curr.val

                if curr.left:
                    q.append(curr.left)
                if curr.right:
                    q.append(curr.right)
            result.append(float(avg)/num_nodes)
        return result
        
```

## in C++
``` C++
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<double> averageOfLevels(TreeNode* root) {
        queue<TreeNode*> q;
        vector<double> res;
        q.push(root);

        while(q.size()) // q is not empty
        {
            double avg = 0;
            double node_nums = q.size();

            for(int i = 0; i < node_nums; i ++)
            {
                TreeNode* curr = q.front(); q.pop();
                avg += curr->val;

                if(curr->left)
                {
                    q.push(curr->left);
                }
                if(curr->right)
                {
                    q.push(curr->right);
                }
            }
            res.push_back(avg/node_nums);
        }
        return res;
    }
};
```



