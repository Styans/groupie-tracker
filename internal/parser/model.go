package parser

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	ConcertDates ConcertDates
	Relations    Relations
}
type (
	IndexLocations struct {
		Locations []Locations `json:"index"`
	}
	IndexConcert struct {
		ConcertsDates []ConcertDates `json:"index"`
	}
	IndexRelations struct {
		Relations []Relations `json:"index"`
	}
)
type (
	Locations struct {
		Id           int      `json:"id"`
		Locations    []string `json:"locations"`
		ConcertDates ConcertDates
	}
	ConcertDates struct {
		Id            int      `json:"id"`
		ConcertsDates []string `json:"dates"`
	}
	Relations struct {
		Id           int                 `json:"id"`
		Relationsmap map[string][]string `json:"datesLocations"`
	}
)
type Err struct {
	StatusCode int
	StatusText string
}
