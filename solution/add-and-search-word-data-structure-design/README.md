## 211. Add and Search Word - Data structure design (Medium)

<p>Design a data structure that supports the following two operations:</p>

<pre>
void addWord(word)
bool search(word)
</pre>

<p>search(word) can search a literal word or a regular expression string containing only letters <code>a-z</code> or <code>.</code>. A <code>.</code> means it can represent any one letter.</p>

<p><strong>Example:</strong></p>

<pre>
addWord(&quot;bad&quot;)
addWord(&quot;dad&quot;)
addWord(&quot;mad&quot;)
search(&quot;pad&quot;) -&gt; false
search(&quot;bad&quot;) -&gt; true
search(&quot;.ad&quot;) -&gt; true
search(&quot;b..&quot;) -&gt; true
</pre>

<p><b>Note:</b><br />
You may assume that all words are consist of lowercase letters <code>a-z</code>.</p>


### Similar Questions
  1. [Implement Trie (Prefix Tree)](https://github.com/openset/leetcode/tree/master/solution/implement-trie-prefix-tree) (Medium)
  1. [Prefix and Suffix Search](https://github.com/openset/leetcode/tree/master/solution/prefix-and-suffix-search) (Hard)
