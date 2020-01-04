package library

import "errors"

//import "fmt"

/**
音乐库管理类型
***/
type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//音乐库音乐数量
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//按索引查找音乐
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= m.Len() {
		return nil, errors.New("Index out of range.")
	}
	//fmt.Println("\t音乐库管理: ",m.musics[index])
	return &m.musics[index], nil
}

//按名字查找歌曲
func (m *MusicManager) Find(name string) (index int, music *MusicEntry) {
	if m.Len() == 0 {
		return -1, nil
	}
	for index, m := range m.musics {
		if m.Name == name {
			return index, &m
		}
	}
	return -1, nil
}

//添加音乐
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//删除指定索引处的元素
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		return nil
	}
	removeMusic := &m.musics[index]
	if index < len(m.musics)-1 { //不止一个元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { //删除仅有的元素
		m.musics = make([]MusicEntry, 0)
	} else { //删除的是最后一个元素
		m.musics = m.musics[:index-1]
	}
	return removeMusic
}
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	index, _ := m.Find(name)
	if index > -1 {
		return m.Remove(index)
	}
	return nil
}
