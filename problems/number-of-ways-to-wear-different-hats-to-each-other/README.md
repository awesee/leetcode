<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../check-if-a-string-can-break-another-string "Check If a String Can Break Another String")
　　　　　　　　　　　　　　　　
[Next >](../create-a-session-bar-chart "Create a Session Bar Chart")

## [1434. Number of Ways to Wear Different Hats to Each Other (Hard)](https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other "每个人戴不同帽子的方案数")

<p>There are <code>n</code> people and <code>40</code> types of hats labeled from <code>1</code> to <code>40</code>.</p>

<p>Given a 2D integer array <code>hats</code>, where <code>hats[i]</code> is a list of all hats preferred by the <code>i<sup>th</sup></code> person.</p>

<p>Return <em>the number of ways that the <code>n</code> people wear different hats to each other</em>.</p>

<p>Since the answer may be too large, return it modulo <code>10<sup>9</sup> + 7</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> hats = [[3,4],[4,5],[5]]
<strong>Output:</strong> 1
<strong>Explanation:</strong> There is only one way to choose hats given the conditions. 
First person choose hat 3, Second person choose hat 4 and last one hat 5.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> hats = [[3,5,1],[3,5]]
<strong>Output:</strong> 4
<strong>Explanation:</strong> There are 4 ways to choose hats:
(3,5), (5,3), (1,3) and (1,5)
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> hats = [[1,2,3,4],[1,2,3,4],[1,2,3,4],[1,2,3,4]]
<strong>Output:</strong> 24
<strong>Explanation:</strong> Each person can choose hats labeled from 1 to 4.
Number of Permutations of (1,2,3,4) = 24.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>n == hats.length</code></li>
	<li><code>1 &lt;= n &lt;= 10</code></li>
	<li><code>1 &lt;= hats[i].length &lt;= 40</code></li>
	<li><code>1 &lt;= hats[i][j] &lt;= 40</code></li>
	<li><code>hats[i]</code> contains a list of <strong>unique</strong> integers.</li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Dynamic Programming](../../tag/dynamic-programming/README.md)]
  [[Bit Manipulation](../../tag/bit-manipulation/README.md)]
  [[Bitmask](../../tag/bitmask/README.md)]

### Similar Questions
  1. [The Number of Good Subsets](../the-number-of-good-subsets) (Hard)

### Hints
<details>
<summary>Hint 1</summary>
Dynamic programming + bitmask.
</details>

<details>
<summary>Hint 2</summary>
dp(peopleMask, idHat) number of ways to wear different hats given a bitmask (people visited) and used hats from 1 to idHat-1.
</details>
