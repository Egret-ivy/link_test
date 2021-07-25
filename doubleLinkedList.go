package doubleLinkedList

import (
	"errors"
	"fmt"
)

type ElemType interface{}

//链表内的每个结点
type Node struct {
	Data ElemType
	Pre  *Node
	Next *Node
}

//指整个链表，用头车厢尾车厢和整体长度表示
type List struct {
	First *Node
	Last  *Node
	Size  int
}

//工厂函数
func CreateList() *List {
	s := new(Node)
	s.Next, s.Pre = s, s
	return &List{s, s, 0} //返回链表结构体 头车厢s尾s长0
}

//尾插法
func (list *List) PushBack(x ElemType) {
	s := new(Node) //new()分配空间返回地址 创建Node类型的s 也就是一个节点
	s.Data = x

	/* 关于创建出的这个新节点要接在最后
	先连自己的线（空的，不会出现覆盖）
	s.Pre指向前一个节点，这里前一个节点怎么找呢
	因为刚好前一个就是原来的Last 就可以直接定位
	s.Pre=list.Last

	接下来关于给s授位为Last时就要大大注意啦
	先要改变原来的Last.next指向s 再传位给s 否则Last代表的东西就变成s了
	*/
	s.Pre = list.Last
	list.Last.Next = s

	list.Last = s

	//剩下的这部分就只有Last跟First之间的回合啦
	//总结出的连线技巧--一次把同一个区的两条线（向前和向后的）牵好
	list.Last.Next = list.First
	list.First.Pre = list.Last

	list.Size++
}

//yes! 这个Insert暂时是对的，接下来再考虑一下边界情况即可
func (list *List) Insert(x ElemType, pos int) {
	s := new(Node)
	s.Data = x
	p := list.First

	//找位置
	for i := 0; i < pos; i++ {
		p = p.Next
	}

	/*思考过程
	先定位原来这个位置上的节点p
	现在要在此位置上插进来，就要把p往后面挤一个
	故是s.Next=p 等

	*/
	s.Next = p
	s.Pre = p.Pre

	/*技巧：思考前后节点会发生覆盖的连线时
	先连不是直接命名的节点的那一根
	这里直接有命名的是p 就先连p.Pre的Next
	这样能保证不会因为迭代而出错
	*/
	p.Pre.Next = s
	p.Pre = s

	list.Size++
}

//头插法
/*这一段看起来将s放在第二个位置不是First 有问题，
但实际上因为在本文中First不占size，没有数据，输出的时候也是从第二个开始输出的
*/
func (list *List) PushFront(x ElemType) {
	s := new(Node)
	s.Data = x
	s.Next = list.First.Next
	list.First.Next.Pre = s

	list.First.Next = s
	s.Pre = list.First
	//这里重要，若是初始的只有一节，填掉Last
	if list.Size == 0 {
		list.Last = s
	}
	list.Size++
}

//尾删法
func (list *List) PopBack() bool {
	if list.IsEmpty() {
		return false
	}
	s := list.Last.Pre //找到最后一个节点的前驱
	s.Next = list.First
	list.First.Pre = s

	list.Last = s
	list.Size--
	return true
}

/*删结尾
list.Last.Pre.Next=list.First
list.First.Pre=list.Last.Pre
list.Last=list.Last.Pre
文中更简洁的用了s代表Last的前一个节点
*/

//头删法
func (list *List) PopFront() bool {
	if list.IsEmpty() {
		return false
	}
	s := list.First.Next //找到第一个节点,注意第一个不是First
	list.First.Next = s.Next
	s.Next.Pre = list.First

	if list.Size == 1 {
		list.Last = list.First //这里太帅了 非常简洁 只有first的时候就是0size
	}
	list.Size--
	return true
}

//查找指定元素
func (list *List) Find(x ElemType) *Node {
	s := list.First.Next
	for s != list.First { //可以看出first不是实际节点的好处，可以使循环有一个终点条件
		if x == s.Data {
			return s
		} else {
			s = s.Next
		}
	}
	return nil
}

//按值删除结点
func (list *List) DeleteVal(x ElemType) bool {
	s := list.Find(x)
	if s != nil {
		s.Pre.Next = s.Next
		s.Next.Pre = s.Pre
		list.Size--
		//如果删除的是最后一个结点
		if s == list.Last {
			list.Last = s.Pre
		}
		return true
	}
	return false
}

//把值为x的元素的值修改为y
func (list *List) Modify(x, y ElemType) bool {
	s := list.Find(x)
	if s != nil {
		s.Data = y
		return true
	}
	return false
}

//判断链表是否为空
func (list *List) IsEmpty() bool {
	return list.Size == 0
}

//反转链表
//保留第一个结点，将剩余的结点游离出来，然后依次头插到保留的结点中
//?????没看懂
func (list *List) Reverse() {
	if list.Size > 1 {
		s := list.First.Next
		p := s.Next
		s.Next = list.First //第一个结点逆置后成为最后一个结点
		list.Last = s

		for p != list.First {
			s = p
			p = p.Next

			s.Next = list.First.Next
			list.First.Next.Pre = s

			s.Pre = list.First
			list.First.Next = s
		}
	}
}

//打印链表
func (list *List) Print() error {
	if list.IsEmpty() {
		return errors.New("this is an empty list")
	}
	s := list.First.Next
	for s != list.First {
		fmt.Printf("%v  ", s.Data)
		s = s.Next
	}
	fmt.Println()
	return nil
}
