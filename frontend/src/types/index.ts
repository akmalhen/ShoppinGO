export interface Product {
  ID: number;
  name: string;
  price: number;
  stock: number;
  description: string;
  image_url: string | null;
  responsible_user_id: number;
  responsible_user: User; 
  CreatedAt: string;
}

export interface DashboardStats {
  total_users:        number;
  active_users:       number;
  total_products:     number;
  available_products: number;
  latest_products:    Product[];
}

export interface User {
    ID: number;
    name: string;   
    email: string;  
    is_active: boolean; 
    CreatedAt: string;  
}