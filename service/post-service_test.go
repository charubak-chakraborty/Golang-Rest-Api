package service

import (
	"testing"

	"../entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectation
	post := entity.Post{ID: 1, Title: "titletest", Text: "texttest"}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	//create test
	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)

	//data assertion
	assert.Equal(t, post.ID, result[0].ID)
	assert.Equal(t, post.Title, result[0].Title)
	assert.Equal(t, post.Text, result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectation
	post := entity.Post{ID: 1, Title: "titletest", Text: "texttest"}
	mockRepo.On("Save").Return(&post, nil)

	//create test
	testService := NewPostService(mockRepo)
	result, _ := testService.Create(&post)

	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)

	//data assertion
	assert.Equal(t, post.ID, result.ID)
	assert.Equal(t, post.Title, result.Title)
	assert.Equal(t, post.Text, result.Text)
}
func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "test"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post Title is empty")
}
