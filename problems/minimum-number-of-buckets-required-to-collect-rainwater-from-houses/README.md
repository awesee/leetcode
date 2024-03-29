<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../count-common-words-with-one-occurrence "Count Common Words With One Occurrence")
　　　　　　　　　　　　　　　　
[Next >](../minimum-cost-homecoming-of-a-robot-in-a-grid "Minimum Cost Homecoming of a Robot in a Grid")

## [2086. Minimum Number of Buckets Required to Collect Rainwater from Houses (Medium)](https://leetcode.com/problems/minimum-number-of-buckets-required-to-collect-rainwater-from-houses "从房屋收集雨水需要的最少水桶数")

<p>You are given a <b>0-index</b><strong>ed</strong> string <code>street</code>. Each character in <code>street</code> is either <code>&#39;H&#39;</code> representing a house or <code>&#39;.&#39;</code> representing an empty space.</p>

<p>You can place buckets on the <strong>empty spaces</strong> to collect rainwater that falls from the adjacent houses. The rainwater from a house at index <code>i</code> is collected if a bucket is placed at index <code>i - 1</code> <strong>and/or</strong> index <code>i + 1</code>. A single bucket, if placed adjacent to two houses, can collect the rainwater from <strong>both</strong> houses.</p>

<p>Return <em>the <strong>minimum </strong>number of buckets needed so that for <strong>every</strong> house, there is <strong>at least</strong> one bucket collecting rainwater from it, or </em><code>-1</code><em> if it is impossible.</em></p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> street = &quot;H..H&quot;
<strong>Output:</strong> 2
<strong>Explanation:</strong>
We can put buckets at index 1 and index 2.
&quot;H..H&quot; -&gt; &quot;HBBH&quot; (&#39;B&#39; denotes where a bucket is placed).
The house at index 0 has a bucket to its right, and the house at index 3 has a bucket to its left.
Thus, for every house, there is at least one bucket collecting rainwater from it.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> street = &quot;.H.H.&quot;
<strong>Output:</strong> 1
<strong>Explanation:</strong>
We can put a bucket at index 2.
&quot;.H.H.&quot; -&gt; &quot;.HBH.&quot; (&#39;B&#39; denotes where a bucket is placed).
The house at index 1 has a bucket to its right, and the house at index 3 has a bucket to its left.
Thus, for every house, there is at least one bucket collecting rainwater from it.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> street = &quot;.HHH.&quot;
<strong>Output:</strong> -1
<strong>Explanation:</strong>
There is no empty space to place a bucket to collect the rainwater from the house at index 2.
Thus, it is impossible to collect the rainwater from all the houses.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= street.length &lt;= 10<sup>5</sup></code></li>
	<li><code>street[i]</code> is either<code>&#39;H&#39;</code> or <code>&#39;.&#39;</code>.</li>
</ul>

### Related Topics
  [[Greedy](../../tag/greedy/README.md)]
  [[String](../../tag/string/README.md)]
  [[Dynamic Programming](../../tag/dynamic-programming/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
When is it impossible to collect the rainwater from all the houses?
</details>

<details>
<summary>Hint 2</summary>
When one or more houses do not have an empty space adjacent to it.
</details>

<details>
<summary>Hint 3</summary>
Assuming the rainwater from all previous houses is collected. If there is a house at index i and you are able to place a bucket at index i - 1 or i + 1, where should you put it?
</details>

<details>
<summary>Hint 4</summary>
It is always better to place a bucket at index i + 1 because it can collect the rainwater from the next house as well.
</details>
