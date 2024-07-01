package service

/*type CourseService interface {

}

type courseProxy interface{
	LearnGo(ctx, req) (resp, err)
}

type CourseServiceImpl struct {
	courseClient courseProxy
}

func NewCourseService(courseClient courseProxy) CourseService{
	return &courseServiceImpl{
		courseClient: courseClient,
	}
}

func (c *CourseServiceImpl) LearnGo(ctx, req) (resp, err) {
	...
	//c := client.newCourseClient()
	//resp, err = c.LearnGo(ctx, req)
	c.CourseClient.LearnGo(ctx, req)
	...
}

func (c *CourseServiceImpl) LearnJava(ctx, req) (resp, err) {
	http.POST()
}

type mockCourseClient struct{}


func (c *mockCourseClient) LearnGo(ctx, req) (resp, err) {
	return map[string]interface{},nil
}


func Test_courseServiceImpl_LearnGo(t *testing.T){
	mockService := NewCourseService(&mockCourseClient{})
}

*/
