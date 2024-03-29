<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../valid-palindrome-iii "Valid Palindrome III")
　　　　　　　　　　　　　　　　
[Next >](../longest-arithmetic-subsequence-of-given-difference "Longest Arithmetic Subsequence of Given Difference")

## [1217. Minimum Cost to Move Chips to The Same Position (Easy)](https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position "玩筹码")

<p>We have <code>n</code> chips, where the position of the <code>i<sup>th</sup></code> chip is <code>position[i]</code>.</p>

<p>We need to move all the chips to <strong>the same position</strong>. In one step, we can change the position of the <code>i<sup>th</sup></code> chip from <code>position[i]</code> to:</p>

<ul>
	<li><code>position[i] + 2</code> or <code>position[i] - 2</code> with <code>cost = 0</code>.</li>
	<li><code>position[i] + 1</code> or <code>position[i] - 1</code> with <code>cost = 1</code>.</li>
</ul>

<p>Return <em>the minimum cost</em> needed to move all the chips to the same position.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/08/15/chips_e1.jpg" style="width: 750px; height: 217px;" />
<pre>
<strong>Input:</strong> position = [1,2,3]
<strong>Output:</strong> 1
<strong>Explanation:</strong> First step: Move the chip at position 3 to position 1 with cost = 0.
Second step: Move the chip at position 2 to position 1 with cost = 1.
Total cost is 1.
</pre>

<p><strong>Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/08/15/chip_e2.jpg" style="width: 750px; height: 306px;" />
<pre>
<strong>Input:</strong> position = [2,2,2,3,3]
<strong>Output:</strong> 2
<strong>Explanation:</strong> We can move the two chips at position  3 to position 2. Each move has cost = 1. The total cost = 2.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> position = [1,1000000000]
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= position.length &lt;= 100</code></li>
	<li><code>1 &lt;= position[i] &lt;= 10^9</code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Math](../../tag/math/README.md)]
  [[Greedy](../../tag/greedy/README.md)]

### Similar Questions
  1. [Minimum Number of Operations to Move All Balls to Each Box](../minimum-number-of-operations-to-move-all-balls-to-each-box) (Medium)

### Hints
<details>
<summary>Hint 1</summary>
The first move keeps the parity of the element as it is.
</details>

<details>
<summary>Hint 2</summary>
The second move changes the parity of the element.
</details>

<details>
<summary>Hint 3</summary>
Since the first move is free, if all the numbers have the same parity, the answer would be zero.
</details>

<details>
<summary>Hint 4</summary>
Find the minimum cost to make all the numbers have the same parity.
</details>
