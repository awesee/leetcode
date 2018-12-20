## 672. Bulb Switcher II

<p>
There is a room with <code>n</code> lights which are turned on initially and 4 buttons on the wall. After performing exactly <code>m</code> unknown operations towards buttons, you need to return how many different kinds of status of the <code>n</code> lights could be.
</p>

<p>
Suppose <code>n</code> lights are labeled as number [1, 2, 3 ..., n], function of these 4 buttons are given below:

<ol>
<li>Flip all the lights.</li>
<li>Flip lights with even numbers.</li>
<li>Flip lights with odd numbers.</li>
<li>Flip lights with (3k + 1) numbers, k = 0, 1, 2, ...</li>
</ol>
</p>


<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> n = 1, m = 1.
<b>Output:</b> 2
<b>Explanation:</b> Status can be: [on], [off]
</pre>
</p>


<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> n = 2, m = 1.
<b>Output:</b> 3
<b>Explanation:</b> Status can be: [on, off], [off, on], [off, off]
</pre>
</p>


<p><b>Example 3:</b><br />
<pre>
<b>Input:</b> n = 3, m = 1.
<b>Output:</b> 4
<b>Explanation:</b> Status can be: [off, on, off], [on, off, on], [off, off, off], [off, on, on].
</pre>
</p>

<p><b>Note:</b>
<code>n</code> and <code>m</code> both fit in range [0, 1000].
</p>


### Similar Questions
  1. [Bulb Switcher](https://github.com/openset/leetcode/tree/master/solution/bulb-switcher)(Medium)
