<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../coin-change "Coin Change")
　　　　　　　　　　　　　　　　
[Next >](../wiggle-sort-ii "Wiggle Sort II")

## [323. Number of Connected Components in an Undirected Graph (Medium)](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph "无向图中连通分量的数目")

<p>Given <code>n</code> nodes labeled from <code>0</code> to <code>n - 1</code> and a list of undirected edges (each edge is a pair of nodes), write a function to find the number of connected components in an undirected graph.</p>

<p><b>Example 1:</b></p>

<pre>
<strong>Input: </strong><code>n = 5</code> and <code>edges = [[0, 1], [1, 2], [3, 4]]</code>

     0          3
     |          |
     1 --- 2    4 

<strong>Output: </strong>2
</pre>

<p><b>Example 2:</b></p>

<pre>
<strong>Input: </strong><code>n = 5</code> and <code>edges = [[0, 1], [1, 2], [2, 3], [3, 4]]</code>

     0           4
     |           |
     1 --- 2 --- 3

<strong>Output:&nbsp;&nbsp;</strong>1
</pre>

<p><b>Note:</b><br />
You can assume that no duplicate edges will appear in <code>edges</code>. Since all edges are undirected, <code>[0, 1]</code> is the same as <code>[1, 0]</code> and thus will not appear together in <code>edges</code>.</p>

### Related Topics
  [[Depth-First Search](../../tag/depth-first-search/README.md)]
  [[Breadth-First Search](../../tag/breadth-first-search/README.md)]
  [[Union Find](../../tag/union-find/README.md)]
  [[Graph](../../tag/graph/README.md)]

### Similar Questions
  1. [Number of Islands](../number-of-islands) (Medium)
  1. [Graph Valid Tree](../graph-valid-tree) (Medium)
  1. [Number of Provinces](../number-of-provinces) (Medium)
