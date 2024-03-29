<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../line-reflection "Line Reflection")
　　　　　　　　　　　　　　　　
[Next >](../rearrange-string-k-distance-apart "Rearrange String k Distance Apart")

## [357. Count Numbers with Unique Digits (Medium)](https://leetcode.com/problems/count-numbers-with-unique-digits "计算各个位数不同的数字个数")

<p>Given an integer <code>n</code>, return the count of all numbers with unique digits, <code>x</code>, where <code>0 &lt;= x &lt; 10<sup>n</sup></code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> n = 2
<strong>Output:</strong> 91
<strong>Explanation:</strong> The answer should be the total numbers in the range of 0 &le; x &lt; 100, excluding 11,22,33,44,55,66,77,88,99
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> n = 0
<strong>Output:</strong> 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>0 &lt;= n &lt;= 8</code></li>
</ul>

### Related Topics
  [[Math](../../tag/math/README.md)]
  [[Dynamic Programming](../../tag/dynamic-programming/README.md)]
  [[Backtracking](../../tag/backtracking/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
A direct way is to use the backtracking approach.
</details>

<details>
<summary>Hint 2</summary>
Backtracking should contains three states which are (the current number, number of steps to get that number and a bitmask which represent which number is marked as visited so far in the current number). Start with state (0,0,0) and count all valid number till we reach number of steps equals to 10<sup>n</sup>.
</details>

<details>
<summary>Hint 3</summary>
This problem can also be solved using a dynamic programming approach and some knowledge of combinatorics.
</details>

<details>
<summary>Hint 4</summary>
Let f(k) = count of numbers with unique digits with length equals k.
</details>

<details>
<summary>Hint 5</summary>
f(1) = 10, ..., f(k) = 9 * 9 * 8 * ... (9 - k + 2) [The first factor is 9 because a number cannot start with 0].
</details>
