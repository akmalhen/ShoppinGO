export async function apiFetch<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('authToken');
  const API_BASE_URL = 'http://localhost:8080';

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  };

  if (options.headers) {
      Object.assign(headers, options.headers);
  }

  if (token && (endpoint.startsWith('/admin') || endpoint.startsWith('/users'))) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  
  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers: headers,
  });

  if (response.status === 401) {
    localStorage.removeItem('authToken');
    window.location.href = '/login.html';
    throw new Error('Unauthorized');
  }
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ error: `Request failed with status: ${response.status}` }));
    throw new Error(errorData.error || 'API request failed');
  }
  
  if (response.status === 204) {
      return null as T;
  }
  
  return response.json() as Promise<T>;
}