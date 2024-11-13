package api

import "fmt"

// 实现相关数据结构的文件

type Queue struct {
	Items []interface{}
	//MaxSize int // 最大长度
}

func (q *Queue) Init() *Queue {
	return &Queue{}
}

func (q *Queue) Add(item interface{}) {
	q.Items = append(q.Items, item)
}

// Remove 获取头元素并出队列
func (q *Queue) Remove() *interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	headItem := q.Items[0]
	q.Items = q.Items[1:]
	return &headItem
}

// Delete 直接出队列不返回
func (q *Queue) Delete() {
	if len(q.Items) == 0 {
		return
	}
	q.Items = q.Items[1:]
}

func (q *Queue) GetSize() int {
	return len(q.Items)
}

func (q *Queue) Show() {
	for _, item := range q.Items {
		fmt.Println(item)
	}
}

func (q *Queue) GetByIndex(idx int) interface{} {
	return q.Items[idx]
}
