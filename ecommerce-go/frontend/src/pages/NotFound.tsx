import { Link } from 'react-router-dom';

export default function NotFound() {
  return (
    <section className="flex min-h-[70vh] items-center justify-center bg-slate-50 px-4 py-10 sm:px-6 lg:px-8">
      <div className="w-full max-w-2xl rounded-[2rem] border border-slate-200 bg-white p-10 text-center shadow-soft">
        <p className="text-sm uppercase tracking-[.3em] text-slate-400">Page not found</p>
        <h1 className="mt-6 text-5xl font-semibold text-slate-900">404</h1>
        <p className="mt-4 text-base text-slate-600">The page you are looking for doesn’t exist, or it has been moved.</p>
        <Link to="/" className="mt-8 inline-flex rounded-full bg-slate-900 px-6 py-3 text-sm font-semibold text-white transition hover:bg-slate-700">
          Return home
        </Link>
      </div>
    </section>
  );
}
