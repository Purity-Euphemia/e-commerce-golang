import { Link } from 'react-router-dom';
import { Product } from '../types';

export default function ProductCard({ product }: { product: Product }) {
  return (
    <div className="group overflow-hidden rounded-3xl border border-slate-200 bg-white shadow-soft transition hover:-translate-y-1 hover:shadow-xl">
      <Link to={`/products/${product.id}`} className="block overflow-hidden">
        <div className="flex h-64 items-center justify-center bg-slate-100">
          <img
            src={product.image || 'https://via.placeholder.com/360x240'}
            alt={product.name}
            className="h-full w-full object-cover transition duration-300 group-hover:scale-105"
          />
        </div>
      </Link>
      <div className="space-y-3 p-5">
        <div className="flex items-center justify-between gap-3">
          <h3 className="text-lg font-semibold text-slate-900">{product.name}</h3>
          <span className="rounded-full bg-slate-100 px-2 py-1 text-xs font-semibold uppercase tracking-wide text-slate-600">
            {product.rating ? product.rating.toFixed(1) : 'New'}
          </span>
        </div>
        <p className="text-sm leading-6 text-slate-600">{product.description || 'Delightful product for everyday life.'}</p>
        <div className="flex items-center justify-between gap-4">
          <div className="text-xl font-bold text-slate-900">${product.discount_price ?? product.price}</div>
          <Link to={`/products/${product.id}`} className="text-sm font-semibold text-slate-700 transition hover:text-slate-900">
            View product
          </Link>
        </div>
      </div>
    </div>
  );
}
