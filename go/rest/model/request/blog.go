package request

type BlogListFilter struct {
	Title  string
	Offset int64
	Limit  int64
}

type CreateBlogRequest struct {
	Title   string
	Content string
}
