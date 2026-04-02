import { type ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

const isAuthenticated = () => Boolean(localStorage.getItem('token'));

export default function ProtectedRoute({ children }: { children: ReactNode }) {
  if (!isAuthenticated()) {
    return <Navigate to="/login" replace />;
  }
  return <>{children}</>;
}
