<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../the-most-recent-three-orders "The Most Recent Three Orders")
　　　　　　　　　　　　　　　　
[Next >](../count-good-triplets "Count Good Triplets")

## [1533. Find the Index of the Large Integer (Medium)](https://leetcode.com/problems/find-the-index-of-the-large-integer "找到最大整数的索引")



### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Binary Search](../../tag/binary-search/README.md)]
  [[Interactive](../../tag/interactive/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Do a binary search over the array, exclude the half of the array that doesn't contain the largest number.
</details>

<details>
<summary>Hint 2</summary>
Keep shrinking the search space till it reaches the size of 2 where you can easily determine which one has the largest integer.
</details>
