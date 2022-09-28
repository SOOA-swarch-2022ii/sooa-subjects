package models

type Subject struct {
	Id                string   `bson:"id,omitempty" json:"id,omitempty"`
	Name              string   `bson:"name,omitempty" json:"name,omitempty"`
	Vigency           bool     `bson:"vigency,omitempty" json:"vigency,omitempty"`
	Level             string   `bson:"level,omitempty" json:"level,omitempty"`
	Credits           int      `bson:"credits,omitempty" json:"credits,omitempty"`
	Campus            string   `bson:"campus,omitempty" json:"campus,omitempty"`
	Faculty           string   `bson:"faculty,omitempty" json:"faculty,omitempty"`
	Department        string   `bson:"department,omitempty" json:"subfaculty,omitempty"`
	BasicAcademicUnit string   `bson:"basic_academic_unit,omitempty" json:"basic_academic_unit,omitempty"`
	Academic_level    string   `bson:"academic_level,omitempty" json:"academic_level,omitempty"`
	Content           []string `bson:"content,omitempty" json:"content,omitempty"`
}

type Course struct {
	Id       string `bson:"_id,omitempty" json:"_id,omitempty"`
	Subject  string `bson:"subject,omitempty" json:"subject,omitempty"`
	Schedule []struct {
		Day      string `bson:"day,omitempty" json:"day,omitempty"`
		StartH   string `bson:"start_h,omitempty" json:"start_h,omitempty"`
		EndH     string `bson:"end_h,omitempty" json:"end_h,omitempty"`
		Location string `bson:"location,omitempty" json:"location,omitempty"`
	} `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Academic_semester string   `bson:"academic_semester,omitempty" json:"academic_semester,omitempty"`
	Start_date        string   `bson:"start_date,omitempty" json:"start_date,omitempty"`
	End_date          string   `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Group_number      int32    `bson:"group_number,omitempty" json:"group_number,omitempty"`
	Places            int32    `bson:"places,omitempty" json:"places,omitempty"`
	Professors        []string `bson:"professors,omitempty" json:"professors,omitempty"`
	StudentsRecord    []struct {
		Student string    `bson:"student,omitempty" json:"student,omitempty"`
		Grades  []float64 `bson:"grades,omitempty" json:"grades,omitempty"`
	} `bson:"students_record,omitempty" json:"students_record,omitempty"`
}
