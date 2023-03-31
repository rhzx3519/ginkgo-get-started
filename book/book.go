package book

type Category string

const (
	CategoryNovel      Category = "NOVEL"
	CategoryShortStory Category = "SHORT STORY"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) Category() Category {
	if b.Pages >= 300 {
		return CategoryNovel
	}
	return CategoryShortStory
}
