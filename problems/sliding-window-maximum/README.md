<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../product-of-array-except-self "Product of Array Except Self")
　　　　　　　　　　　　　　　　
[Next >](../search-a-2d-matrix-ii "Search a 2D Matrix II")

## [239. Sliding Window Maximum (Hard)](https://leetcode.com/problems/sliding-window-maximum "滑动窗口最大值")

<p>You are given an array of integers&nbsp;<code>nums</code>, there is a sliding window of size <code>k</code> which is moving from the very left of the array to the very right. You can only see the <code>k</code> numbers in the window. Each time the sliding window moves right by one position.</p>

<p>Return <em>the max sliding window</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,-1,-3,5,3,6,7], k = 3
<strong>Output:</strong> [3,3,5,5,6,7]
<strong>Explanation:</strong> 
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       <strong>3</strong>
 1 [3  -1  -3] 5  3  6  7       <strong>3</strong>
 1  3 [-1  -3  5] 3  6  7      <strong> 5</strong>
 1  3  -1 [-3  5  3] 6  7       <strong>5</strong>
 1  3  -1  -3 [5  3  6] 7       <strong>6</strong>
 1  3  -1  -3  5 [3  6  7]      <strong>7</strong>
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1], k = 1
<strong>Output:</strong> [1]
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,-1], k = 1
<strong>Output:</strong> [1,-1]
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> nums = [9,11], k = 2
<strong>Output:</strong> [11]
</pre>

<p><strong>Example 5:</strong></p>

<pre>
<strong>Input:</strong> nums = [4,-2], k = 2
<strong>Output:</strong> [4]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= k &lt;= nums.length</code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Queue](../../tag/queue/README.md)]
  [[Sliding Window](../../tag/sliding-window/README.md)]
  [[Heap (Priority Queue)](../../tag/heap-priority-queue/README.md)]
  [[Monotonic Queue](../../tag/monotonic-queue/README.md)]

### Similar Questions
  1. [Minimum Window Substring](../minimum-window-substring) (Hard)
  1. [Min Stack](../min-stack) (Easy)
  1. [Longest Substring with At Most Two Distinct Characters](../longest-substring-with-at-most-two-distinct-characters) (Medium)
  1. [Paint House II](../paint-house-ii) (Hard)
  1. [Jump Game VI](../jump-game-vi) (Medium)

### Hints
<details>
<summary>Hint 1</summary>
How about using a data structure such as deque (double-ended queue)?
</details>

<details>
<summary>Hint 2</summary>
The queue size need not be the same as the window’s size.
</details>

<details>
<summary>Hint 3</summary>
Remove redundant elements and the queue should store only elements that need to be considered.
</details>
