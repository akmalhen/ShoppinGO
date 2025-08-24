import './product-detail-modal.css';
import modalHTML from './product-detail-modal.html?raw';
import type { Product } from '../../types';

export class ProductDetailModal {
    private modalElement: HTMLElement;
    private viewProductName: HTMLHeadingElement;
    private viewProductImage: HTMLDivElement;
    private viewProductPrice: HTMLParagraphElement;
    private viewProductStock: HTMLParagraphElement;
    private viewProductDescription: HTMLParagraphElement;

    constructor(container: HTMLElement) {
        container.innerHTML = modalHTML;
        
        this.modalElement = container.querySelector('#view-product-modal')!;
        this.viewProductName = container.querySelector('#view-product-name')!;
        this.viewProductImage = container.querySelector('#view-product-image')!;
        this.viewProductPrice = container.querySelector('#view-product-price')!;
        this.viewProductStock = container.querySelector('#view-product-stock')!;
        this.viewProductDescription = container.querySelector('#view-product-description')!;
        
        const closeButton = container.querySelector('#close-view-modal-button')!;
        closeButton.addEventListener('click', () => this.hide());
    }

    public show(product: Product): void {
        this.viewProductName.textContent = product.name;
        this.viewProductPrice.textContent = new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(product.price);
        this.viewProductStock.textContent = product.stock.toString();
        this.viewProductDescription.textContent = product.description || 'Tidak ada deskripsi.';

        if (product.image_url) {
            this.viewProductImage.style.backgroundImage = `url(${product.image_url})`;
        } else {
            this.viewProductImage.style.backgroundImage = 'none';
        }

        this.modalElement.classList.remove('hidden');
    }

    public hide(): void {
        this.modalElement.classList.add('hidden');
    }
}