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

// ----------------------------------- 以下需要实现 -----------------------------------------------

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
