package model

type RequestQuery struct {
	From                 int               `json:"from"`
	Size                 int               `json:"size"`
	SearchKeyword        string            `json:"searchKeyword"`
	SearchFields         []string          `json:"searchFields"`
	AuthorName           []string          `json:"authorName"`
	TagId                []string          `json:"tagId"`
	PostTermTaxonomyId   []string          `json:"postTermTaxonomyId"`
	EntityTermTaxonomyId string            `json:"entityTermTaxonomyId"`
	Sort                 map[string]string `json:"sort"`
}

type Request struct {
	Body   string                 `json:"body"`
	Params map[string]interface{} `json:"params"`
}
