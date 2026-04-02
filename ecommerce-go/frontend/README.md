# Ecommerce Frontend

A React + Tailwind CSS storefront for the Go ecommerce backend.

## Setup

1. Open a terminal in `frontend/`
2. Install dependencies:

```bash
npm install
```

3. Create a `.env` file if needed and set API URL:

```env
VITE_API_URL=http://localhost:8080
```

4. Start the dev server:

```bash
npm run dev
```

5. Visit the app at `http://localhost:5173`

## Available scripts

- `npm run dev` - Start development server
- `npm run build` - Build production assets
- `npm run preview` - Preview production build

## Notes

- The frontend uses React Router for navigation.
- The app stores JWT tokens in `localStorage`.
- This frontend is ready for production build and deployment.
