import { useEffect, useState } from 'react';

const TOKEN_KEY = 'token';
const PROFILE_KEY = 'profile';

export function useAuth() {
  const [token, setTokenState] = useState<string | null>(null);
  const [profile, setProfileState] = useState<any>(null);

  const refreshAuth = () => {
    const stored = localStorage.getItem(TOKEN_KEY);
    const user = localStorage.getItem(PROFILE_KEY);
    setTokenState(stored);
    setProfileState(user ? JSON.parse(user) : null);
  };

  useEffect(() => {
    refreshAuth();
    window.addEventListener('authChange', refreshAuth);
    return () => window.removeEventListener('authChange', refreshAuth);
  }, []);

  const login = (tokenValue: string, user?: any) => {
    localStorage.setItem(TOKEN_KEY, tokenValue);
    if (user) localStorage.setItem(PROFILE_KEY, JSON.stringify(user));
    setTokenState(tokenValue);
    setProfileState(user || null);
    window.dispatchEvent(new Event('authChange'));
  };

  const logout = () => {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(PROFILE_KEY);
    setTokenState(null);
    setProfileState(null);
    window.dispatchEvent(new Event('authChange'));
  };

  const isAuthenticated = Boolean(token);

  return { token, profile, login, logout, isAuthenticated };
}
