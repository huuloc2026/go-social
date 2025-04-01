export interface IAuthService {
  login(email: string, password: string): Promise<AuthResponse>
  register(name: string, email: string, password: string): Promise<AuthResponse>
  logout(): Promise<void>
  getCurrentUser(): Promise<User>
}

/**
 * Interface for post-related API calls
 */
export interface IPostService {
  getPosts(page?: number): Promise<PostsResponse>
  createPost(content: string): Promise<Post>
  likePost(postId: string): Promise<void>
  unlikePost(postId: string): Promise<void>
}

/**
 * Interface for comment-related API calls
 */
export interface ICommentService {
  getComments(postId: string): Promise<CommentsResponse>
  createComment(postId: string, content: string): Promise<Comment>
}

/**
 * Response types
 */
export interface AuthResponse {
  user: User
  token?: string
}

export interface PostsResponse {
  posts: Post[]
  hasMore: boolean
}

export interface CommentsResponse {
  comments: Comment[]
}

/**
 * Data types
 */
export interface User {
  id: string
  name: string
  email: string
}

export interface Post {
  id: string
  content: string
  createdAt: string
  user: {
    id: string
    name: string
  }
  likesCount: number
  commentsCount: number
  isLiked: boolean
}

export interface Comment {
  id: string
  content: string
  createdAt: string
  user: {
    id: string
    name: string
  }
}

