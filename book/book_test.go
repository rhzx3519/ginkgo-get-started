package book_test

import (
	"ginkgo.guide.com/v2/book"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var foxInSocks, lesMis *book.Book

	BeforeEach(func() {
		lesMis = &book.Book{
			Title:  "Les Miserable",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		foxInSocks = &book.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing books", func() {
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(book.CategoryNovel))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(book.CategoryShortStory))
			})
		})
	})
})
