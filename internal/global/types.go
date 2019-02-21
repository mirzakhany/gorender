package global

// D provide string array
type D map[string]interface{}

// DS provide string to string map
type DS map[string]string

type RenderReq struct {
	HtmlTemplate string `json:"html_template"`
	Data interface{} `json:"data"`
	Options struct{
		Pdf struct{
			PageMargin []int `json:"page_margin"`
		} `json:"pdf"`

	} `json:"options"`
}

type RenderRes struct {
	Files struct{
		Pdf string `json:"pdf"`
		JPG string `json:"jpg"`
		HTML string `json:"html"`
	} `json:"files"`
}