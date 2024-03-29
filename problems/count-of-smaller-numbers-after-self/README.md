<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../binary-tree-vertical-order-traversal "Binary Tree Vertical Order Traversal")
　　　　　　　　　　　　　　　　
[Next >](../remove-duplicate-letters "Remove Duplicate Letters")

## [315. Count of Smaller Numbers After Self (Hard)](https://leetcode.com/problems/count-of-smaller-numbers-after-self "计算右侧小于当前元素的个数")

<p>You are given an integer array <code>nums</code> and you have to return a new <code>counts</code> array. The <code>counts</code> array has the property where <code>counts[i]</code> is the number of smaller elements to the right of <code>nums[i]</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [5,2,6,1]
<strong>Output:</strong> [2,1,1,0]
<strong>Explanation:</strong>
To the right of 5 there are <b>2</b> smaller elements (2 and 1).
To the right of 2 there is only <b>1</b> smaller element (1).
To the right of 6 there is <b>1</b> smaller element (1).
To the right of 1 there is <b>0</b> smaller element.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1]
<strong>Output:</strong> [0]
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,-1]
<strong>Output:</strong> [0,0]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>

### Related Topics
  [[Binary Indexed Tree](../../tag/binary-indexed-tree/README.md)]
  [[Segment Tree](../../tag/segment-tree/README.md)]
  [[Array](../../tag/array/README.md)]
  [[Binary Search](../../tag/binary-search/README.md)]
  [[Divide and Conquer](../../tag/divide-and-conquer/README.md)]
  [[Ordered Set](../../tag/ordered-set/README.md)]
  [[Merge Sort](../../tag/merge-sort/README.md)]

### Similar Questions
  1. [Count of Range Sum](../count-of-range-sum) (Hard)
  1. [Queue Reconstruction by Height](../queue-reconstruction-by-height) (Medium)
  1. [Reverse Pairs](../reverse-pairs) (Hard)
