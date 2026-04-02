import { useState, type FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';
import { login } from '../api';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setLoading(true);
    setError('');

    try {
      const response = await login({ email, password });
      const token = response.data.data.token || response.data.token || null;
      if (token) {
        localStorage.setItem('token', token);
        window.dispatchEvent(new Event('authChange'));
        navigate('/');
      } else {
        setError('Login failed, please verify your credentials.');
      }
    } catch (error) {
      console.error(error);
      setError('Unable to login. Please check your email and password.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-md rounded-[2rem] bg-white p-8 shadow-soft">
        <h1 className="text-3xl font-semibold text-slate-900">Login</h1>
        <p className="mt-2 text-sm text-slate-600">Welcome back. Enter your credentials to access your account.</p>
        <form onSubmit={handleSubmit} className="mt-8 space-y-6">
          <div>
            <label className="block text-sm font-medium text-slate-700">Email</label>
            <input
              type="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-slate-700">Password</label>
            <input
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
            />
          </div>
          {error && <div className="rounded-3xl bg-rose-50 px-4 py-3 text-sm text-rose-700">{error}</div>}
          <button
            type="submit"
            disabled={loading}
            className="w-full rounded-3xl bg-slate-900 px-6 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700 disabled:cursor-not-allowed disabled:bg-slate-400"
          >
            {loading ? 'Signing in...' : 'Sign in'}
          </button>
        </form>
      </div>
    </section>
  );
}
