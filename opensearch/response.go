package opensearch

type SearchResponse struct {
	Status    string   `json:"status"`
	RequestId string   `json:"request_id"`
	Res       *Result  `json:"result"`
	Errors    []*Error `json:"errors"`
}

type Result struct {
	SearchTime float64 `json:"searchtime"`
	Total      int     `json:"total"`
	Num        int     `json:"num"`
	ViewTotal  int     `json:"viewtotal"`
	Items      []*Item `json:"items"`
}

type Item struct {
	Fields         map[string]string `json:"fields"`
	SortExprValues []string          `json:"sortExprValues"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
