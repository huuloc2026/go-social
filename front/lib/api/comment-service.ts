import type { HttpClient } from "./http-client"
import type { ICommentService, Comment, CommentsResponse } from "./interfaces"

/**
 * Concrete implementation of the ICommentService interface
 */
export class CommentService implements ICommentService {
  private client: HttpClient

  constructor(client: HttpClient) {
    this.client = client
  }

  /**
   * Get comments for a post
   */
  async getComments(postId: string): Promise<CommentsResponse> {
    return this.client.get<CommentsResponse>(`/posts/${postId}/comments`)
  }

  /**
   * Create a new comment on a post
   */
  async createComment(postId: string, content: string): Promise<Comment> {
    const response = await this.client.post<{ comment: Comment }>(`/posts/${postId}/comments`, { content })
    return response.comment
  }
}

