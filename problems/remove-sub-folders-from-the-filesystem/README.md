<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    awesee <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/awesee                                 |-->
<!--|@home      https://github.com/awesee/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](../check-if-it-is-a-straight-line "Check If It Is a Straight Line")
　　　　　　　　　　　　　　　　
[Next >](../replace-the-substring-for-balanced-string "Replace the Substring for Balanced String")

## [1233. Remove Sub-Folders from the Filesystem (Medium)](https://leetcode.com/problems/remove-sub-folders-from-the-filesystem "删除子文件夹")

<p>Given a list of folders <code>folder</code>, return <em>the folders after removing all <strong>sub-folders</strong> in those folders</em>. You may return the answer in <strong>any order</strong>.</p>

<p>If a <code>folder[i]</code> is located within another <code>folder[j]</code>, it is called a <strong>sub-folder</strong> of it.</p>

<p>The format of a path is one or more concatenated strings of the form: <code>&#39;/&#39;</code> followed by one or more lowercase English letters.</p>

<ul>
	<li>For example, <code>&quot;/leetcode&quot;</code> and <code>&quot;/leetcode/problems&quot;</code> are valid paths while an empty string and <code>&quot;/&quot;</code> are not.</li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> folder = [&quot;/a&quot;,&quot;/a/b&quot;,&quot;/c/d&quot;,&quot;/c/d/e&quot;,&quot;/c/f&quot;]
<strong>Output:</strong> [&quot;/a&quot;,&quot;/c/d&quot;,&quot;/c/f&quot;]
<strong>Explanation:</strong> Folders &quot;/a/b/&quot; is a subfolder of &quot;/a&quot; and &quot;/c/d/e&quot; is inside of folder &quot;/c/d&quot; in our filesystem.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> folder = [&quot;/a&quot;,&quot;/a/b/c&quot;,&quot;/a/b/d&quot;]
<strong>Output:</strong> [&quot;/a&quot;]
<strong>Explanation:</strong> Folders &quot;/a/b/c&quot; and &quot;/a/b/d/&quot; will be removed because they are subfolders of &quot;/a&quot;.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> folder = [&quot;/a/b/c&quot;,&quot;/a/b/ca&quot;,&quot;/a/b/d&quot;]
<strong>Output:</strong> [&quot;/a/b/c&quot;,&quot;/a/b/ca&quot;,&quot;/a/b/d&quot;]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= folder.length &lt;= 4 * 10<sup>4</sup></code></li>
	<li><code>2 &lt;= folder[i].length &lt;= 100</code></li>
	<li><code>folder[i]</code> contains only lowercase letters and <code>&#39;/&#39;</code>.</li>
	<li><code>folder[i]</code> always starts with the character <code>&#39;/&#39;</code>.</li>
	<li>Each folder name is <strong>unique</strong>.</li>
</ul>

### Related Topics
  [[Trie](../../tag/trie/README.md)]
  [[Array](../../tag/array/README.md)]
  [[String](../../tag/string/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
Sort the folders lexicographically.
</details>

<details>
<summary>Hint 2</summary>
Insert the current element in an array and then loop until we get rid of all of their subfolders, repeat this until no element is left.
</details>
