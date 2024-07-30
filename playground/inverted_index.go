package playground

type Doc struct {
	Id       int
	KeyWords []string
}

func BuildInvertIndex(docs []*Doc) map[string][]int {
	index := make(map[string][]int, 100)
	for _, doc := range docs {
		for _, keyword := range doc.KeyWords {
			// go 中 slice 的零值 nil slice 长度/容量为0, 可以进行append操作
			index[keyword] = append(index[keyword], doc.Id)
		}
	}
	return index
}
