import type { HttpClient } from "./http-client"
import type { IPostService, Post, PostsResponse } from "./interfaces"

/**
 * Concrete implementation of the IPostService interface
 */
export class PostService implements IPostService {
  private client: HttpClient

  constructor(client: HttpClient) {
    this.client = client
  }

  /**
   * Get posts with pagination
   */
  async getPosts(page = 1): Promise<PostsResponse> {
    return this.client.get<PostsResponse>("/posts", { page })
  }

  /**
   * Create a new post
   */
  async createPost(content: string): Promise<Post> {
    const response = await this.client.post<{ post: Post }>("/posts", { content })
    return response.post
  }

  /**
   * Like a post
   */
  async likePost(postId: string): Promise<void> {
    await this.client.post<void>(`/posts/${postId}/like`)
  }

  /**
   * Unlike a post
   */
  async unlikePost(postId: string): Promise<void> {
    await this.client.delete<void>(`/posts/${postId}/like`)
  }
}

