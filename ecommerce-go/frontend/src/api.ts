import axios from 'axios';
import { CartItem, Category, LoginPayload, Order, Product, Review, User } from './types';

const baseURL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

const client = axios.create({
  baseURL,
  headers: {
    'Content-Type': 'application/json'
  }
});

client.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token && config.headers) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const getProducts = async (query = '') => {
  const response = await client.get<{ data: Product[] }>(`/products${query}`);
  return response.data.data;
};

export const getProduct = async (id: string) => {
  const response = await client.get<{ data: Product }>(`/products/${id}`);
  return response.data.data;
};

export const getCategories = async () => {
  const response = await client.get<{ data: Category[] }>('/categories');
  return response.data.data;
};

export const login = async (payload: LoginPayload) => {
  return client.post('/login', payload);
};

export const register = async (payload: { name: string; email: string; password: string }) => {
  return client.post('/register', payload);
};

export const getProfile = async () => {
  const response = await client.get<{ data: User }>('/profile');
  return response.data.data;
};

export const getCart = async () => {
  const response = await client.get<{ data: { cart: { items: CartItem[] } } }>('/cart');
  return response.data.data.cart.items;
};

export const addToCart = async (productId: number, quantity = 1) => {
  return client.post('/cart', { product_id: productId, quantity });
};

export const updateCart = async (productId: number, quantity: number) => {
  return client.put('/cart', { product_id: productId, quantity });
};

export const getProductReviews = async (productId: number) => {
  const response = await client.get<{ data: Review[] }>(`/reviews/product/${productId}`);
  return response.data.data;
};

export const addReview = async (productId: number, payload: { rating: number; title?: string; comment: string }) => {
  return client.post(`/reviews/product/${productId}`, payload);
};

export const removeFromCart = async (productId: number) => {
  return client.delete(`/cart/${productId}`);
};

export const checkout = async (payload: { shipping_address: string; coupon_code?: string }) => {
  return client.post('/checkout', payload);
};

export const fetchDashboardStats = async () => {
  const response = await client.get('/admin/dashboard/stats');
  return response.data;
};

export const fetchOrders = async () => {
  const response = await client.get<{ data: Order[] }>('/my-orders');
  return response.data.data;
};
