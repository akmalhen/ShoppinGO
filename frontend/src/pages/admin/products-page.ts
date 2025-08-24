import '../../styles/products.css';
import type { Product } from '../../types';
import { apiFetch } from '../../utils/api';
import { ProductDetailModal } from '../../components/product-detail-modal/product-detail-modal';

if (!localStorage.getItem('authToken')) {
    window.location.href = '/login.html';
}

const productTableBody = document.querySelector<HTMLTableSectionElement>('#product-table-body')!;
const addProductButton = document.querySelector<HTMLButtonElement>('#add-product-button')!;
const productModal = document.querySelector<HTMLDivElement>('#product-modal')!;
const closeModalButton = document.querySelector<HTMLButtonElement>('#close-modal-button')!;
const productForm = document.querySelector<HTMLFormElement>('#product-form')!;
const formTitle = document.querySelector<HTMLHeadingElement>('#form-title')!;
const productIdInput = document.querySelector<HTMLInputElement>('#product-id')!;
const nameInput = document.querySelector<HTMLInputElement>('#name')!;
const priceInput = document.querySelector<HTMLInputElement>('#price')!;
const stockInput = document.querySelector<HTMLInputElement>('#stock')!;
const descriptionInput = document.querySelector<HTMLTextAreaElement>('#description')!;
const imageInput = document.querySelector<HTMLInputElement>('#image-input')!;
const imagePreview = document.querySelector<HTMLDivElement>('#image-preview')!;
const userResponsibleInput = document.querySelector<HTMLInputElement>('#user-responsible')!;
const adminNameDisplay = document.querySelector<HTMLSpanElement>('#admin-name-display')!;
const logoutButton = document.querySelector<HTMLAnchorElement>('#logout-button')!;
const exportButton = document.querySelector<HTMLButtonElement>('#export-button')!;

const modalContainer = document.getElementById('modal-container')!;
const productDetailModal = new ProductDetailModal(modalContainer);

let selectedFile: File | null = null;
let adminName = 'Admin';

const token = localStorage.getItem('authToken');
if (token) {
    try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        adminName = payload.name || payload.email || 'Admin';
        if (adminNameDisplay) adminNameDisplay.textContent = adminName;
    } catch (e) {
        console.error('Gagal parse token:', e);
        if (adminNameDisplay) adminNameDisplay.textContent = 'Admin';
    }
}

// Definisi Fungsi
const showModal = (modal: HTMLElement) => modal.classList.remove('hidden');
const hideModal = (modal: HTMLElement) => modal.classList.add('hidden');

function updateImagePreview(imageUrl?: string | null) {
    if (!imagePreview) return;
    const placeholder = imagePreview.querySelector<HTMLElement>('.upload-placeholder'); 
    if (!placeholder) return;

    if (imageUrl) {
        imagePreview.style.backgroundImage = `url(${imageUrl})`;
        placeholder.style.display = 'none';
        imagePreview.classList.add('has-image');
    } else {
        imagePreview.style.backgroundImage = 'none';
        placeholder.style.display = 'flex';
        imagePreview.classList.remove('has-image');
    }
}

async function loadProducts() {
    try {
        const products = await apiFetch<Product[]>('/products');
        productTableBody.innerHTML = products.map((product, index) => {
            const imageCell = product.image_url ? `<img src="${product.image_url}" alt="${product.name}">` : `<span>No Image</span>`;
            
            const responsibleUserName = product.responsible_user?.name || 'N/A';

            return `
                <tr>
                    <td>${index + 1}</td>
                    <td><div class="product-cell">${imageCell}<span>${product.name}</span></div></td>
                    <td>${new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(product.price)}</td>
                    <td>${responsibleUserName}</td>
                    <td>${product.stock}</td>
                    <td>
                        <button class="icon-button view-btn" data-product='${JSON.stringify(product)}'><i class='bx bx-show'></i></button>
                        <button class="icon-button edit-btn" data-product='${JSON.stringify(product)}'><i class='bx bxs-edit'></i></button>
                    </td>
                </tr>`;
        }).join('');
    } catch (error) {
        console.error('Gagal memuat produk:', error);
        productTableBody.innerHTML = `<tr><td colspan="6">Gagal memuat data.</td></tr>`;
    }
}

