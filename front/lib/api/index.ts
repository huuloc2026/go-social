import { ApiFactory } from "./api-factory"
import type { IAuthService, IPostService, ICommentService } from "./interfaces"

/**
 * API Client - Main entry point for the API
 */
class ApiClient {
  private factory: ApiFactory
  private static instance: ApiClient

  private constructor() {
    const baseUrl = process.env.NEXT_PUBLIC_API_URL || ""
    this.factory = new ApiFactory(baseUrl)
  }

  /**
   * Get the singleton instance
   */
  public static getInstance(): ApiClient {
    if (!ApiClient.instance) {
      ApiClient.instance = new ApiClient()
    }
    return ApiClient.instance
  }

  /**
   * Get the auth service
   */
  public get auth(): IAuthService {
    return this.factory.getAuthService()
  }

  /**
   * Get the post service
   */
  public get posts(): IPostService {
    return this.factory.getPostService()
  }

  /**
   * Get the comment service
   */
  public get comments(): ICommentService {
    return this.factory.getCommentService()
  }
}

// Export a singleton instance
export const api = ApiClient.getInstance()

// Export types
export * from "./interfaces"

