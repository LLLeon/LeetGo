0155: Min Stack

## 描述

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- getMin() -- Retrieve the minimum element in the stack.

**Example:**

```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```

## 解答

可以用一个辅助栈 minIndex 来存数据栈 dataStack 中最小值的索引。由于代码中使用的是数组栈，所有操作的时间复杂度都是 O(1)，由于使用了辅助栈，空间复杂度是 O(n)。

下面两点需要注意：

> - Push 值时，如果入栈的值在 dataStack 中的索引小于 minIndex 栈顶的值（即 dataStack 中最小值的索引），索引才 Push 进 minIndex。这样空间复杂度虽然没有降低数量级，空间占用也稍微小了一点。
> - Pop 时，如果 dataStack 中出栈的值的索引在 minIndex 的栈顶，minStack 栈顶的值才会出栈。