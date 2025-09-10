# Mini E-Commerce Platform


A full-stack, production-style MVP e-commerce platform built with Go (Fiber), Vue 3, and PostgreSQL. Features secure authentication, product management, order processing, and M-Pesa payment simulation. Frontend uses Vite, Pinia, TypeScript, and is ready for deployment on Vercel. Backend is ready for deployment on Render.com.

---

## � Project Status

- MVP complete: All core features implemented and production-ready
- Order totals, product names, and cart logic fully aligned between backend and frontend
- Profile page, admin dashboard, and error handling implemented
- Frontend and backend ready for cloud deployment

---

## �🚀 Project Overview

- **Product Vision:**
  - Enable secure product management, customer shopping, and integrated M-Pesa payment processing with real-time updates.
- **Success Metrics:**
  - Complete user authentication with role-based access
  - Functional product management
  - Working checkout and payment simulation
  - Responsive UI (desktop & mobile)
  - Clean, maintainable codebase

---

## 🧑‍💻 Tech Stack
- **Backend:** Go (Fiber), PostgreSQL, JWT
- **Frontend:** Vue.js 3 (Vite, TypeScript, Tailwind CSS)
- **Other:** M-Pesa payment simulation, RESTful API

---

## 🗂️ Monorepo Structure
```
go-ecom/
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── uploads/
│   ├── docs/
│   └── go.mod
├── frontend/
│   ├── src/
│   ├── public/
│   ├── package.json
│   ├── vercel.json
│   └── .env.example
└── README.md
```

---

## 👤 User Personas & Use Cases

### Admin
- Manage product inventory
- Monitor orders & transactions
- Full CRUD on products
- View/manage all orders

### Customer
- Register, log in, manage account
- Browse products, add to cart, checkout
- Make payments (M-Pesa simulation)
- View order history & payment status

---

## 🔒 Core Features

### Authentication & Authorization
- JWT-based registration & login
- Role-based access (admin/customer)

### Product Management
- Admin: Create, update, delete (soft), upload images
- Customer: Browse, view details

### Orders & Checkout
- Place orders, view order history
- Admin: View/manage all orders, update status

### Payment Integration
- M-Pesa payment simulation (mock API)
- Real-time transaction status

---

## 🏗️ Getting Started

### Prerequisites
- Go 1.24+
- Node.js 20+
- PostgreSQL 14+
- Git
- Postman (for API testing)

### Backend Setup
```sh
cd backend
cp .env.example .env # Edit DB credentials, JWT secret, etc.
go mod tidy
# Create DB & run migrations
psql -U <user> -d <db> -f migrations/001_initial_schema.sql
# Start server
go run cmd/main.go
```

### Frontend Setup
```sh
cd frontend
npm install
npm run dev
```

---


## 🗄️ Database Schema (Simplified)
- **users:** id, email, password_hash, full_name, role, created_at, updated_at
- **products:** id, name, description, price, stock, category, image_url
- **orders:** id, user_id, status, total_amount, created_at, updated_at
- **order_items:** id, order_id, product_id, product_name, quantity, unit_price
- **transactions:** id, order_id, mpesa_ref, status, amount, created_at

---

## 🚀 Deployment

### Backend (Render.com or similar)
1. Set up PostgreSQL and create your database.
2. Configure environment variables in Render.com dashboard (see `.env.example`).
3. Deploy backend from your repo; Render will run `go run cmd/main.go`.
4. Migrations: Run SQL files in `backend/migrations/` as needed.

### Frontend (Vercel)
1. Copy `.env.example` to `.env` and set `VITE_API_BASE_URL` to your backend API URL.
2. Ensure `vercel.json` exists for SPA routing:
   ```json
   {
     "rewrites": [
       { "source": "/(.*)", "destination": "/" }
     ]
   }
   ```
3. Push to GitHub and connect repo to Vercel.
4. Vercel will auto-detect Vite and build/deploy your app.

---

## 🛠️ Environment Variables

- Backend: See `backend/.env.example` for DB, JWT, etc.
- Frontend: See `frontend/.env.example` for `VITE_API_BASE_URL` (must point to deployed backend API)

---

## 📝 Contribution & Commit Guidelines

- Use [Conventional Commits](https://www.conventionalcommits.org/) (e.g., `feat:`, `fix:`, `chore:`)
- Commit in logical, incremental steps (see commit plan in project docs)
- Write clear PR descriptions and reference issues where possible

---

---

## 📚 API Overview
- `/api/auth/register` — Register user
- `/api/auth/login` — Login
- `/api/products` — List/browse products
- `/api/orders` — Place/view orders
- `/api/mpesa/stkpush` — Simulate payment
- `/api/admin/orders` — Admin order management

---

## 📝 Developer Notes
- Use `.env` for config (DB, JWT, etc.)
- All endpoints return JSON
- Use Postman for API testing
- See `Mini E-Commerce Platform - PRD.md` and `Mini E-Commerce Platform - Developer Implementation Guide.md` for full details

---

## 🌐 Credits

Created by Cleve-codes, 2025

---

## 📄 License
MIT
