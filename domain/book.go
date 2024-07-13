package domain


type Book struct {
	Id int64 `json:"id"`
	Name string    `json:"name"`
	Owner   string    `json:"owner"`
	Url     string    `json:"url"`
	SummarizedUrl string `json:"summarized_url"`
}
