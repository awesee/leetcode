<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../calculate-money-in-leetcode-bank "Calculate Money in Leetcode Bank")
　　　　　　　　　　　　　　　　
[Next >](../construct-the-lexicographically-largest-valid-sequence "Construct the Lexicographically Largest Valid Sequence")

## [1717. Maximum Score From Removing Substrings (Medium)](https://leetcode.com/problems/maximum-score-from-removing-substrings "删除子字符串的最大得分")

<p>You are given a string <code>s</code> and two integers <code>x</code> and <code>y</code>. You can perform two types of operations any number of times.</p>

<ul>
	<li>Remove substring <code>&quot;ab&quot;</code> and gain <code>x</code> points.
	<ul>
		<li>For example, when removing <code>&quot;ab&quot;</code> from <code>&quot;c<u>ab</u>xbae&quot;</code> it becomes <code>&quot;cxbae&quot;</code>.</li>
	</ul>
	</li>
	<li>Remove substring <code>&quot;ba&quot;</code> and gain <code>y</code> points.
	<ul>
		<li>For example, when removing <code>&quot;ba&quot;</code> from <code>&quot;cabx<u>ba</u>e&quot;</code> it becomes <code>&quot;cabxe&quot;</code>.</li>
	</ul>
	</li>
</ul>

<p>Return <em>the maximum points you can gain after applying the above operations on</em> <code>s</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;cdbcbbaaabab&quot;, x = 4, y = 5
<strong>Output:</strong> 19
<strong>Explanation:</strong>
- Remove the &quot;ba&quot; underlined in &quot;cdbcbbaaa<u>ba</u>b&quot;. Now, s = &quot;cdbcbbaaab&quot; and 5 points are added to the score.
- Remove the &quot;ab&quot; underlined in &quot;cdbcbbaa<u>ab</u>&quot;. Now, s = &quot;cdbcbbaa&quot; and 4 points are added to the score.
- Remove the &quot;ba&quot; underlined in &quot;cdbcb<u>ba</u>a&quot;. Now, s = &quot;cdbcba&quot; and 5 points are added to the score.
- Remove the &quot;ba&quot; underlined in &quot;cdbc<u>ba</u>&quot;. Now, s = &quot;cdbc&quot; and 5 points are added to the score.
Total score = 5 + 4 + 5 + 5 = 19.</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;aabbaaxybbaabb&quot;, x = 5, y = 4
<strong>Output:</strong> 20
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= x, y &lt;= 10<sup>4</sup></code></li>
	<li><code>s</code> consists of lowercase English letters.</li>
</ul>

### Related Topics
  [[String](../../tag/string/README.md)]
  [[Stack](../../tag/stack/README.md)]
  [[Greedy](../../tag/greedy/README.md)]

### Similar Questions
  1. [Count Words Obtained After Adding a Letter](../count-words-obtained-after-adding-a-letter) (Medium)

### Hints
<details>
<summary>Hint 1</summary>
Note that it is always more optimal to take one type of substring before another
</details>

<details>
<summary>Hint 2</summary>
You can use a stack to handle erasures
</details>
