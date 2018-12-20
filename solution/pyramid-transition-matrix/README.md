## 756. Pyramid Transition Matrix

<p>
We are stacking blocks to form a pyramid.  Each block has a color which is a one letter string, like `'Z'`.
</p><p>
For every block of color `C` we place not in the bottom row, we are placing it on top of a left block of color `A` and right block of color `B`.  We are allowed to place the block there only if `(A, B, C)` is an allowed triple.
</p><p>
We start with a bottom row of <code>bottom</code>, represented as a single string.  We also start with a list of allowed triples <code>allowed</code>.  Each allowed triple is represented as a string of length 3.
</p><p>
Return true if we can build the pyramid all the way to the top, otherwise false.
</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> bottom = "XYZ", allowed = ["XYD", "YZE", "DEA", "FFF"]
<b>Output:</b> true
<b>Explanation:</b>
We can stack the pyramid like this:
    A
   / \
  D   E
 / \ / \
X   Y   Z

This works because ('X', 'Y', 'D'), ('Y', 'Z', 'E'), and ('D', 'E', 'A') are allowed triples.
</pre>
</p>

<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> bottom = "XXYX", allowed = ["XXX", "XXY", "XYX", "XYY", "YXZ"]
<b>Output:</b> false
<b>Explanation:</b>
We can't stack the pyramid to the top.
Note that there could be allowed triples (A, B, C) and (A, B, D) with C != D.
</pre>
</p>

<p><b>Note:</b><br>
<ol>
<li><code>bottom</code> will be a string with length in range <code>[2, 8]</code>.</li>
<li><code>allowed</code> will have length in range <code>[0, 200]</code>.</li>
<li>Letters in all strings will be chosen from the set <code>{'A', 'B', 'C', 'D', 'E', 'F', 'G'}</code>.</li>
</ol>
</p>