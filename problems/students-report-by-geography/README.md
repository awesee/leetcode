## 618. Students Report By Geography (Hard)

A U.S graduate school has students from Asia, Europe and America. The students' location information are stored in table <code>student</code> as below.<p></p>
 
<pre>| name   | continent |
|--------|-----------|
| Jack   | America   |
| Pascal | Europe    |
| Xi     | Asia      |
| Jane   | America   |
</pre><p></p>
 
<a href="https://en.wikipedia.org/wiki/Pivot_table"> Pivot</a> the continent column in this table so that each name is sorted alphabetically and displayed underneath its corresponding continent. The output headers should be America, Asia and Europe respectively. It is guaranteed that the student number from America is no less than either Asia or Europe.<p></p>
 
For the sample input, the output is:<p></p>
<pre>| America | Asia | Europe |
|---------|------|--------|
| Jack    | Xi   | Pascal |
| Jane    |      |        |
</pre><p></p>
 
<b>Follow-up:</b> If it is unknown which continent has the most students, can you write a query to generate the student report?<p></p>
