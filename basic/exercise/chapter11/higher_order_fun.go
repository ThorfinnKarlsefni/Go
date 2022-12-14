package chapter11

import "fmt"

type Album struct {
	name    string
	author  string
	pubDate int
}

type Albums []Album

func HigherOrderFunc() {
	album1 := Album{"美丽谎言", "陈奕迅", 1997}
	album2 := Album{"半岛铁盒", "周杰伦", 2000}
	album3 := Album{"六月的雨", "胡歌", 2008}
	album4 := Album{"忘记时间", "胡歌", 2010}
	album5 := Album{"Don't be afraid", "Moxizishi", 2020}
	allAlbums := Albums{album1, album2, album3, album4,album5}
	allNewAlbums := allAlbums.FindAll(func(album Album) bool {
		return album.pubDate > 1999
	})

	fmt.Println(allNewAlbums)
	appender,newAllAlbums := allAlbums.MakeAppenderSorted()
	allAlbums.Process(appender)
	//fmt.Println(allAlbums)
	fmt.Println(newAllAlbums)
	fmt.Println("count 胡歌",len(newAllAlbums["胡歌"]))
}

func (al Albums) Process(f func(a Album)) {
	for _, a := range al {
		f(a)
	}
}

func (al Albums) FindAll(f func(album Album) bool) Albums {
	albums := make([]Album, 0)

	al.Process(func(a Album) {
		if f(a) {
			albums = append(albums, a)
		}
	})

	return albums
}


func (al Albums) MakeAppenderSorted() (func (album Album), map[string][]Album){
	newAlbums := make(map[string][]Album)
	newAlbums["Default"] = make([]Album,0)

	appender := func(a Album){
		if len(a.author) > 0 {
			newAlbums[a.author] = append(newAlbums[a.author],a)
		}else{
			newAlbums["Default"] = append(newAlbums["Default"],a)
		}
	}
	return appender, newAlbums
}


