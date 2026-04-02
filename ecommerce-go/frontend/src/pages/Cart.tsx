import { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { getCart, removeFromCart, updateCart } from '../api';
import { CartItem } from '../types';
import Loading from '../components/Loading';

export default function Cart() {
  const [items, setItems] = useState<CartItem[]>([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const load = async () => {
      try {
        const cartItems = await getCart();
        setItems(cartItems);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    load();
  }, []);

  const handleUpdate = async (productId: number, quantity: number) => {
    if (quantity < 1) return;
    try {
      await updateCart(productId, quantity);
      setItems((prev) => prev.map((item) => (item.product_id === productId ? { ...item, quantity } : item)));
    } catch (error) {
      console.error(error);
    }
  };

  const handleRemove = async (productId: number) => {
    try {
      await removeFromCart(productId);
      setItems((prev) => prev.filter((item) => item.product_id !== productId));
    } catch (error) {
      console.error(error);
    }
  };

  const total = items.reduce((sum, item) => sum + (item.product.discount_price ?? item.product.price) * item.quantity, 0);

  if (loading) return <Loading />;

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-7xl space-y-8">
        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h1 className="text-3xl font-semibold text-slate-900">Shopping cart</h1>
              <p className="mt-2 text-sm text-slate-600">Review items before checkout.</p>
            </div>
            <button
              type="button"
              onClick={() => navigate('/checkout')}
              className="rounded-full bg-slate-900 px-6 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700"
            >
              Proceed to checkout
            </button>
          </div>
        </div>

        {items.length === 0 ? (
          <div className="rounded-[2rem] bg-slate-100 p-10 text-center text-slate-600">
            Your cart is empty. <Link to="/products" className="font-semibold text-slate-900 underline">Browse products.</Link>
          </div>
        ) : (
          <div className="grid gap-8 lg:grid-cols-[1.5fr_0.5fr]">
            <div className="space-y-6 rounded-[2rem] bg-white p-6 shadow-soft">
              {items.map((item) => (
                <div key={item.id} className="flex flex-col gap-4 rounded-3xl border border-slate-200 p-5 md:flex-row md:items-center md:justify-between">
                  <div className="flex items-center gap-4">
                    <img src={item.product.image || 'https://via.placeholder.com/140'} alt={item.product.name} className="h-28 w-28 rounded-3xl object-cover" />
                    <div>
                      <h2 className="text-lg font-semibold text-slate-900">{item.product.name}</h2>
                      <p className="mt-2 text-sm text-slate-600">${item.product.discount_price ?? item.product.price} each</p>
                    </div>
                  </div>
                  <div className="flex flex-col gap-3 sm:items-end">
                    <input
                      type="number"
                      value={item.quantity}
                      min={1}
                      onChange={(e) => handleUpdate(item.product_id, Number(e.target.value))}
                      className="w-28 rounded-3xl border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none"
                    />
                    <button
                      type="button"
                      onClick={() => handleRemove(item.product_id)}
                      className="text-sm font-semibold text-slate-600 transition hover:text-slate-900"
                    >
                      Remove
                    </button>
                  </div>
                </div>
              ))}
            </div>
            <aside className="space-y-4 rounded-[2rem] bg-white p-6 shadow-soft">
              <h2 className="text-lg font-semibold text-slate-900">Order summary</h2>
              <div className="space-y-3 text-sm text-slate-600">
                <div className="flex items-center justify-between">
                  <span>Subtotal</span>
                  <span>${total.toFixed(2)}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span>Shipping</span>
                  <span>Calculated at checkout</span>
                </div>
              </div>
              <p className="rounded-3xl bg-slate-50 p-4 text-sm text-slate-700">Proceed to checkout for payment and shipping details.</p>
              <button
                type="button"
                onClick={() => navigate('/checkout')}
                className="w-full rounded-3xl bg-slate-900 px-5 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700"
              >
                Checkout
              </button>
            </aside>
          </div>
        )}
      </div>
    </section>
  );
}
