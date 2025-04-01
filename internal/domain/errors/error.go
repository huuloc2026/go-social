package errors

import "errors"

// General Errors
var (
	ErrUnauthorized = errors.New("unauthorized action")
	ErrForbidden    = errors.New("forbidden action")
	ErrBadRequest   = errors.New("invalid request body")
	ErrNotFound     = errors.New("resource not found")
	ErrInternal     = errors.New("internal server error")
	ErrDatabase     = errors.New("database operation failed")
)

// User/Auth Errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailAlreadyExists = errors.New("email already in use")
	ErrUsernameTaken      = errors.New("username already taken")
	ErrTokenInvalid       = errors.New("invalid or expired token")
	ErrPasswordTooWeak    = errors.New("password is too weak")
	ErrUserNotActivated   = errors.New("user account is not activated")
	ErrUserBanned         = errors.New("user account is banned")
	ErrSessionExpired     = errors.New("session has expired")
)

// Post Errors
var (
	ErrPostNotFound         = errors.New("post not found")
	ErrInvalidPostData      = errors.New("invalid post data")
	ErrPostTooLong          = errors.New("post content exceeds character limit")
	ErrImageUploadFail      = errors.New("failed to upload image")
	ErrPostEditNotAllowed   = errors.New("editing post not allowed")
	ErrPostDeleteNotAllowed = errors.New("deleting post not allowed")
)

// Comment Errors
var (
	ErrCommentNotFound         = errors.New("comment not found")
	ErrInvalidCommentData      = errors.New("invalid comment data")
	ErrCommentTooLong          = errors.New("comment content exceeds character limit")
	ErrCommentEditNotAllowed   = errors.New("editing comment not allowed")
	ErrCommentDeleteNotAllowed = errors.New("deleting comment not allowed")
)

// Like/Reaction Errors
var (
	ErrAlreadyLiked   = errors.New("post already liked")
	ErrNotLiked       = errors.New("post not liked yet")
	ErrReactionFailed = errors.New("failed to add reaction")
)
