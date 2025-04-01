// User types
export interface User {
  id: string
  name: string
  username: string
  email: string
  avatar?: string
  bio?: string
  createdAt: string
}

// Post types
export interface Post {
  id: string
  content: string
  image?: string
  author: User
  likeCount: number
  commentCount: number
  isLiked: boolean
  comments?: Comment[]
  createdAt: string
  updatedAt: string
}

export interface Comment {
  id: string
  content: string
  author: User
  postId: string
  createdAt: string
}

// Notification types
export interface Notification {
  id: string
  type: "friend_request" | "friend_accept" | "post_like" | "post_comment"
  sender: User
  postId?: string
  read: boolean
  createdAt: string
}

// Friendship types
export type FriendshipStatus = "none" | "pending_sent" | "pending_received" | "friends"

