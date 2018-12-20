## 394. Decode String

<p>
Given an encoded string, return it's decoded string.
</p>
<p>
The encoding rule is: <code>k[encoded_string]</code>, where the <i>encoded_string</i> inside the square brackets is being repeated exactly <i>k</i> times. Note that <i>k</i> is guaranteed to be a positive integer.</p>

<p>
You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.</p>

<p>Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, <i>k</i>. For example, there won't be input like <code>3a</code> or <code>2[4]</code>.
</p>

<p><b>Examples:</b>
<pre>
s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
</pre>
</p>

### Similar Questions
  1. [Encode String with Shortest Length](https://github.com/openset/leetcode/tree/master/solution/encode-string-with-shortest-length)(Hard)
  1. [Number of Atoms](https://github.com/openset/leetcode/tree/master/solution/number-of-atoms)(Hard)
