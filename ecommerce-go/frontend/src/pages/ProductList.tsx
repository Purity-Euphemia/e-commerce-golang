import { useEffect, useMemo, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import { getCategories, getProducts } from '../api';
import { Category, Product } from '../types';
import ProductCard from '../components/ProductCard';
import Loading from '../components/Loading';

export default function ProductList() {
  const [searchParams, setSearchParams] = useSearchParams();
  const [products, setProducts] = useState<Product[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');

  const categoryId = searchParams.get('category');

  const queryString = useMemo(() => {
    const query = new URLSearchParams();
    query.set('page', '1');
    query.set('page_size', '12');
    if (search) query.set('search', search);
    if (categoryId) query.set('category_id', categoryId);
    return `?${query.toString()}`;
  }, [categoryId, search]);

  useEffect(() => {
    const load = async () => {
      setLoading(true);
      try {
        const [categoriesData, productsData] = await Promise.all([
          getCategories(),
          getProducts(queryString)
        ]);
        setCategories(categoriesData);
        setProducts(productsData);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [queryString]);

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-7xl space-y-8">
        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h1 className="text-3xl font-semibold text-slate-900">All products</h1>
              <p className="mt-2 text-sm text-slate-600">Browse and filter the entire catalog with search.</p>
            </div>
            <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
              <input
                type="text"
                placeholder="Search products..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                className="w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-slate-900 sm:w-80"
              />
            </div>
          </div>
        </div>

        <div className="grid gap-8 xl:grid-cols-[280px_1fr]">
          <aside className="space-y-6">
            <div className="rounded-[2rem] bg-white p-6 shadow-soft">
              <h2 className="text-lg font-semibold text-slate-900">Categories</h2>
              <div className="mt-4 grid gap-3">
                <button
                  type="button"
                  onClick={() => {
                    setSearchParams({});
                    setSearch('');
                  }}
                  className="w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-left text-sm font-medium text-slate-700 transition hover:bg-slate-100"
                >
                  All categories
                </button>
                {categories.map((category) => (
                  <button
                    key={category.id}
                    type="button"
                    onClick={() => setSearchParams({ category: String(category.id) })}
                    className="w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-left text-sm font-medium text-slate-700 transition hover:bg-slate-100"
                  >
                    {category.name}
                  </button>
                ))}
              </div>
            </div>
          </aside>

          <div className="space-y-6">
            {loading ? (
              <div className="rounded-[2rem] bg-white p-10 shadow-soft">
                <Loading />
              </div>
            ) : (
              <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
                {products.length ? products.map((product) => <ProductCard key={product.id} product={product} />) : (
                  <div className="rounded-[2rem] bg-slate-100 p-10 text-center text-slate-600">No products found.</div>
                )}
              </div>
            )}
          </div>
        </div>
      </div>
    </section>
  );
}
