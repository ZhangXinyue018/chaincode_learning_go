package course

type Course struct {
	Name        string `json:name`
	Description string `json:description`
	SchoolName  string `json:school_name`
}
