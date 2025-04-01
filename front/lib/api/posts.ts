import type { Post, Comment } from "@/types"
import { getAuthToken } from "@/lib/auth"

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api"

// Helper function to get headers with auth token
const getHeaders = (includeContentType = true) => {
  const headers: Record<string, string> = {
    Authorization: `Bearer ${getAuthToken()}`,
  }

  if (includeContentType) {
    headers["Content-Type"] = "application/json"
  }

  return headers
}

export async function getPosts(filter = "all"): Promise<{ posts: Post[] }> {
  const response = await fetch(`${API_URL}/posts?filter=${filter}`, {
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to fetch posts")
  }

  return await response.json()
}

export async function getUserPosts(userId: string): Promise<{ posts: Post[] }> {
  const response = await fetch(`${API_URL}/users/${userId}/posts`, {
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to fetch user posts")
  }

  return await response.json()
}

export async function createPost(content: string, image: File | null): Promise<Post> {
  const formData = new FormData()
  formData.append("content", content)

  if (image) {
    formData.append("image", image)
  }

  const response = await fetch(`${API_URL}/posts`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${getAuthToken()}`,
    },
    body: formData,
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to create post")
  }

  return (await response.json()).post
}

export async function updatePost(postId: string, content: string): Promise<Post> {
  const response = await fetch(`${API_URL}/posts/${postId}`, {
    method: "PUT",
    headers: getHeaders(),
    body: JSON.stringify({ content }),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to update post")
  }

  return (await response.json()).post
}

export async function deletePost(postId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/posts/${postId}`, {
    method: "DELETE",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to delete post")
  }

  return await response.json()
}

export async function likePost(postId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/posts/${postId}/like`, {
    method: "POST",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to like post")
  }

  return await response.json()
}

export async function unlikePost(postId: string): Promise<{ message: string }> {
  const response = await fetch(`${API_URL}/posts/${postId}/unlike`, {
    method: "POST",
    headers: getHeaders(),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to unlike post")
  }

  return await response.json()
}

export async function createComment(postId: string, content: string): Promise<Comment> {
  const response = await fetch(`${API_URL}/posts/${postId}/comments`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ content }),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "Failed to create comment")
  }

  return (await response.json()).comment
}

