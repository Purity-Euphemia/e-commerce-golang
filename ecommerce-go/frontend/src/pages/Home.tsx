import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { getCategories, getProducts } from '../api';
import { Category, Product } from '../types';
import ProductCard from '../components/ProductCard';
import Loading from '../components/Loading';

export default function Home() {
  const [products, setProducts] = useState<Product[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const load = async () => {
      try {
        const [fetchedCategories, fetchedProducts] = await Promise.all([
          getCategories(),
          getProducts('?page=1&page_size=8')
        ]);
        setCategories(fetchedCategories);
        setProducts(fetchedProducts);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    load();
  }, []);

  return (
    <section className="space-y-12 px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-7xl rounded-[2rem] bg-gradient-to-r from-slate-950 to-slate-800 px-8 py-16 text-white shadow-soft">
        <div className="max-w-3xl space-y-6">
          <p className="inline-flex items-center rounded-full bg-white/10 px-4 py-1 text-xs uppercase tracking-[.25em] text-slate-200">Modern ecommerce experience</p>
          <h1 className="text-4xl font-semibold tracking-tight sm:text-5xl">Launch your store with a premium shopping experience.</h1>
          <p className="text-lg leading-8 text-slate-300">Browse products, add to cart, apply coupons, and manage orders with a responsive frontend built for production.</p>
          <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
            <Link to="/products" className="inline-flex items-center justify-center rounded-full bg-white px-6 py-3 text-sm font-semibold text-slate-950 shadow-lg shadow-slate-900/10 transition hover:bg-slate-100">
              Browse products
            </Link>
            <Link to="/login" className="inline-flex items-center justify-center rounded-full border border-white/20 bg-white/5 px-6 py-3 text-sm font-semibold text-white transition hover:border-white hover:bg-white/10">
              Login or Register
            </Link>
          </div>
        </div>
      </div>

      <div className="grid gap-8 lg:grid-cols-[1.25fr_0.75fr]">
        <div className="space-y-6">
          <div className="rounded-3xl bg-white p-6 shadow-soft">
            <div className="flex items-center justify-between gap-4">
              <div>
                <h2 className="text-xl font-semibold text-slate-900">Featured products</h2>
                <p className="mt-1 text-sm text-slate-600">Top items ready for your store launch.</p>
              </div>
              <Link to="/products" className="text-sm font-semibold text-slate-700 hover:text-slate-900">View all</Link>
            </div>
            {loading ? (
              <Loading />
            ) : (
              <div className="mt-6 grid gap-4 sm:grid-cols-2 lg:grid-cols-2">
                {products.map((product) => (
                  <ProductCard key={product.id} product={product} />
                ))}
              </div>
            )}
          </div>
        </div>

        <aside className="space-y-6">
          <div className="rounded-3xl bg-white p-6 shadow-soft">
            <h2 className="text-xl font-semibold text-slate-900">Shop by category</h2>
            <div className="mt-5 grid gap-3">
              {categories.slice(0, 6).map((category) => (
                <Link
                  key={category.id}
                  to={`/products?category=${category.id}`}
                  className="rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm font-medium text-slate-700 transition hover:border-slate-300 hover:bg-slate-100"
                >
                  {category.name}
                </Link>
              ))}
            </div>
          </div>

          <div className="rounded-3xl bg-slate-950 p-6 text-white shadow-soft">
            <h2 className="text-xl font-semibold">Design ready for production</h2>
            <p className="mt-4 text-sm leading-6 text-slate-300">This frontend is built with Tailwind CSS and modern React practices so it is fast, maintainable, and visually polished.</p>
          </div>
        </aside>
      </div>
    </section>
  );
}
