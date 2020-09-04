package main

import (
	"fmt"

	"learn.com/gopractise/grammar/oop/library"
)

func main() {
	var m *library.MusicManager = library.NewMusicManager()
	m.Add(&library.MusicEntry{Name: "悟空", Type: "mp3"})
	entry, _ := m.Get(0)
	fmt.Println("1. 音乐库测试:")
	fmt.Println("\t歌曲名称:", entry)
	entry.Name = "八戒"
	fmt.Println("\t通过entry.Name修改歌曲名称为:", entry)
	index, entry1 := m.Find("悟空")
	if entry1 != nil {
		fmt.Println("\t查找名为'八戒'的歌曲:序号index=", index, " , 名称:", entry1.Name)
	}
	fmt.Println("2. 播放音乐测试:")
	library.Play("老神仙", "MP3")
}
