package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"pujie.org/practise/bigdata/zookeep/util"
)

var nodes map[string]int = make(map[string]int, 2)

func main() {
	zkList := []string{util.Zk_host}
	conn := util.GetConnect(zkList)
	defer conn.Close()

	watch_root(conn, util.Zk_work_path)

}

func freshWatch(conn *zk.Conn, root string) []string {
	children, _, err := conn.Children(root)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("==children of path: %s start..... \n", root)
	for _, path := range children {
		fmt.Printf("\tchildren: %s \n", path)
		/*		if _,ok := nodes[path];!ok{
				nodes[path] = 1
				//go test_Wath_Znode(conn, util.Zk_work_path+"/"+path)
			}*/
	}
	fmt.Println("==children of path end \n")
	//fmt.Printf("zk_work_path=%s: %v \n",util.Zk_work_path, children)
	return children
}

func watch_root(conn *zk.Conn, root string) {
	freshWatch(conn, root) //初始化
	for {
		children, _, childCh, err := conn.ChildrenW(root)
		if err != nil {
			fmt.Println("watch children error, ", err)
			return
		}
		select {
		case Event := <-childCh:
			if Event.Type == zk.EventNodeChildrenChanged {
				fmt.Printf("root %s,event-type:%v  \n", root, Event.Type)
				for _, path := range children {
					fmt.Printf("\tChildrenW: %s \n", path)
				}
				freshWatch(conn, root)
			} else {
				fmt.Printf("\t %v : root %s : Event  \n", Event.Type, root)
			}
		}
	}
}

/**
 * 获取所有节点
 */
func test_Wath_Znode(conn *zk.Conn, path string) {

	for {
		children, _, childCh, err := conn.ChildrenW(path)
		if err != nil {
			fmt.Println("watch children error, ", err)
			return
		}
		//fmt.Println("watch children result, ", children, state)
		select {
		case Event := <-childCh:
			if Event.Type == zk.EventNodeDeleted {
				delete(nodes, path) //删除本地缓存
				fmt.Printf("\t %v : Znode %s,%s \n", Event.Type, path, children)
				return
			} else {
				fmt.Printf("\t %v : Znode %s : Event \n", Event.Type, path)
			}
		}
	}

}
