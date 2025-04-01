import type { RequestConfig } from "./http-client"

/**
 * Authentication interceptor that adds the JWT token to requests
 * Follows Open/Closed Principle by extending functionality without modifying the HttpClient
 */
export class AuthInterceptor {
  private getToken: () => string | null

  constructor(getToken: () => string | null) {
    this.getToken = getToken
  }

  /**
   * Intercept requests to add authentication headers
   */
  intercept(config: RequestConfig): RequestConfig {
    const token = this.getToken()

    if (token) {
      return {
        ...config,
        headers: {
          ...config.headers,
          Authorization: `Bearer ${token}`,
        },
      }
    }

    return config
  }
}

