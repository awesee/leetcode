## 438. Find All Anagrams in a String (Easy)

<p>Given a string <b>s</b> and a <b>non-empty</b> string <b>p</b>, find all the start indices of <b>p</b>'s anagrams in <b>s</b>.</p>

<p>Strings consists of lowercase English letters only and the length of both strings <b>s</b> and <b>p</b> will not be larger than 20,100.</p>

<p>The order of output does not matter.</p>

<p><b>Example 1:</b>
<pre>
<b>Input:</b>
s: "cbaebabacd" p: "abc"

<b>Output:</b>
[0, 6]

<b>Explanation:</b>
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
</pre>
</p>

<p><b>Example 2:</b>
<pre>
<b>Input:</b>
s: "abab" p: "ab"

<b>Output:</b>
[0, 1, 2]

<b>Explanation:</b>
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
</pre>
</p>

### Similar Questions
  1. [Valid Anagram](https://github.com/openset/leetcode/tree/master/solution/valid-anagram) (Easy)
  1. [Permutation in String](https://github.com/openset/leetcode/tree/master/solution/permutation-in-string) (Medium)
