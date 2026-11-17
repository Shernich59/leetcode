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
        