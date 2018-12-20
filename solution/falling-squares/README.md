## 699. Falling Squares

<p>On an infinite number line (x-axis), we drop given squares in the order they are given.</p>
<p>The <code>i</code>-th square dropped (<code>positions[i] = (left, side_length)</code>) is a square with the left-most point being <code>positions[i][0]</code> and sidelength <code>positions[i][1]</code>.</p>
<p>The square is dropped with the bottom edge parallel to the number line, and from a higher height than all currently landed squares.  We wait for each square to stick before dropping the next.</p>
<p>The squares are infinitely sticky on their bottom edge, and will remain fixed to any positive length surface they touch (either the number line or another square).  Squares dropped adjacent to each other will not stick together prematurely.</p>

<br>
<p>Return a list <code>ans</code> of heights.  Each height <code>ans[i]</code> represents the current highest height of any square we have dropped, after dropping squares represented by <code>positions[0], positions[1], ..., positions[i]</code>.
</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> [[1, 2], [2, 3], [6, 1]]
<b>Output:</b> [2, 5, 5]
<b>Explanation:</b>
<p>
After the first drop of <code>positions[0] = [1, 2]:
_aa
_aa
-------
</code>The maximum height of any square is 2.
</p><p>
After the second drop of <code>positions[1] = [2, 3]:
__aaa
__aaa
__aaa
_aa__
_aa__
--------------
</code>The maximum height of any square is 5.  
The larger square stays on top of the smaller square despite where its center
of gravity is, because squares are infinitely sticky on their bottom edge.
</p><p>
After the third drop of <code>positions[1] = [6, 1]:
__aaa
__aaa
__aaa
_aa
_aa___a
--------------
</code>The maximum height of any square is still 5.

Thus, we return an answer of <code>[2, 5, 5]</code>.
</pre>
</p>

<br>

<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> [[100, 100], [200, 100]]
<b>Output:</b> [100, 100]
<b>Explanation:</b> Adjacent squares don't get stuck prematurely - only their bottom edge can stick to surfaces.
</pre>
</p>

<p><b>Note:</b>
<li><code>1 <= positions.length <= 1000</code>.</li>
<li><code>1 <= positions[i][0] <= 10^8</code>.</li>
<li><code>1 <= positions[i][1] <= 10^6</code>.</li>
</p>

### Similar Questions
  1. [The Skyline Problem](https://github.com/openset/leetcode/tree/master/solution/the-skyline-problem)(Hard)
