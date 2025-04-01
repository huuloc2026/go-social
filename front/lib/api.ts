export async function fetchWithAuth(url: string, options: RequestInit = {}) {
  // Get the token from the cookie
  const token = document.cookie
    .split("; ")
    .find((row) => row.startsWith("auth_token="))
    ?.split("=")[1]

  // Set up headers with authorization if token exists
  const headers = {
    ...(options.headers || {}),
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  }

  // Make the request
  return fetch(url, {
    ...options,
    headers,
  })
}

