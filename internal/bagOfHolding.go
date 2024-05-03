package internal

type Bag struct {
	Owner int     `json:"owner"`
	Items []*Item `json:"items"`
}
