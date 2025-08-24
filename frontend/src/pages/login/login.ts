import '../../styles/login.css';

const loginForm = document.querySelector<HTMLFormElement>('#login-form')!;
const emailInput = document.querySelector<HTMLInputElement>('#email')!;
const passwordInput = document.querySelector<HTMLInputElement>('#password')!;
const errorMessageElement = document.querySelector<HTMLParagraphElement>('#error-message')!;

loginForm.addEventListener('submit', async (event) => {
  event.preventDefault();
  errorMessageElement.textContent = '';
  const email = emailInput.value;
  const password = passwordInput.value;
  const API_URL = 'http://localhost:8080/admin/login';

  try {
    const response = await fetch(API_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    const data = await response.json();
    if (!response.ok) { throw new Error(data.error || 'Terjadi kesalahan saat login.'); }
    if (data.token) {
      localStorage.setItem('authToken', data.token);
      window.location.href = '/dashboard.html';
    } else {
      throw new Error('Token tidak diterima dari server.');
    }
  } catch (error: any) {
    errorMessageElement.textContent = error.message;
  }
});