<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../rearrange-array-elements-by-sign "Rearrange Array Elements by Sign")
　　　　　　　　　　　　　　　　
[Next >](../maximum-good-people-based-on-statements "Maximum Good People Based on Statements")

## [2150. Find All Lonely Numbers in the Array (Medium)](https://leetcode.com/problems/find-all-lonely-numbers-in-the-array "找出数组中的所有孤独数字")

<p>You are given an integer array <code>nums</code>. A number <code>x</code> is <strong>lonely</strong> when it appears only <strong>once</strong>, and no <strong>adjacent</strong> numbers (i.e. <code>x + 1</code> and <code>x - 1)</code> appear in the array.</p>

<p>Return <em><strong>all</strong> lonely numbers in </em><code>nums</code>. You may return the answer in <strong>any order</strong>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [10,6,5,8]
<strong>Output:</strong> [10,8]
<strong>Explanation:</strong> 
- 10 is a lonely number since it appears exactly once and 9 and 11 does not appear in nums.
- 8 is a lonely number since it appears exactly once and 7 and 9 does not appear in nums.
- 5 is not a lonely number since 6 appears in nums and vice versa.
Hence, the lonely numbers in nums are [10, 8].
Note that [8, 10] may also be returned.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,5,3]
<strong>Output:</strong> [1,5]
<strong>Explanation:</strong> 
- 1 is a lonely number since it appears exactly once and 0 and 2 does not appear in nums.
- 5 is a lonely number since it appears exactly once and 4 and 6 does not appear in nums.
- 3 is not a lonely number since it appears twice.
Hence, the lonely numbers in nums are [1, 5].
Note that [5, 1] may also be returned.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Hash Table](../../tag/hash-table/README.md)]
  [[Counting](../../tag/counting/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
For a given element x, how can you quickly check if x - 1 and x + 1 are present in the array without reiterating through the entire array?
</details>

<details>
<summary>Hint 2</summary>
Use a set or a hash map.
</details>
