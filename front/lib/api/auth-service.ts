import type { HttpClient } from "./http-client"
import type { IAuthService, AuthResponse, User } from "./interfaces"

/**
 * Concrete implementation of the IAuthService interface
 * Follows Liskov Substitution Principle by properly implementing the interface
 */
export class AuthService implements IAuthService {
  private client: HttpClient
  private tokenKey = "auth_token"

  constructor(client: HttpClient) {
    this.client = client
  }

  /**
   * Get the stored JWT token
   */
  getToken(): string | null {
    if (typeof document !== "undefined") {
      return (
        document.cookie
          .split("; ")
          .find((row) => row.startsWith(`${this.tokenKey}=`))
          ?.split("=")[1] || null
      )
    }
    return null
  }

  /**
   * Set the JWT token in a cookie
   */
  private setToken(token: string): void {
    if (typeof document !== "undefined") {
      const secure = window.location.protocol === "https:"
      document.cookie = `${this.tokenKey}=${token}; path=/; ${secure ? "secure; " : ""}max-age=${60 * 60 * 24 * 7}` // 1 week
    }
  }

  /**
   * Clear the JWT token
   */
  private clearToken(): void {
    if (typeof document !== "undefined") {
      document.cookie = `${this.tokenKey}=; path=/; max-age=0`
    }
  }

  /**
   * Login with email and password
   */
  async login(email: string, password: string): Promise<AuthResponse> {
    const response = await this.client.post<AuthResponse>("/auth/login", { email, password })

    if (response.token) {
      this.setToken(response.token)
    }

    return response
  }

  /**
   * Register a new user
   */
  async register(name: string, email: string, password: string): Promise<AuthResponse> {
    const response = await this.client.post<AuthResponse>("/auth/register", { name, email, password })

    if (response.token) {
      this.setToken(response.token)
    }

    return response
  }

  /**
   * Logout the current user
   */
  async logout(): Promise<void> {
    await this.client.post<void>("/auth/logout")
    this.clearToken()
  }

  /**
   * Get the current user's information
   */
  async getCurrentUser(): Promise<User> {
    const response = await this.client.get<AuthResponse>("/auth/me")
    return response.user
  }
}

