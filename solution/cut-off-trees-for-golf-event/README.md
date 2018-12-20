## 675. Cut Off Trees for Golf Event (Hard)

<p>
You are asked to cut off trees in a forest for a golf event. The forest is represented as a non-negative 2D map, in this map:
<ol>
<li><code>0</code> represents the <code>obstacle</code> can't be reached.</li>
<li><code>1</code> represents the <code>ground</code> can be walked through.</li>
<li><code>The place with number bigger than 1</code> represents a <code>tree</code> can be walked through, and this positive number represents the tree's height.</li>
</ol>
</p>

<p>
You are asked to cut off <b>all</b> the trees in this forest in the order of tree's height - always cut off the tree with lowest height first. And after cutting, the original place has the tree will become a grass (value 1). 
</p>

<p>
You will start from the point (0, 0) and you should output the minimum steps <b>you need to walk</b> to cut off all the trees. If you can't cut off all the trees, output -1 in that situation.
</p>

<p>
You are guaranteed that no two <code>trees</code> have the same height and there is at least one tree needs to be cut off.
</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> 
[
 [1,2,3],
 [0,0,4],
 [7,6,5]
]
<b>Output:</b> 6
</pre>
</p>

<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> 
[
 [1,2,3],
 [0,0,0],
 [7,6,5]
]
<b>Output:</b> -1
</pre>
</p>

<p><b>Example 3:</b><br />
<pre>
<b>Input:</b> 
[
 [2,3,4],
 [0,0,5],
 [8,7,6]
]
<b>Output:</b> 6
<b>Explanation:</b> You started from the point (0,0) and you can cut off the tree in (0,0) directly without walking.
</pre>
</p>

<p>
<b>Hint</b>: size of the given matrix will not exceed 50x50.
</p>