function resetForm() {
    productForm.reset();
    productIdInput.value = '';
    formTitle.textContent = 'Tambah Produk';
    updateImagePreview(null);
    selectedFile = null;
}

function populateFormForEdit(product: Product) {
    resetForm();
    formTitle.textContent = 'Edit Produk';
    productIdInput.value = product.ID.toString();
    nameInput.value = product.name;
    priceInput.value = product.price.toString();
    stockInput.value = product.stock.toString();
    descriptionInput.value = product.description;
    
    userResponsibleInput.value = product.responsible_user?.name || adminName;
    updateImagePreview(product.image_url);
    showModal(productModal);
}

async function handleExport() {
    exportButton.textContent = 'Mengekspor...';
    exportButton.disabled = true;
    try {
        const products = await apiFetch<Product[]>('/products');
        const csvRows = ['No,ID,Nama Produk,Harga,Stok,Penanggung Jawab'];
        products.forEach((product, index) => {
            const row = [
                index + 1, product.ID, `"${product.name.replace(/"/g, '""')}"`,
                product.price, product.stock, 
                `"${product.responsible_user?.name || 'N/A'}"`
            ];
            csvRows.push(row.join(','));
        });
        const blob = new Blob([csvRows.join('\n')], { type: 'text/csv;charset=utf-8;' });
        const link = document.createElement('a');
        const url = URL.createObjectURL(blob);
        link.setAttribute('href', url);
        link.setAttribute('download', 'shoppinGO_products.csv');
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    } catch (error) {
        console.error('Gagal mengekspor data:', error);
        alert('Terjadi kesalahan saat mengekspor data.');
    } finally {
        exportButton.textContent = 'Export Produk';
        exportButton.disabled = false;
    }
}

addProductButton.addEventListener('click', () => {
    resetForm();
    userResponsibleInput.value = adminName;
    showModal(productModal);
});

closeModalButton.addEventListener('click', () => hideModal(productModal));
imagePreview.addEventListener('click', () => imageInput.click());
exportButton.addEventListener('click', handleExport);

imageInput.addEventListener('change', () => {
    if (imageInput.files && imageInput.files[0]) {
        selectedFile = imageInput.files[0];
        const reader = new FileReader();
        reader.onload = e => updateImagePreview(e.target?.result as string);
        reader.readAsDataURL(selectedFile);
    }
});

productForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const productData = {
        name: nameInput.value, price: parseInt(priceInput.value),
        stock: parseInt(stockInput.value), description: descriptionInput.value,
    };
    const isEditMode = !!productIdInput.value;
    const endpoint = isEditMode ? `/products/${productIdInput.value}` : '/products';
    const method = isEditMode ? 'PUT' : 'POST';
    
    try {
        const savedProduct = await apiFetch<Product>(endpoint, { method, body: JSON.stringify(productData) });
        
        if (selectedFile) {
            const idToUpload = isEditMode ? productIdInput.value : savedProduct.ID;
            const formData = new FormData();
            formData.append('image', selectedFile);
            await fetch(`http://localhost:8080/admin/products/${idToUpload}/upload`, {
                method: 'POST', headers: { 'Authorization': `Bearer ${token}` }, body: formData
            });
        }
        hideModal(productModal);
        await loadProducts();
    } catch (error) {
        alert(`Gagal menyimpan produk: ${error}`);
    }
});

productTableBody.addEventListener('click', (e) => {
    const target = e.target as HTMLElement;
    const editBtn = target.closest('.edit-btn');
    const viewBtn = target.closest('.view-btn');

    if (editBtn) {
        const product = JSON.parse(editBtn.getAttribute('data-product')!);
        populateFormForEdit(product);
    }
    if (viewBtn) {
        const product = JSON.parse(viewBtn.getAttribute('data-product')!);
        productDetailModal.show(product);
    }
});

logoutButton.addEventListener('click', (e) => {
    e.preventDefault();
    localStorage.removeItem('authToken');
    window.location.href = '/login.html';
});

loadProducts();