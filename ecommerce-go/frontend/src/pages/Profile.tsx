import { useEffect, useState } from 'react';
import { getProfile } from '../api';
import Loading from '../components/Loading';

export default function Profile() {
  const [profile, setProfile] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const load = async () => {
      try {
        const data = await getProfile();
        setProfile(data);
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
      <div className="mx-auto max-w-3xl rounded-[2rem] bg-white p-8 shadow-soft">
        <h1 className="text-3xl font-semibold text-slate-900">My profile</h1>
        <p className="mt-2 text-sm text-slate-600">View your account details and order status.</p>
        <div className="mt-8 grid gap-6 sm:grid-cols-2">
          <div className="rounded-3xl bg-slate-50 p-6">
            <p className="text-sm uppercase tracking-[.18em] text-slate-500">Name</p>
            <p className="mt-3 text-lg font-semibold text-slate-900">{profile?.name || 'N/A'}</p>
          </div>
          <div className="rounded-3xl bg-slate-50 p-6">
            <p className="text-sm uppercase tracking-[.18em] text-slate-500">Email</p>
            <p className="mt-3 text-lg font-semibold text-slate-900">{profile?.email || 'N/A'}</p>
          </div>
          <div className="rounded-3xl bg-slate-50 p-6">
            <p className="text-sm uppercase tracking-[.18em] text-slate-500">Phone</p>
            <p className="mt-3 text-lg font-semibold text-slate-900">{profile?.phone || 'N/A'}</p>
          </div>
          <div className="rounded-3xl bg-slate-50 p-6">
            <p className="text-sm uppercase tracking-[.18em] text-slate-500">Shipping address</p>
            <p className="mt-3 text-lg font-semibold text-slate-900">{profile?.address || 'N/A'}</p>
          </div>
        </div>
      </div>
    </section>
  );
}
