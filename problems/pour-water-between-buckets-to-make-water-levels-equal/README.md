<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../earliest-possible-day-of-full-bloom "Earliest Possible Day of Full Bloom")
　　　　　　　　　　　　　　　　
[Next >](../divide-a-string-into-groups-of-size-k "Divide a String Into Groups of Size k")

## [2137. Pour Water Between Buckets to Make Water Levels Equal (Medium)](https://leetcode.com/problems/pour-water-between-buckets-to-make-water-levels-equal "")



### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Binary Search](../../tag/binary-search/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
What is the range that the answer must fall into?
</details>

<details>
<summary>Hint 2</summary>
The answer has to be in the range [0, max(buckets)] (inclusive).
</details>

<details>
<summary>Hint 3</summary>
For a number x, is there an efficient way to check if it is possible to make the amount of water in each bucket x.
</details>

<details>
<summary>Hint 4</summary>
Let in be the total amount of water that needs to be poured into buckets and out be the total amount of water that needs to be poured out of buckets to make the amount of water in each bucket x. If out - (out * loss) >= in, then it is possible.
</details>
