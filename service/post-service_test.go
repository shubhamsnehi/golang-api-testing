package service

import (
	"testing"

	"github.com/shubhamsnehi/golang-api-testing/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.User) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}

func (mock *MockRepository) FindByID(id string) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}
func (mock *MockRepository) Delete(post *entity.User) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) FindAll() ([]entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.User), args.Error(1)
}
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier string = "1"
	post := entity.User{
		ID:   identifier,
		Name: "A",
	}
	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.User{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion: behavioral
	mockRepo.AssertExpectations(t)

	//Data Assetion
	assert.Equal(t, "A", result[0].Name)
	assert.Equal(t, "1", identifier)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.User{
		Name: "A",
	}

	//Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create1(&post)

	//Mock Assertion: behavioral
	mockRepo.AssertExpectations(t)

	//Data Assetion
	assert.Equal(t, "A", result.Name)
	assert.NotNil(t, result.ID)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}
func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.User{
		ID:   "1",
		Name: "",
	}
	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}
