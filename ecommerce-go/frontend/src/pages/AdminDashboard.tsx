import { useEffect, useState } from 'react';
import { fetchDashboardStats } from '../api';
import Loading from '../components/Loading';

export default function AdminDashboard() {
  const [stats, setStats] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const load = async () => {
      try {
        const data = await fetchDashboardStats();
        setStats(data);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    load();
  }, []);

  if (loading) return <Loading />;

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-7xl space-y-6">
        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <h1 className="text-3xl font-semibold text-slate-900">Admin dashboard</h1>
          <p className="mt-2 text-sm text-slate-600">View store performance and order insights.</p>
        </div>
        <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-4">
          <div className="rounded-[2rem] bg-slate-950 p-6 text-white shadow-soft">
            <p className="text-sm uppercase tracking-[.18em] text-slate-400">Total orders</p>
            <p className="mt-4 text-3xl font-semibold">{stats?.order_count ?? '0'}</p>
          </div>
          <div className="rounded-[2rem] bg-slate-950 p-6 text-white shadow-soft">
            <p className="text-sm uppercase tracking-[.18em] text-slate-400">Total revenue</p>
            <p className="mt-4 text-3xl font-semibold">${stats?.revenue ?? '0.00'}</p>
          </div>
          <div className="rounded-[2rem] bg-slate-950 p-6 text-white shadow-soft">
            <p className="text-sm uppercase tracking-[.18em] text-slate-400">Active users</p>
            <p className="mt-4 text-3xl font-semibold">{stats?.user_count ?? '0'}</p>
          </div>
          <div className="rounded-[2rem] bg-slate-950 p-6 text-white shadow-soft">
            <p className="text-sm uppercase tracking-[.18em] text-slate-400">Products</p>
            <p className="mt-4 text-3xl font-semibold">{stats?.product_count ?? '0'}</p>
          </div>
        </div>
      </div>
    </section>
  );
}
