<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../leftmost-column-with-at-least-a-one "Leftmost Column with at Least a One")
　　　　　　　　　　　　　　　　
[Next >](../check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree "Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree")

## [1429. First Unique Number (Medium)](https://leetcode.com/problems/first-unique-number "第一个唯一数字")



### Related Topics
  [[Array](../../tag/array/README.md)]
  [[Hash Table](../../tag/hash-table/README.md)]
  [[Design](../../tag/design/README.md)]
  [[Queue](../../tag/queue/README.md)]
  [[Data Stream](../../tag/data-stream/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Use doubly Linked list with hashmap of pointers to linked list nodes. add unique number to the linked list. When add is called check if the added number is unique then it have to be added to the linked list and if it is repeated remove it from the linked list if exists. When showFirstUnique is called retrieve the head of the linked list.
</details>

<details>
<summary>Hint 2</summary>
Use queue and check that first element of the queue is always unique.
</details>

<details>
<summary>Hint 3</summary>
Use set or heap to make running time of each function O(logn).
</details>
