<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../count-all-valid-pickup-and-delivery-options "Count All Valid Pickup and Delivery Options")
　　　　　　　　　　　　　　　　
[Next >](../validate-binary-tree-nodes "Validate Binary Tree Nodes")

## [1360. Number of Days Between Two Dates (Easy)](https://leetcode.com/problems/number-of-days-between-two-dates "日期之间隔几天")

<p>Write a program to count the number of days between two dates.</p>

<p>The two dates are given as strings, their format is <code>YYYY-MM-DD</code>&nbsp;as shown in the examples.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> date1 = "2019-06-29", date2 = "2019-06-30"
<strong>Output:</strong> 1
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> date1 = "2020-01-15", date2 = "2019-12-31"
<strong>Output:</strong> 15
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The given dates are valid&nbsp;dates between the years <code>1971</code> and <code>2100</code>.</li>
</ul>

### Related Topics
  [[Math](../../tag/math/README.md)]
  [[String](../../tag/string/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Create a function f(date) that counts the number of days from 1900-01-01 to date. How can we calculate the answer ?
</details>

<details>
<summary>Hint 2</summary>
The answer is just |f(date1) - f(date2)|.
</details>

<details>
<summary>Hint 3</summary>
How to construct f(date) ?
</details>

<details>
<summary>Hint 4</summary>
For each year from 1900 to year - 1 sum up 365 or 366 in case of leap years. Then sum up for each month the number of days, consider the case when the current year is leap, finally sum up the days.
</details>
