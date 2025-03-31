package usecases

// import (
// 	"encoding/json"
// 	"time"

// 	"github.com/huuloc2026/go-social/internal/domain"
// )

// type FeedUsecase interface {
// 	GetUserFeed(userID uint, limit, offset int) ([]*domain.Post, error)
// }

// type feedUsecase struct {
// 	postRepo       domain.PostRepository
// 	friendshipRepo domain.FriendshipRepository
// 	cache          cache.Cache
// }

// func NewFeedUsecase(
// 	postRepo domain.PostRepository,
// 	friendshipRepo domain.FriendshipRepository,
// 	cache cache.Cache,
// ) FeedUsecase {
// 	return &feedUsecase{
// 		postRepo:       postRepo,
// 		friendshipRepo: friendshipRepo,
// 		cache:          cache,
// 	}
// }

// func (uc *feedUsecase) GetUserFeed(userID uint, limit, offset int) ([]*domain.Post, error) {
// 	cacheKey := cache.GenerateFeedKey(userID, limit, offset)

// 	// Try cache first
// 	if cached, err := uc.cache.Get(cacheKey); err == nil {
// 		var posts []*domain.Post
// 		if err := json.Unmarshal([]byte(cached), &posts); err == nil {
// 			return posts, nil
// 		}
// 	}

// 	// Get friends list
// 	friends, err := uc.friendshipRepo.GetFriends(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	friendIDs := make([]uint, len(friends))
// 	for i, f := range friends {
// 		friendIDs[i] = f.ID
// 	}
// 	friendIDs = append(friendIDs, userID) // Include user's own posts

// 	// Get posts from database
// 	posts, err := uc.postRepo.GetFeed(friendIDs, limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Cache the result
// 	if data, err := json.Marshal(posts); err == nil {
// 		_ = uc.cache.Set(cacheKey, string(data), 5*time.Minute)
// 	}

// 	return posts, nil
// }
