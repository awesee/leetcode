<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../maximum-score-of-a-good-subarray "Maximum Score of a Good Subarray")
　　　　　　　　　　　　　　　　
[Next >](../rearrange-products-table "Rearrange Products Table")

## [1794. Count Pairs of Equal Substrings With Minimum Difference (Medium)](https://leetcode.com/problems/count-pairs-of-equal-substrings-with-minimum-difference "统计距离最小的子串对个数")



### Related Topics
  [[Greedy](../../tag/greedy/README.md)]
  [[Hash Table](../../tag/hash-table/README.md)]
  [[String](../../tag/string/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
If the chosen substrings are of size larger than 1, then you can remove all but the first character from both substrings, and you'll get equal substrings of size 1, with the same a but less j. Hence, it's always optimal to choose substrings of size 1.
</details>

<details>
<summary>Hint 2</summary>
If you choose a specific letter, then it's optimal to choose its first occurrence in firstString, and its last occurrence in secondString, to minimize j-a.
</details>