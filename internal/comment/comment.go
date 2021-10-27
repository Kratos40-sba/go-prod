package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - define the comment structr
type Comment struct {
	gorm.Model
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

// GetComment - retrieves comments by their ID from database
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, nil
	}
	return comment, nil
}

// UpdateComment - updates a comment by the newComment
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// DeleteComment - deletes a comment from database by ID
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - retrieves all comments from databases
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
