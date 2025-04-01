import { HttpClient } from "./http-client"
import { AuthInterceptor } from "./auth-interceptor"
import { AuthService } from "./auth-service"
import { PostService } from "./post-service"
import { CommentService } from "./comment-service"
import type { IAuthService, IPostService, ICommentService } from "./interfaces"

/**
 * API Factory - Creates and configures API services
 * Follows Dependency Inversion Principle by providing abstractions
 */
export class ApiFactory {
  private client: HttpClient
  private authService: AuthService
  private postService: PostService
  private commentService: CommentService

  constructor(baseUrl: string) {
    // Create the HTTP client
    this.client = new HttpClient(baseUrl)

    // Create the auth service first
    this.authService = new AuthService(this.client)

    // Create the auth interceptor using the auth service's getToken method
    const authInterceptor = new AuthInterceptor(this.authService.getToken.bind(this.authService))

    // Add the auth interceptor to the client
    this.client.addRequestInterceptor(authInterceptor.intercept.bind(authInterceptor))

    // Create other services
    this.postService = new PostService(this.client)
    this.commentService = new CommentService(this.client)

    // Add error handling interceptor
    this.client.addErrorInterceptor(this.handleApiError.bind(this))
  }

  /**
   * Get the auth service
   */
  getAuthService(): IAuthService {
    return this.authService
  }

  /**
   * Get the post service
   */
  getPostService(): IPostService {
    return this.postService
  }

  /**
   * Get the comment service
   */
  getCommentService(): ICommentService {
    return this.commentService
  }

  /**
   * Handle API errors
   */
  private handleApiError(error: any): any {
    // Handle 401 Unauthorized errors
    if (error.status === 401) {
      // Clear the token if unauthorized
      this.authService.logout().catch(console.error)
    }

    // Enhance the error with more information
    return {
      ...error,
      message: error.data?.message || error.statusText || "An unknown error occurred",
      isApiError: true,
    }
  }
}

