<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../patients-with-a-condition "Patients With a Condition")
　　　　　　　　　　　　　　　　
[Next >](../minimum-suffix-flips "Minimum Suffix Flips")

## [1528. Shuffle String (Easy)](https://leetcode.com/problems/shuffle-string "重新排列字符串")

<p>You are given a string <code>s</code> and an integer array <code>indices</code> of the <strong>same length</strong>. The string <code>s</code> will be shuffled such that the character at the <code>i<sup>th</sup></code> position moves to <code>indices[i]</code> in the shuffled string.</p>

<p>Return <em>the shuffled string</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/07/09/q1.jpg" style="width: 321px; height: 243px;" />
<pre>
<strong>Input:</strong> s = &quot;codeleet&quot;, <code>indices</code> = [4,5,6,7,0,2,1,3]
<strong>Output:</strong> &quot;leetcode&quot;
<strong>Explanation:</strong> As shown, &quot;codeleet&quot; becomes &quot;leetcode&quot; after shuffling.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;abc&quot;, <code>indices</code> = [0,1,2]
<strong>Output:</strong> &quot;abc&quot;
<strong>Explanation:</strong> After shuffling, each character remains in its position.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>s.length == indices.length == n</code></li>
	<li><code>1 &lt;= n &lt;= 100</code></li>
	<li><code>s</code> consists of only lowercase English letters.</li>
	<li><code>0 &lt;= indices[i] &lt; n</code></li>
	<li>All values of <code>indices</code> are <strong>unique</strong>.</li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[String](../../tag/string/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
You can create an auxiliary string t of length n.
</details>

<details>
<summary>Hint 2</summary>
Assign t[indexes[i]] to s[i] for each i from 0 to n-1.
</details>
