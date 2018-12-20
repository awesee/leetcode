## 155. Min Stack

<p>
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
<ul>
<li>
push(x) -- Push element x onto stack.
</li>
<li>
pop() -- Removes the element on top of the stack.
</li>
<li>
top() -- Get the top element.
</li>
<li>
getMin() -- Retrieve the minimum element in the stack.
</li>
</ul>
</p>

<p><b>Example:</b><br />
<pre>
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
</pre>
</p>

### Similar Questions
  1. [Sliding Window Maximum](https://github.com/openset/leetcode/tree/master/solution/sliding-window-maximum)(Hard)
  1. [Max Stack](https://github.com/openset/leetcode/tree/master/solution/max-stack)(Easy)
