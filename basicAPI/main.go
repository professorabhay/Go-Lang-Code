package main 


// Modal for course - file 
type Course struct{
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	CoursePrice int `json:"coursePrice"`
	Author *Author `json:"author"` //Author Pointing to the custom type that I defined just below
}

type Author struct{
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB 
var courses []Course

// middleware, helper - file 
func (c *Course) IsEmpty() bool{
	return c.CourseId == "" && c.CourseName == "" 
}

func main(){

}