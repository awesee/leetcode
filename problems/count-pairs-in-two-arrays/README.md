<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../egg-drop-with-2-eggs-and-n-floors "Egg Drop With 2 Eggs and N Floors")
　　　　　　　　　　　　　　　　
[Next >](../determine-whether-matrix-can-be-obtained-by-rotation "Determine Whether Matrix Can Be Obtained By Rotation")

## [1885. Count Pairs in Two Arrays (Medium)](https://leetcode.com/problems/count-pairs-in-two-arrays "统计数对")



### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Binary Search](../../tag/binary-search/README.md)]
  [[Sorting](../../tag/sorting/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
We can write it as nums1[i] - nums2[i] > nums2[j] - nums1[j] instead of nums1[i] + nums1[j] > nums2[i] + nums2[j].
</details>

<details>
<summary>Hint 2</summary>
Store nums1[idx] - nums2[idx] in a data structure.
</details>

<details>
<summary>Hint 3</summary>
Store nums2[idx] - nums1[idx] in a different data structure.
</details>

<details>
<summary>Hint 4</summary>
For each integer in the first data structure, count the number of the strictly smaller integers in the second data structure with a larger index in the original array.
</details>
