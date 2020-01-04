package main

import (
	"fmt"
	"sort"
)

//售卖方式
type SaleMode struct {
	Type      uint32   //类型. 1:api; 2:直投;
	Weight    int      //权重大的, 能获得更大的分发机会.
	Priority  int      //处理优先级.(先通过权限选择)
	SalesList []string //售卖目标列表: 如果类型为api则(渠道编码:权重);如果类型为直投,则(CID:权重)
}

//为*Person添加String()方法，便于输出
func (p *SaleMode) String() string {
	return fmt.Sprintf("( %d,%d,%d)", p.Type, p.Weight, p.Priority)
}

type SaleModeList []*SaleMode

func (list SaleModeList) Len() int {
	return len(list)
}

//排序规则：首先按年龄排序（由小到大），年龄相同时按姓名进行排序（按字符串的自然顺序）
func (list SaleModeList) Less(i, j int) bool {
	if list[i].Priority <= list[j].Priority {
		return true
	} else {
		return false
	}
}

func (list SaleModeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

type ApiMode SaleMode

func (p *ApiMode) String() string {
	return fmt.Sprintf("apiMode.String()")
}

func Start_Sort_test() {
	fmt.Println("------")
	p1 := &SaleMode{1, 24, 5, []string{"(001,5)"}}
	p2 := &SaleMode{1, 24, 1, []string{"(001,5)"}}
	p3 := &SaleMode{1, 24, 9, []string{"(001,5)"}}
	p4 := &SaleMode{1, 24, 2, []string{"(001,5)"}}
	p5 := &SaleMode{1, 24, 0, []string{"(001,5)"}}
	p6 := &SaleMode{1, 24, 6, []string{"(001,5)"}}

	p7 := ApiMode{1, 8, 2, []string{"(001,5)"}}

	pList := SaleModeList([]*SaleMode{p1, p2, p3, p4, p5, p6})

	//p7 = SaleMode(p7)
	fmt.Println(p7.String())
	//pList = append(pList, p7)
	//sort.Sort(pList) //顺序
	sort.Sort(sort.Reverse(pList)) //逆序

	fmt.Println(pList)
}
