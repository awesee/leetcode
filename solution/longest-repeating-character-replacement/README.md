## 424. Longest Repeating Character Replacement (Medium)

<p>Given a string that consists of only uppercase English letters, you can replace any letter in the string with another letter at most <i>k</i> times. Find the length of a longest substring containing all repeating letters you can get after performing the above operations.</p>

<p><b>Note:</b><br />
Both the string's length and <i>k</i> will not exceed 10<sup>4</sup>.
</p>

<p>
<b>Example 1:</b>
<pre>
<b>Input:</b>
s = "ABAB", k = 2

<b>Output:</b>
4

<b>Explanation:</b>
Replace the two 'A's with two 'B's or vice versa.
</pre>
</p>

<p>
<b>Example 2:</b>
<pre>
<b>Input:</b>
s = "AABABBA", k = 1

<b>Output:</b>
4

<b>Explanation:</b>
Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
</pre>
</p>

### Similar Questions
  1. [Longest Substring with At Most K Distinct Characters](https://github.com/openset/leetcode/tree/master/solution/longest-substring-with-at-most-k-distinct-characters) (Hard)
