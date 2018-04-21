package models

// ListItem stores the items on the template
type ListItem struct {
	Title string
	URL   string
}

// IndexPageData is defined to pass data to the indexpage
type IndexPageData struct {
	PageTitle string
	List      []ListItem
}
