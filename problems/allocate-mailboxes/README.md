<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../find-two-non-overlapping-sub-arrays-each-with-target-sum "Find Two Non-overlapping Sub-arrays Each With Target Sum")
　　　　　　　　　　　　　　　　
[Next >](../sales-by-day-of-the-week "Sales by Day of the Week")

## [1478. Allocate Mailboxes (Hard)](https://leetcode.com/problems/allocate-mailboxes "安排邮筒")

<p>Given the array <code>houses</code> where <code>houses[i]</code> is the location of the <code>i<sup>th</sup></code> house along a street and an integer <code>k</code>, allocate <code>k</code> mailboxes in the street.</p>

<p>Return <em>the <strong>minimum</strong> total distance between each house and its nearest mailbox</em>.</p>

<p>The test cases are generated so that the answer fits in a 32-bit integer.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/05/07/sample_11_1816.png" style="width: 454px; height: 154px;" />
<pre>
<strong>Input:</strong> houses = [1,4,8,10,20], k = 3
<strong>Output:</strong> 5
<strong>Explanation:</strong> Allocate mailboxes in position 3, 9 and 20.
Minimum total distance from each houses to nearest mailboxes is |3-1| + |4-3| + |9-8| + |10-9| + |20-20| = 5 
</pre>

<p><strong>Example 2:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/05/07/sample_2_1816.png" style="width: 433px; height: 154px;" />
<pre>
<strong>Input:</strong> houses = [2,3,5,12,18], k = 2
<strong>Output:</strong> 9
<strong>Explanation:</strong> Allocate mailboxes in position 3 and 14.
Minimum total distance from each houses to nearest mailboxes is |2-3| + |3-3| + |5-3| + |12-14| + |18-14| = 9.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= houses.length &lt;= 100</code></li>
	<li><code>1 &lt;= houses[i] &lt;= 10<sup>4</sup></code></li>
	<li>All the integers of <code>houses</code> are <strong>unique</strong>.</li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Math](../../tag/math/README.md)]
  [[Dynamic Programming](../../tag/dynamic-programming/README.md)]
  [[Sorting](../../tag/sorting/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
If k =1, the minimum distance is obtained allocating the mailbox in the median of the array houses.
</details>

<details>
<summary>Hint 2</summary>
Generalize this idea, using dynamic programming allocating k mailboxes.
</details>
