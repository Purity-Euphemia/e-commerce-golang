export default function Footer() {
  return (
    <footer className="bg-slate-950 text-slate-200">
      <div className="mx-auto flex max-w-7xl flex-col gap-8 px-4 py-10 text-sm sm:px-6 md:flex-row md:items-center md:justify-between">
        <div>
          <p className="text-base font-semibold text-white">ShopWave</p>
          <p className="mt-2 text-slate-400">Modern ecommerce experience built with React + Tailwind.</p>
        </div>
        <div className="flex flex-wrap gap-4 text-slate-400">
          <span>© 2026 ShopWave</span>
          <span>Built for production</span>
          <span>React + Tailwind</span>
        </div>
      </div>
    </footer>
  );
}
