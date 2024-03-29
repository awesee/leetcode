<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../as-far-from-land-as-possible "As Far from Land as Possible")
　　　　　　　　　　　　　　　　
[Next >](../product-price-at-a-given-date "Product Price at a Given Date")

## [1163. Last Substring in Lexicographical Order (Hard)](https://leetcode.com/problems/last-substring-in-lexicographical-order "按字典序排在最后的子串")

<p>Given a string <code>s</code>, return <em>the last substring of</em> <code>s</code> <em>in lexicographical order</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;abab&quot;
<strong>Output:</strong> &quot;bab&quot;
<strong>Explanation:</strong> The substrings are [&quot;a&quot;, &quot;ab&quot;, &quot;aba&quot;, &quot;abab&quot;, &quot;b&quot;, &quot;ba&quot;, &quot;bab&quot;]. The lexicographically maximum substring is &quot;bab&quot;.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;leetcode&quot;
<strong>Output:</strong> &quot;tcode&quot;
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 4 * 10<sup>5</sup></code></li>
	<li><code>s</code> contains only lowercase English letters.</li>
</ul>

### Related Topics
  [[Two Pointers](../../tag/two-pointers/README.md)]
  [[String](../../tag/string/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Assume that the answer is a sub-string from index i to j. If you add the character at index j+1 you get a better answer.
</details>

<details>
<summary>Hint 2</summary>
The answer is always a suffix of the given string.
</details>

<details>
<summary>Hint 3</summary>
Since the limits are high, we need an efficient data structure.
</details>

<details>
<summary>Hint 4</summary>
Use suffix array.
</details>
