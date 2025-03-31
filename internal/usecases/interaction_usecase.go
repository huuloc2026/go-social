// package usecases

// import "github.com/huuloc2026/go-social/internal/domain"

// type InteractionUsecase interface {
// 	AddComment(userID, postID uint, content string) (*domain.Comment, error)
// 	DeleteComment(userID, commentID uint) error
// 	AddLike(userID, postID uint) (*domain.Like, error)
// 	RemoveLike(userID, postID uint) error
// 	GetComments(postID uint, limit, offset int) ([]*domain.Comment, error)
// 	GetLikes(postID uint, limit, offset int) ([]*domain.Like, error)
// }

// type interactionUsecase struct {
// 	commentRepo  domain.CommentRepository
// 	likeRepo     domain.LikeRepository
// 	postRepo     domain.PostRepository
// 	notification domain.NotificationService
// }

// // Implement interaction methods
