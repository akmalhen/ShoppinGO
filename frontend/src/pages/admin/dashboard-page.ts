import '../../styles/dashboard.css';
import type { Product, DashboardStats } from '../../types'; 
import { apiFetch } from '../../utils/api';

if (!localStorage.getItem('authToken')) {
    window.location.href = '/login.html';
}

const statsContainer = document.querySelector<HTMLDivElement>('#stats-container')!;
const latestProductsBody = document.querySelector<HTMLTableSectionElement>('#latest-products-body')!;
const adminNameDisplay = document.querySelector<HTMLSpanElement>('#admin-name-display')!;
const logoutButton = document.querySelector<HTMLAnchorElement>('#logout-button')!;

function displayUserInfo() {
    const token = localStorage.getItem('authToken');
    if (token) {
        try {
            const payload = JSON.parse(atob(token.split('.')[1]));
            adminNameDisplay.textContent = payload.name || payload.email || 'Admin';
        } catch (e) {
            adminNameDisplay.textContent = 'Admin';
        }
    }
}

async function loadDashboardData() {
    try {
        const data = await apiFetch<DashboardStats>('/dashboard');
        
        statsContainer.innerHTML = `
            <div class="stat-card"><h3>Jumlah User</h3><p class="stat-value">${data.total_users} <span>User</span></p></div>
            <div class="stat-card"><h3>User Aktif</h3><p class="stat-value">${data.active_users} <span>User</span></p></div>
            <div class="stat-card"><h3>Total Produk</h3><p class="stat-value">${data.total_products} <span>Produk</span></p></div>
            <div class="stat-card"><h3>Produk Tersedia</h3><p class="stat-value">${data.available_products} <span>Produk</span></p></div>
        `;

        latestProductsBody.innerHTML = data.latest_products.map((product: Product) => `
            <tr>
                <td>
                    <div class="product-cell">
                        <img src="${product.image_url || 'https://via.placeholder.com/40'}" alt="${product.name}">
                        <span>${product.name}</span>
                    </div>
                </td>
                <td>${new Date(product.CreatedAt).toLocaleDateString('id-ID', { day: '2-digit', month: 'long', year: 'numeric' })}</td>
                <td>${new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(product.price)}</td>
            </tr>
        `).join('');
    } catch (error) {
        console.error("Gagal memuat data dashboard:", error);
    }
}

logoutButton.addEventListener('click', (e) => {
    e.preventDefault();
    localStorage.removeItem('authToken');
    window.location.href = '/login.html';
});

displayUserInfo();
loadDashboardData();