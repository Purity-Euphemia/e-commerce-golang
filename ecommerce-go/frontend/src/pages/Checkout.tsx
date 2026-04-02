import { useState, type FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';
import { checkout } from '../api';

export default function Checkout() {
  const [address, setAddress] = useState('');
  const [coupon, setCoupon] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setLoading(true);
    try {
      const { data } = await checkout({ shipping_address: address, coupon_code: coupon || undefined });
      setMessage(data.message || 'Order created successfully.');
      setTimeout(() => navigate('/'), 1200);
    } catch (error) {
      console.error(error);
      setMessage('Unable to complete checkout. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-3xl space-y-8 rounded-[2rem] bg-white p-8 shadow-soft">
        <div>
          <h1 className="text-3xl font-semibold text-slate-900">Checkout</h1>
          <p className="mt-2 text-sm text-slate-600">Complete your purchase with shipping and coupon support.</p>
        </div>
        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-slate-700">Shipping address</label>
            <textarea
              required
              rows={4}
              value={address}
              onChange={(e) => setAddress(e.target.value)}
              className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none focus:border-slate-900"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-slate-700">Coupon code</label>
            <input
              value={coupon}
              onChange={(e) => setCoupon(e.target.value)}
              className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
              placeholder="Enter coupon code"
            />
          </div>
          {message && <div className="rounded-3xl bg-slate-100 p-4 text-sm text-slate-700">{message}</div>}
          <button
            type="submit"
            disabled={loading}
            className="w-full rounded-3xl bg-slate-900 px-6 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700 disabled:cursor-not-allowed disabled:bg-slate-400"
          >
            {loading ? 'Processing...' : 'Place order'}
          </button>
        </form>
      </div>
    </section>
  );
}
