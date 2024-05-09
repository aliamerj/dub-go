package tags

type color string

const (
	Red    color = "red"
	Yellow color = "yellow"
	Green  color = "green"
	Blue   color = "blue"
	Purple color = "purple"
	Pink   color = "pink"
	Brown  color = "brown"
)

type Options struct {
	Tag   string `json:"tag"`
	Color color  `json:"color"`
}
type resTag struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
