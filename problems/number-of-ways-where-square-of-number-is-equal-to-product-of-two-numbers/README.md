<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../replace-all-s-to-avoid-consecutive-repeating-characters "Replace All ?'s to Avoid Consecutive Repeating Characters")
　　　　　　　　　　　　　　　　
[Next >](../minimum-time-to-make-rope-colorful "Minimum Time to Make Rope Colorful")

## [1577. Number of Ways Where Square of Number Is Equal to Product of Two Numbers (Medium)](https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers "数的平方等于两数乘积的方法数")

<p>Given two arrays of integers <code>nums1</code> and <code>nums2</code>, return the number of triplets formed (type 1 and type 2) under the following rules:</p>

<ul>
	<li>Type 1: Triplet (i, j, k) if <code>nums1[i]<sup>2</sup> == nums2[j] * nums2[k]</code> where <code>0 &lt;= i &lt; nums1.length</code> and <code>0 &lt;= j &lt; k &lt; nums2.length</code>.</li>
	<li>Type 2: Triplet (i, j, k) if <code>nums2[i]<sup>2</sup> == nums1[j] * nums1[k]</code> where <code>0 &lt;= i &lt; nums2.length</code> and <code>0 &lt;= j &lt; k &lt; nums1.length</code>.</li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [7,4], nums2 = [5,2,8,9]
<strong>Output:</strong> 1
<strong>Explanation:</strong> Type 1: (1, 1, 2), nums1[1]<sup>2</sup> = nums2[1] * nums2[2]. (4<sup>2</sup> = 2 * 8). 
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [1,1], nums2 = [1,1,1]
<strong>Output:</strong> 9
<strong>Explanation:</strong> All Triplets are valid, because 1<sup>2</sup> = 1 * 1.
Type 1: (0,0,1), (0,0,2), (0,1,2), (1,0,1), (1,0,2), (1,1,2).  nums1[i]<sup>2</sup> = nums2[j] * nums2[k].
Type 2: (0,0,1), (1,0,1), (2,0,1). nums2[i]<sup>2</sup> = nums1[j] * nums1[k].
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums1 = [7,7,8,3], nums2 = [1,2,9,7]
<strong>Output:</strong> 2
<strong>Explanation:</strong> There are 2 valid triplets.
Type 1: (3,0,2).  nums1[3]<sup>2</sup> = nums2[0] * nums2[2].
Type 2: (3,0,1).  nums2[3]<sup>2</sup> = nums1[0] * nums1[1].
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length, nums2.length &lt;= 1000</code></li>
	<li><code>1 &lt;= nums1[i], nums2[i] &lt;= 10<sup>5</sup></code></li>
</ul>

### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Hash Table](../../tag/hash-table/README.md)]
  [[Math](../../tag/math/README.md)]
  [[Two Pointers](../../tag/two-pointers/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Precalculate the frequencies of all nums1[i]^2 and nums2[i]^2
</details>
