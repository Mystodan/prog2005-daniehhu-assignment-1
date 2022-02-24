package handler

type countryName struct {
	Name string `json:"common"`
}
type openStreetMap struct {
	Map string `json:"openStreetMaps"`
}

type getUnii struct {
	Country  string   `json:"country"` // Suppress field in JSON output if it is empty
	Alpha_2  string   `json:"alpha_two_code"`
	Name     string   `json:"name"`
	Webpages []string `json:"web_pages"`
}
type getCountry struct {
	Name      countryName       `json:"name"`
	Languages map[string]string `json:"languages"`
	Maps      openStreetMap     `json:"maps"`
	Borders   []string          `json:"borders"`
	Isocode   string            `json:"cca2"`
}

type Universities struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
}

type Diagnostics struct {
	UniversityAPI string `json:"universitiesapi"`
	CountriesAPI  string `json:"countriesapi"`
	Version       string `json:"version"`
	Uptime        string `json:"uptime"`
}
