<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../guess-the-majority-in-a-hidden-array "Guess the Majority in a Hidden Array")
　　　　　　　　　　　　　　　　
[Next >](../can-convert-string-in-k-moves "Can Convert String in K Moves")

## [1539. Kth Missing Positive Number (Easy)](https://leetcode.com/problems/kth-missing-positive-number "第 k 个缺失的正整数")

<p>Given an array <code>arr</code>&nbsp;of positive integers&nbsp;sorted in a <strong>strictly increasing order</strong>, and an integer <code><font face="monospace">k</font></code>.</p>

<p><em>Find the </em><font face="monospace"><code>k<sup>th</sup></code></font><em>&nbsp;positive integer that is missing from this array.</em></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> arr = [2,3,4,7,11], k = 5
<strong>Output:</strong> 9
<strong>Explanation: </strong>The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5<sup>th</sup>&nbsp;missing positive integer is 9.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> arr = [1,2,3,4], k = 2
<strong>Output:</strong> 6
<strong>Explanation: </strong>The missing positive integers are [5,6,7,...]. The 2<sup>nd</sup> missing positive integer is 6.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= arr.length &lt;= 1000</code></li>
	<li><code>1 &lt;= arr[i] &lt;= 1000</code></li>
	<li><code>1 &lt;= k &lt;= 1000</code></li>
	<li><code>arr[i] &lt; arr[j]</code> for <code>1 &lt;= i &lt; j &lt;= arr.length</code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Binary Search](../../tag/binary-search/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Keep track of how many positive numbers are missing as you scan the array.
</details>
