import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { addToCart, getProduct } from '../api';
import { Product } from '../types';
import Loading from '../components/Loading';

export default function ProductDetail() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [product, setProduct] = useState<Product | null>(null);
  const [quantity, setQuantity] = useState(1);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const load = async () => {
      if (id) {
        try {
          const data = await getProduct(id);
          setProduct(data);
        } catch (error) {
          console.error(error);
        } finally {
          setLoading(false);
        }
      }
    };
    load();
  }, [id]);

  const handleAddToCart = async () => {
    if (!product) return;
    try {
      await addToCart(product.id, quantity);
      navigate('/cart');
    } catch (error) {
      console.error(error);
      alert('Please login or try again.');
    }
  };

  if (loading) return <Loading />;

  if (!product) {
    return <div className="px-4 py-10 text-center text-slate-700">Product not found.</div>;
  }

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto grid max-w-6xl gap-10 lg:grid-cols-[1.2fr_0.8fr]">
        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <img src={product.image || 'https://via.placeholder.com/720x480'} alt={product.name} className="w-full rounded-[1.75rem] object-cover" />
          <div className="mt-8 space-y-4">
            <h1 className="text-3xl font-semibold text-slate-900">{product.name}</h1>
            <p className="text-lg text-slate-600">{product.description || 'A premium product ready for your customers.'}</p>
            <div className="flex flex-wrap items-center gap-3 text-slate-900">
              <span className="text-3xl font-bold">${product.discount_price ?? product.price}</span>
              {product.discount_price ? <span className="text-sm text-slate-500 line-through">${product.price}</span> : null}
            </div>
          </div>
        </div>

        <div className="space-y-6 rounded-[2rem] bg-white p-8 shadow-soft">
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <p className="text-sm uppercase tracking-[.2em] text-slate-500">Product details</p>
              <span className="text-sm font-semibold text-slate-900">{product.stock ?? 0} in stock</span>
            </div>
            <div className="rounded-3xl border border-slate-200 bg-slate-50 p-4">
              <p className="text-sm text-slate-600">SKU: {product.sku || 'SW-001'}</p>
              <p className="mt-2 text-sm text-slate-600">Category ID: {product.category_id || 'N/A'}</p>
            </div>
          </div>
          <div className="space-y-4">
            <label className="block text-sm font-medium text-slate-700">Quantity</label>
            <input
              type="number"
              value={quantity}
              min={1}
              onChange={(e) => setQuantity(Number(e.target.value))}
              className="w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
            />
          </div>
          <button
            type="button"
            onClick={handleAddToCart}
            className="w-full rounded-3xl bg-slate-900 px-5 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700"
          >
            Add to cart
          </button>
        </div>
      </div>
    </section>
  );
}
