export interface Product {
  id: number;
  name: string;
  description?: string;
  price: number;
  discount_price?: number;
  image?: string;
  category_id?: number;
  sku?: string;
  rating?: number;
  stock?: number;
}

export interface Category {
  id: number;
  name: string;
  slug: string;
}

export interface User {
  id: number;
  name: string;
  email: string;
  phone?: string;
  address?: string;
  city?: string;
  state?: string;
  zip_code?: string;
}

export interface CartItem {
  id: number;
  product_id: number;
  quantity: number;
  product: Product;
}

export interface Order {
  id: number;
  order_number: string;
  total_amount: number;
  status: string;
}

export interface LoginPayload {
  email: string;
  password: string;
}
