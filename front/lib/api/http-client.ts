export class HttpClient {
  private baseUrl: string
  private defaultHeaders: Record<string, string>
  private requestInterceptors: ((config: RequestConfig) => RequestConfig)[] = []
  private responseInterceptors: ((response: any) => any)[] = []
  private errorInterceptors: ((error: any) => any)[] = []

  constructor(baseUrl: string, defaultHeaders: Record<string, string> = {}) {
    this.baseUrl = baseUrl
    this.defaultHeaders = {
      "Content-Type": "application/json",
      ...defaultHeaders,
    }
  }

  /**
   * Add a request interceptor
   */
  addRequestInterceptor(interceptor: (config: RequestConfig) => RequestConfig): void {
    this.requestInterceptors.push(interceptor)
  }

  /**
   * Add a response interceptor
   */
  addResponseInterceptor(interceptor: (response: any) => any): void {
    this.responseInterceptors.push(interceptor)
  }

  /**
   * Add an error interceptor
   */
  addErrorInterceptor(interceptor: (error: any) => any): void {
    this.errorInterceptors.push(interceptor)
  }

  /**
   * Process request config through all request interceptors
   */
  private processRequestInterceptors(config: RequestConfig): RequestConfig {
    return this.requestInterceptors.reduce((acc, interceptor) => interceptor(acc), config)
  }

  /**
   * Process response through all response interceptors
   */
  private processResponseInterceptors(response: any): any {
    return this.responseInterceptors.reduce((acc, interceptor) => interceptor(acc), response)
  }

  /**
   * Process error through all error interceptors
   */
  private processErrorInterceptors(error: any): any {
    return this.errorInterceptors.reduce((acc, interceptor) => interceptor(acc), error)
  }

  /**
   * Make an HTTP request
   */
  async request<T>(config: RequestConfig): Promise<T> {
    try {
      // Apply request interceptors
      const processedConfig = this.processRequestInterceptors(config)

      // Build the full URL
      const url = new URL(processedConfig.path, processedConfig.path.startsWith("http") ? undefined : this.baseUrl)

      // Add query parameters if they exist
      if (processedConfig.params) {
        Object.entries(processedConfig.params).forEach(([key, value]) => {
          if (value !== undefined && value !== null) {
            url.searchParams.append(key, String(value))
          }
        })
      }

      // Merge headers
      const headers = {
        ...this.defaultHeaders,
        ...processedConfig.headers,
      }

      // Prepare the fetch options
      const fetchOptions: RequestInit = {
        method: processedConfig.method,
        headers,
        credentials: processedConfig.withCredentials ? "include" : "same-origin",
      }

      // Add body for non-GET requests
      if (processedConfig.method !== "GET" && processedConfig.data) {
        fetchOptions.body = JSON.stringify(processedConfig.data)
      }

      // Make the request
      const response = await fetch(url.toString(), fetchOptions)

      // Handle different response types
      let data: any
      const contentType = response.headers.get("Content-Type") || ""

      if (contentType.includes("application/json")) {
        data = await response.json()
      } else if (contentType.includes("text/")) {
        data = await response.text()
      } else {
        data = await response.blob()
      }

      // Check if the response is successful
      if (!response.ok) {
        throw {
          status: response.status,
          statusText: response.statusText,
          data,
        }
      }

      // Apply response interceptors
      return this.processResponseInterceptors(data)
    } catch (error) {
      // Apply error interceptors and rethrow
      const processedError = this.processErrorInterceptors(error)
      throw processedError
    }
  }

  /**
   * HTTP GET request
   */
  async get<T>(path: string, params?: Record<string, any>, headers?: Record<string, string>): Promise<T> {
    return this.request<T>({
      method: "GET",
      path,
      params,
      headers,
    })
  }

  /**
   * HTTP POST request
   */
  async post<T>(path: string, data?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>({
      method: "POST",
      path,
      data,
      headers,
    })
  }

  /**
   * HTTP PUT request
   */
  async put<T>(path: string, data?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>({
      method: "PUT",
      path,
      data,
      headers,
    })
  }

  /**
   * HTTP PATCH request
   */
  async patch<T>(path: string, data?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>({
      method: "PATCH",
      path,
      data,
      headers,
    })
  }

  /**
   * HTTP DELETE request
   */
  async delete<T>(path: string, data?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>({
      method: "DELETE",
      path,
      data,
      headers,
    })
  }
}

/**
 * Request configuration interface
 */
export interface RequestConfig {
  method: "GET" | "POST" | "PUT" | "PATCH" | "DELETE"
  path: string
  params?: Record<string, any>
  data?: any
  headers?: Record<string, string>
  withCredentials?: boolean
}

