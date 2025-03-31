// package usecases

// import "github.com/huuloc2026/go-social/internal/domain"

// type FriendshipUsecase interface {
// 	SendRequest(userID, friendID uint) (*domain.Friendship, error)
// 	AcceptRequest(userID, friendshipID uint) (*domain.Friendship, error)
// 	RejectRequest(userID, friendshipID uint) error
// 	GetFriends(userID uint) ([]*domain.User, error)
// 	GetFriendRequests(userID uint) ([]*domain.Friendship, error)
// 	Unfriend(userID, friendID uint) error
// }

// type friendshipUsecase struct {
// 	friendshipRepo domain.FriendshipRepository
// 	userRepo       domain.UserRepository
// 	notification   domain.NotificationService
// }

// func NewFriendshipUsecase(
// 	friendshipRepo domain.FriendshipRepository,
// 	userRepo domain.UserRepository,
// 	notification domain.NotificationService,
// ) FriendshipUsecase {
// 	return &friendshipUsecase{
// 		friendshipRepo: friendshipRepo,
// 		userRepo:       userRepo,
// 		notification:   notification,
// 	}
// }

// func (uc *friendshipUsecase) SendRequest(userID, friendID uint) (*domain.Friendship, error) {
// 	if userID == friendID {
// 		return nil, domain.ErrInvalidRequest
// 	}

// 	// Check if friendship already exists
// 	existing, err := uc.friendshipRepo.FindByUsers(userID, friendID)
// 	if err == nil && existing != nil {
// 		return nil, domain.ErrFriendshipExists
// 	}

// 	friendship := &domain.Friendship{
// 		UserID:   userID,
// 		FriendID: friendID,
// 		Status:   domain.FriendshipPending,
// 	}

// 	if err := uc.friendshipRepo.Create(friendship); err != nil {
// 		return nil, err
// 	}

// 	// Send notification
// 	_ = uc.notification.Send(friendID, userID, "friendship_request", map[string]interface{}{
// 		"friendship_id": friendship.ID,
// 	})

// 	return friendship, nil
// }

// // Implement other methods (AcceptRequest, RejectRequest, etc.)
