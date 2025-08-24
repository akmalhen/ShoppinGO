import Swiper from 'swiper';
import { Navigation, Pagination, Autoplay } from 'swiper/modules';

import "../../styles/landing.css";
import type { Product } from "../../types";
import { apiFetch } from "../../utils/api";
import { ProductDetailModal } from '../../components/product-detail-modal/product-detail-modal';

const modalContainer = document.getElementById('modal-container')!;
const productDetailModal = new ProductDetailModal(modalContainer);
const newProductsContainer = document.querySelector<HTMLDivElement>('#new-products-slider')!;
const allProductsContainer = document.querySelector<HTMLDivElement>('#all-products-grid')!;

function createProductCard(product: Product): HTMLElement {
    const card = document.createElement('div');
    card.className = 'product-card';
    card.dataset.productId = product.ID.toString();

    const formattedPrice = new Intl.NumberFormat('id-ID', {
        style: 'currency', currency: 'IDR', minimumFractionDigits: 0
    }).format(product.price);

    const imageContentHTML = product.image_url
        ? `<img src="${product.image_url}" alt="${product.name}" class="product-image">`
        : `<i class='bx bx-image-alt placeholder-icon'></i>`;

    card.innerHTML = `
        <div class="product-image-container">
            ${imageContentHTML}
        </div>
        <h3 class="product-name">${product.name}</h3>
        <p class="product-price">${formattedPrice}</p>
    `;
    return card;
}

async function fetchAndDisplayProducts() {
    try {
        const newProducts: Product[] = await apiFetch('/products/latest');
        const allProducts: Product[] = await apiFetch('/products/available');

        const newProductsWrapper = document.querySelector<HTMLDivElement>('#new-products-slider');
        if (newProductsWrapper) newProductsWrapper.innerHTML = '';
        allProductsContainer.innerHTML = '';

        newProducts.forEach(product => newProductsWrapper?.appendChild(createProductCard(product)));
        allProducts.forEach(product => allProductsContainer.appendChild(createProductCard(product)));

    } catch (error) {
        console.error('Terjadi kesalahan:', error);
        allProductsContainer.innerHTML = `<p>Gagal memuat produk.</p>`;
    }
}

function handleProductClick(event: MouseEvent) {
    const target = event.target as HTMLElement;
    const card = target.closest<HTMLDivElement>('.product-card');
    
    if (card && card.dataset.productId) {
        const productId = card.dataset.productId;
        apiFetch<Product>(`/products/${productId}`)
            .then(product => {
                productDetailModal.show(product);
            })
            .catch(err => {
                console.error('Gagal mengambil detail produk:', err);
                alert('Gagal memuat detail produk.');
            });
    }
}

Swiper.use([Navigation, Pagination, Autoplay]);

new Swiper('.swiper', {
  loop: true,
  autoplay: {
    delay: 4000,
    disableOnInteraction: false,
  },
  pagination: {
    el: '.swiper-pagination',
    clickable: true,
  },
  navigation: {
    nextEl: '.swiper-button-next',
    prevEl: '.swiper-button-prev',
  },
});

// Event Listeners untuk klik produk
newProductsContainer.addEventListener('click', handleProductClick);
allProductsContainer.addEventListener('click', handleProductClick);

// Panggil fungsi utama untuk memuat data
fetchAndDisplayProducts();