## 520. Detect Capital

<p>
Given a word, you need to judge whether the usage of capitals in it is right or not.
</p>

<p>
We define the usage of capitals in a word to be right when one of the following cases holds:
<ol>
<li>All letters in this word are capitals, like "USA".</li>
<li>All letters in this word are not capitals, like "leetcode".</li>
<li>Only the first letter in this word is capital if it has more than one letter, like "Google".</li>
</ol>
Otherwise, we define that this word doesn't use capitals in a right way.
</p>


<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> "USA"
<b>Output:</b> True
</pre>
</p>

<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> "FlaG"
<b>Output:</b> False
</pre>
</p>

<p><b>Note:</b>
The input will be a non-empty word consisting of uppercase and lowercase latin letters.
</p>