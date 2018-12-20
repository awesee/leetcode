## 809. Expressive Words

<p>Sometimes people repeat letters to represent extra feeling, such as &quot;hello&quot; -&gt; &quot;heeellooo&quot;, &quot;hi&quot; -&gt; &quot;hiiii&quot;.&nbsp; Here, we have&nbsp;groups, of adjacent letters that are all the same character, and adjacent characters to&nbsp;the group are different.&nbsp; A group&nbsp;is extended if that group is length 3 or more, so &quot;e&quot; and &quot;o&quot; would be extended in the first example, and &quot;i&quot; would be extended in the second example.&nbsp; As another example, the groups of &quot;abbcccaaaa&quot; would be &quot;a&quot;, &quot;bb&quot;, &quot;ccc&quot;, and &quot;aaaa&quot;; and &quot;ccc&quot; and &quot;aaaa&quot; are the extended groups of that string.</p>

<p>For some given string S, a query word is <em>stretchy</em> if it can be made to be equal to S by extending some groups.&nbsp; Formally, we are allowed to repeatedly choose a group&nbsp;(as defined above) of characters <code>c</code>, and add some number of the&nbsp;same character <code>c</code> to it so that the length of the group is 3 or more.&nbsp; Note that we cannot extend a group of size one like &quot;h&quot; to a group of size two like &quot;hh&quot; - all extensions must leave the group extended - ie., at least 3 characters long.</p>

<p>Given a list of query words, return the number of words that are stretchy.&nbsp;</p>

<pre>
<strong>Example:</strong>
<strong>Input:</strong> 
S = &quot;heeellooo&quot;
words = [&quot;hello&quot;, &quot;hi&quot;, &quot;helo&quot;]
<strong>Output:</strong> 1
<strong>Explanation:</strong> 
We can extend &quot;e&quot; and &quot;o&quot; in the word &quot;hello&quot; to get &quot;heeellooo&quot;.
We can&#39;t extend &quot;helo&quot; to get &quot;heeellooo&quot; because the group &quot;ll&quot; is not extended.
</pre>

<p><strong>Notes: </strong></p>

<ul>
	<li><code>0 &lt;= len(S) &lt;= 100</code>.</li>
	<li><code>0 &lt;= len(words) &lt;= 100</code>.</li>
	<li><code>0 &lt;= len(words[i]) &lt;= 100</code>.</li>
	<li><code>S</code> and all words in <code>words</code>&nbsp;consist only of&nbsp;lowercase letters</li>
</ul>

<p>&nbsp;</p>
