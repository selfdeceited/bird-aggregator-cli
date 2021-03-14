package flickr

type Photo struct {
	ID       string `json:"id"`
	Owner    string `json:"owner"`
	Secret   string `json:"secret"`
	Server   string `json:"server"`
	Farm     int    `json:"farm"`
	Title    string `json:"title"`
	Ispublic int    `json:"ispublic"`
	Isfriend int    `json:"isfriend"`
	Isfamily int    `json:"isfamily"`
}

type PhotosResponse struct {
	Photos struct {
		Page    int     `json:"page"`
		Pages   int     `json:"pages"`
		Perpage int     `json:"perpage"`
		Total   string  `json:"total"`
		Photo   []Photo `json:"photo"`
	} `json:"photos"`
	Stat string `json:"stat"`
}

type CallOptions struct {
	APIKey  string
	UserID  string
	PerPage int
}

type CallResult struct {
	PagesCount int
	PhotoNames []string
}
