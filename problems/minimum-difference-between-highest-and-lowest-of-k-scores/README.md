<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../widest-pair-of-indices-with-equal-range-sum "Widest Pair of Indices With Equal Range Sum")
　　　　　　　　　　　　　　　　
[Next >](../find-the-kth-largest-integer-in-the-array "Find the Kth Largest Integer in the Array")

## [1984. Minimum Difference Between Highest and Lowest of K Scores (Easy)](https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores "学生分数的最小差值")

<p>You are given a <strong>0-indexed</strong> integer array <code>nums</code>, where <code>nums[i]</code> represents the score of the <code>i<sup>th</sup></code> student. You are also given an integer <code>k</code>.</p>

<p>Pick the scores of any <code>k</code> students from the array so that the <strong>difference</strong> between the <strong>highest</strong> and the <strong>lowest</strong> of the <code>k</code> scores is <strong>minimized</strong>.</p>

<p>Return <em>the <strong>minimum</strong> possible difference</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [90], k = 1
<strong>Output:</strong> 0
<strong>Explanation:</strong> There is one way to pick score(s) of one student:
- [<strong><u>90</u></strong>]. The difference between the highest and lowest score is 90 - 90 = 0.
The minimum possible difference is 0.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [9,4,1,7], k = 2
<strong>Output:</strong> 2
<strong>Explanation:</strong> There are six ways to pick score(s) of two students:
- [<strong><u>9</u></strong>,<strong><u>4</u></strong>,1,7]. The difference between the highest and lowest score is 9 - 4 = 5.
- [<strong><u>9</u></strong>,4,<strong><u>1</u></strong>,7]. The difference between the highest and lowest score is 9 - 1 = 8.
- [<strong><u>9</u></strong>,4,1,<strong><u>7</u></strong>]. The difference between the highest and lowest score is 9 - 7 = 2.
- [9,<strong><u>4</u></strong>,<strong><u>1</u></strong>,7]. The difference between the highest and lowest score is 4 - 1 = 3.
- [9,<strong><u>4</u></strong>,1,<strong><u>7</u></strong>]. The difference between the highest and lowest score is 7 - 4 = 3.
- [9,4,<strong><u>1</u></strong>,<strong><u>7</u></strong>]. The difference between the highest and lowest score is 7 - 1 = 6.
The minimum possible difference is 2.</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= nums.length &lt;= 1000</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Sorting](../../tag/sorting/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
For the difference between the highest and lowest element to be minimized, the k chosen scores need to be as close to each other as possible.
</details>

<details>
<summary>Hint 2</summary>
What if the array was sorted?
</details>

<details>
<summary>Hint 3</summary>
After sorting the scores, any contiguous k scores are as close to each other as possible.
</details>

<details>
<summary>Hint 4</summary>
Apply a sliding window solution to iterate over each contiguous k scores, and find the minimum of the differences of all windows.
</details>