import { NavLink, useNavigate } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';

const navClasses = ({ isActive }: { isActive: boolean }) =>
  `transition text-slate-700 hover:text-slate-900 ${isActive ? 'font-semibold text-slate-900' : 'font-medium'}`;

export default function Header() {
  const { token, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <header className="bg-white shadow-sm sticky top-0 z-30">
      <div className="mx-auto flex max-w-7xl items-center justify-between px-4 py-4 sm:px-6">
        <div>
          <NavLink to="/" className="text-xl font-bold text-slate-900">
            ShopWave
          </NavLink>
        </div>
        <nav className="hidden gap-6 md:flex">
          <NavLink to="/products" className={navClasses}>
            Products
          </NavLink>
          <NavLink to="/cart" className={navClasses}>
            Cart
          </NavLink>
          <NavLink to="/profile" className={navClasses}>
            Profile
          </NavLink>
        </nav>
        <div className="flex items-center gap-3">
          {token ? (
            <button
              type="button"
              onClick={handleLogout}
              className="rounded-full bg-slate-900 px-4 py-2 text-sm font-semibold text-white transition hover:bg-slate-700"
            >
              Logout
            </button>
          ) : (
            <>
              <NavLink to="/login" className="rounded-full bg-slate-900 px-4 py-2 text-sm font-semibold text-white transition hover:bg-slate-700">
                Login
              </NavLink>
              <NavLink to="/register" className="rounded-full border border-slate-300 bg-white px-4 py-2 text-sm font-semibold text-slate-900 transition hover:border-slate-900">
                Register
              </NavLink>
            </>
          )}
        </div>
      </div>
    </header>
  );
}
