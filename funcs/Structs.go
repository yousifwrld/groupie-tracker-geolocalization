package web

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// this struct is for adding the locations for the locations filter
type LocationsForMainData struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type ArtistPageData struct {
	Artists            []Artists
	NumOfMembers       [8]string
	Locations          []string
	SearchLocations    []string
	SearchCreationDate []int
	SearchFirstAlbum   []string
}

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type Dates struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}
type Relations struct {
	Id            int                 `json:"id"`
	LocationDates map[string][]string `json:"datesLocations"`
}

type Cords struct {
	Name string
	Lat  float64
	Lng  float64
}

// struct to hold all the data together
type AllDetails struct {
	Artist    Artists
	Locations Locations
	Dates     Dates
	Relations Relations
	Cords     []Cords
}
