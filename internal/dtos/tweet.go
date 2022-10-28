package dtos

type CreateTweetRequest struct {
	Text string
	Url  string
}

type UpdateTweetRequest struct {
	Text string
	Url  string
}
