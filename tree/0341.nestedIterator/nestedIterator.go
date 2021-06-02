package main

// ---------------------------- LeetCode 中实现的方法 ----------------------------------

type NestedInteger struct{}

func (ni NestedInteger) IsInteger() bool {
	return false
}

func (ni NestedInteger) GetInteger() int {
	return 0
}

func (ni *NestedInteger) SetInteger(val int) {
}

func (ni *NestedInteger) Add(elem NestedInteger) {
}

func (ni NestedInteger) GetList() []*NestedInteger {
	return nil
}

// -------------------------------------------- 递归实现 --------------------------------------------

type NestedIterator struct {
	vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var (
		vals []int
		dfs  func([]*NestedInteger)
	)

	dfs = func(nestedList []*NestedInteger) {
		for _, nest := range nestedList {
			if nest.IsInteger() {
				vals = append(vals, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}

	dfs(nestedList)
	return &NestedIterator{vals: vals}
}

func (ni *NestedIterator) Next() int {
	val := ni.vals[0]
	ni.vals = ni.vals[1:]
	return val
}

func (ni *NestedIterator) HasNext() bool {
	return len(ni.vals) > 0
}

// -------------------------------------------- 栈实现 --------------------------------------------
type NestedIteratorII struct {
	// 栈里面存的是队列
	stack [][]*NestedInteger
}

func ConstructorII(nestedList []*NestedInteger) *NestedIteratorII {
	return &NestedIteratorII{[][]*NestedInteger{nestedList}}
}

// 注意在使用时, 先调用 HasNext 方法再调用 Next 方法:
//   - 如果是 int, 调用 Next 方法会返回该 int.
//   - 如果不是 int, HasNext 方法会将该列表入栈, 基于栈后入先出的特性, 下次调用 HasNext 方法会先处理此列表,
// 这也符合题目「拍平」nestedList 的要求.
func (ni *NestedIteratorII) HasNext() bool {
	for len(ni.stack) > 0 {
		// 查看栈顶的队列, 如果队列为空则将其出栈
		queue := ni.stack[len(ni.stack)-1]
		if len(queue) == 0 {
			ni.stack = ni.stack[:len(ni.stack)-1]
			continue
		}

		// 判断队头元素是不是整数
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}

		// 队头元素是列表, 将其弹出队列, 再将队列入栈, 下次循环会处理栈顶队列
		ni.stack[len(ni.stack)-1] = queue[1:]
		ni.stack = append(ni.stack, nest.GetList())
	}

	return false
}

func (ni *NestedIteratorII) Next() int {
	// 直接返回栈顶队列的队首元素, 将其弹出队列并返回
	quque := ni.stack[len(ni.stack)-1]
	val := quque[0].GetInteger()
	ni.stack[len(ni.stack)-1] = quque[1:]

	return val
}
