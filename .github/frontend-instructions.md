# ScoreThatExam Frontend - AI Agent Context Guide

---

## 🤖 AI Agent Role

**You are an Expert Frontend Developer and UI/UX Specialist** working on the ScoreThatExam project.

### Your Responsibilities:
1. **Maintain Design Consistency**: Follow the Neu-Brutalist design system strictly
2. **Component-Based Architecture**: Build reusable, accessible components
3. **Type Safety**: Use TypeScript strictly, no `any` types
4. **Responsive Design**: Mobile-first approach, all layouts must work on all screen sizes
5. **Performance-Oriented**: Optimize bundle size, lazy load routes, minimize re-renders

### When Contributing Code:
- Follow the component structure: `components/`, `pages/`, `hooks/`, `services/`
- Use TypeScript interfaces for all props and API responses
- Implement proper loading and error states
- Ensure keyboard navigation and accessibility (WCAG 2.1 AA)
- Write semantic HTML

---

## What This Project Is

**AceThatPaper** is a **React-based frontend** for an educational quiz platform designed to help Nigerian students prepare for the **JAMB (Joint Admissions and Matriculation Board)** examination. The platform provides quiz functionality powered by JAMB syllabus and past questions.

---

## Technology Stack

| Component | Technology |
|-----------|------------|
| **Framework** | React 18+ with Vite |
| **Language** | TypeScript (strict mode) |
| **Styling** | TailwindCSS + Custom Neu-Brutalist Design System |
| **Routing** | React Router v6 |
| **State Management** | React Context + useReducer (or Zustand if needed) |
| **HTTP Client** | Axios |
| **Forms** | React Hook Form + Zod validation |
| **Icons** | Lucide React |
| **Date Handling** | date-fns |

---

## 🎨 NEU-BRUTALIST DESIGN SYSTEM

### Design Philosophy

The UI follows a **Neu-Brutalist** aesthetic—bold, raw, unapologetic, and intentionally "undesigned." The design should feel **loud, confident, and standout**, eschewing subtlety in favor of boldness.

### Core Design Principles

1. **Bold & High-Contrast**: Use stark color combinations that demand attention
2. **Thick Borders**: All major elements have 3-4px solid black borders
3. **Harsh Drop Shadows**: Hard-edged shadows (no blur), offset 4-8px
4. **Raw Typography**: Bold, slightly aggressive fonts
5. **Asymmetrical Layouts**: Intentionally break the grid occasionally
6. **Flat Colors**: No gradients, no subtle effects
7. **Unapologetic Presence**: Every element should feel intentional and bold

---

### Color Palette

```css
:root {
  /* Primary Colors */
  --color-primary: #FF6B35;        /* Bright Orange */
  --color-primary-dark: #E55A2B;   /* Darker Orange */
  
  /* Secondary Colors */
  --color-secondary: #004E89;      /* Deep Blue */
  --color-secondary-light: #1A6FB5; /* Lighter Blue */
  
  /* Accent Colors */
  --color-accent-yellow: #FFD23F;  /* Bold Yellow */
  --color-accent-green: #3DDC84;   /* Success Green */
  --color-accent-red: #FF3366;     /* Error Red */
  --color-accent-purple: #7B2CBF;  /* Purple Accent */
  
  /* Neutrals */
  --color-black: #000000;
  --color-white: #FFFFFF;
  --color-cream: #FFF8E7;          /* Off-white background */
  --color-gray-light: #E5E5E5;
  --color-gray: #888888;
  
  /* Semantic Colors */
  --color-success: #3DDC84;
  --color-error: #FF3366;
  --color-warning: #FFD23F;
  --color-info: #004E89;
}
```

### TailwindCSS Configuration

```javascript
// tailwind.config.js
module.exports = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: "#FF6B35",
          dark: "#E55A2B",
        },
        secondary: {
          DEFAULT: "#004E89",
          light: "#1A6FB5",
        },
        accent: {
          yellow: "#FFD23F",
          green: "#3DDC84",
          red: "#FF3366",
          purple: "#7B2CBF",
        },
        cream: "#FFF8E7",
      },
      fontFamily: {
        display: ['"Space Grotesk"', 'Arial Black', 'sans-serif'],
        body: ['"Space Grotesk"', 'Arial', 'sans-serif'],
        mono: ['"JetBrains Mono"', 'monospace'],
      },
      boxShadow: {
        'brutal': '4px 4px 0px 0px #000000',
        'brutal-sm': '2px 2px 0px 0px #000000',
        'brutal-lg': '8px 8px 0px 0px #000000',
        'brutal-primary': '4px 4px 0px 0px #FF6B35',
        'brutal-secondary': '4px 4px 0px 0px #004E89',
      },
      borderWidth: {
        '3': '3px',
        '4': '4px',
      },
    },
  },
  plugins: [],
}
```

---

### Typography

```css
/* Font Imports - Add to index.css */
@import url('https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@400;500;600;700&display=swap');

/* Typography Scale */
.text-display-xl { font-size: 4rem; line-height: 1.1; font-weight: 700; }    /* 64px - Hero */
.text-display-lg { font-size: 3rem; line-height: 1.2; font-weight: 700; }    /* 48px - Page titles */
.text-display-md { font-size: 2.25rem; line-height: 1.2; font-weight: 700; } /* 36px - Section titles */
.text-heading-lg { font-size: 1.875rem; line-height: 1.3; font-weight: 600; } /* 30px */
.text-heading-md { font-size: 1.5rem; line-height: 1.4; font-weight: 600; }   /* 24px */
.text-heading-sm { font-size: 1.25rem; line-height: 1.4; font-weight: 600; }  /* 20px */
.text-body-lg { font-size: 1.125rem; line-height: 1.6; font-weight: 400; }    /* 18px */
.text-body { font-size: 1rem; line-height: 1.6; font-weight: 400; }           /* 16px */
.text-body-sm { font-size: 0.875rem; line-height: 1.5; font-weight: 400; }    /* 14px */
.text-caption { font-size: 0.75rem; line-height: 1.4; font-weight: 500; }     /* 12px */
```

---

### Component Specifications

#### Buttons

```tsx
// Primary Button
<button className="
  px-6 py-3
  bg-primary text-black
  font-display font-bold text-lg uppercase
  border-4 border-black
  shadow-brutal
  hover:bg-accent-yellow hover:shadow-brutal-lg
  active:shadow-none active:translate-x-1 active:translate-y-1
  transition-all duration-100
">
  Start Quiz
</button>

// Secondary Button
<button className="
  px-6 py-3
  bg-secondary text-white
  font-display font-bold text-lg uppercase
  border-4 border-black
  shadow-brutal
  hover:bg-secondary-light
  active:shadow-none active:translate-x-1 active:translate-y-1
">
  View Leaderboard
</button>

// Outline Button
<button className="
  px-6 py-3
  bg-white text-black
  font-display font-bold text-lg uppercase
  border-4 border-black
  shadow-brutal
  hover:bg-black hover:text-white
  active:shadow-none active:translate-x-1 active:translate-y-1
">
  Cancel
</button>

// Danger Button
<button className="
  px-6 py-3
  bg-accent-red text-white
  font-display font-bold text-lg uppercase
  border-4 border-black
  shadow-brutal
  hover:bg-red-700
  active:shadow-none active:translate-x-1 active:translate-y-1
">
  Delete Account
</button>
```

**Button States:**
- **Default**: Flat color + brutal shadow
- **Hover**: Hard color swap (no fade) + larger shadow
- **Active/Pressed**: Shadow removed, element shifts down-right
- **Disabled**: Gray background, no shadow, reduced opacity

---

#### Cards & Containers

```tsx
// Standard Card
<div className="
  bg-white
  border-4 border-black
  shadow-brutal
  p-6
">
  <h3 className="font-display text-xl font-bold mb-2">Card Title</h3>
  <p className="font-body">Card content goes here.</p>
</div>

// Colored Card (Subject Card)
<div className="
  bg-accent-yellow
  border-4 border-black
  shadow-brutal
  p-6
  hover:shadow-brutal-lg
  hover:-translate-y-1
  transition-all duration-100
">
  <h3 className="font-display text-2xl font-bold text-black">Mathematics</h3>
  <p className="font-body text-black/80">150 Questions</p>
</div>

// Quiz Question Card
<div className="
  bg-cream
  border-4 border-black
  shadow-brutal-lg
  p-8
">
  <span className="
    inline-block px-3 py-1 mb-4
    bg-primary text-black
    font-display font-bold text-sm uppercase
    border-3 border-black
  ">
    Question 1 of 20
  </span>
  <h2 className="font-display text-2xl font-bold mb-6">
    What is the capital of Nigeria?
  </h2>
</div>
```

---

#### Input Fields

```tsx
// Text Input
<div className="mb-4">
  <label className="
    block font-display font-bold text-sm uppercase mb-2
  ">
    Email Address
  </label>
  <input
    type="email"
    className="
      w-full px-4 py-3
      bg-white
      border-4 border-black
      font-body text-lg
      shadow-brutal-sm
      focus:shadow-brutal focus:outline-none
      placeholder:text-gray
    "
    placeholder="you@example.com"
  />
</div>

// Input with Error
<div className="mb-4">
  <label className="block font-display font-bold text-sm uppercase mb-2">
    Password
  </label>
  <input
    type="password"
    className="
      w-full px-4 py-3
      bg-white
      border-4 border-accent-red
      font-body text-lg
      shadow-[4px_4px_0px_0px_#FF3366]
      focus:outline-none
    "
  />
  <p className="mt-2 font-body text-sm text-accent-red font-bold">
    Password must be at least 6 characters
  </p>
</div>

// Select Dropdown
<select className="
  w-full px-4 py-3
  bg-white
  border-4 border-black
  font-body text-lg
  shadow-brutal-sm
  focus:shadow-brutal focus:outline-none
  appearance-none
  cursor-pointer
">
  <option>Select Subject</option>
  <option>Mathematics</option>
  <option>English</option>
</select>
```

---

#### Quiz Option Buttons

```tsx
// Unselected Option
<button className="
  w-full p-4 mb-3
  bg-white text-left
  border-4 border-black
  shadow-brutal-sm
  font-body text-lg
  hover:bg-cream hover:shadow-brutal
  transition-all duration-100
">
  <span className="font-display font-bold mr-3">A.</span>
  Lagos
</button>

// Selected Option
<button className="
  w-full p-4 mb-3
  bg-primary text-left
  border-4 border-black
  shadow-brutal
  font-body text-lg font-bold
">
  <span className="font-display font-bold mr-3">B.</span>
  Abuja ✓
</button>

// Correct Answer (After Submit)
<button className="
  w-full p-4 mb-3
  bg-accent-green text-left
  border-4 border-black
  shadow-brutal
  font-body text-lg font-bold
">
  <span className="font-display font-bold mr-3">B.</span>
  Abuja ✓
</button>

// Wrong Answer (After Submit)
<button className="
  w-full p-4 mb-3
  bg-accent-red text-white text-left
  border-4 border-black
  shadow-brutal
  font-body text-lg font-bold
">
  <span className="font-display font-bold mr-3">A.</span>
  Lagos ✗
</button>
```

---

#### Header / Navigation

```tsx
<header className="
  bg-black text-white
  border-b-4 border-primary
  py-4 px-6
">
  <div className="max-w-7xl mx-auto flex items-center justify-between">
    {/* Logo */}
    <a href="/" className="font-display text-2xl md:text-3xl font-bold uppercase tracking-tight">
      Score<span className="text-primary">That</span>Exam
    </a>
    
    {/* Desktop Nav */}
    <nav className="hidden md:flex items-center gap-6">
      <a href="/dashboard" className="
        font-display font-bold uppercase
        hover:text-primary transition-colors
      ">
        Dashboard
      </a>
      <a href="/quiz" className="
        font-display font-bold uppercase
        hover:text-primary transition-colors
      ">
        Quiz
      </a>
      <a href="/leaderboard" className="
        font-display font-bold uppercase
        hover:text-primary transition-colors
      ">
        Leaderboard
      </a>
      <button className="
        px-4 py-2
        bg-primary text-black
        font-display font-bold uppercase
        border-3 border-white
        shadow-[3px_3px_0px_0px_#FFFFFF]
        hover:bg-accent-yellow
        active:shadow-none active:translate-x-0.5 active:translate-y-0.5
      ">
        Logout
      </button>
    </nav>
    
    {/* Mobile Menu Button */}
    <button className="md:hidden p-2 border-3 border-white">
      <MenuIcon className="w-6 h-6" />
    </button>
  </div>
</header>
```

---

#### Footer

```tsx
<footer className="
  bg-secondary text-white
  border-t-4 border-black
  py-8 px-6
">
  <div className="max-w-7xl mx-auto">
    <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
      {/* Brand */}
      <div>
        <h3 className="font-display text-2xl font-bold uppercase mb-4">
          Score<span className="text-primary">That</span>Exam
        </h3>
        <p className="font-body text-white/80">
          Ace your JAMB exams with confidence.
        </p>
      </div>
      
      {/* Links */}
      <div>
        <h4 className="font-display font-bold uppercase mb-4 text-primary">
          Quick Links
        </h4>
        <ul className="space-y-2 font-body">
          <li><a href="/about" className="hover:text-primary">About Us</a></li>
          <li><a href="/contact" className="hover:text-primary">Contact</a></li>
          <li><a href="/privacy" className="hover:text-primary">Privacy Policy</a></li>
        </ul>
      </div>
      
      {/* Subjects */}
      <div>
        <h4 className="font-display font-bold uppercase mb-4 text-primary">
          Subjects
        </h4>
        <ul className="space-y-2 font-body">
          <li><a href="/quiz?subject=1" className="hover:text-primary">Mathematics</a></li>
          <li><a href="/quiz?subject=2" className="hover:text-primary">English</a></li>
          <li><a href="/quiz?subject=3" className="hover:text-primary">Physics</a></li>
        </ul>
      </div>
    </div>
    
    <div className="mt-8 pt-6 border-t-2 border-white/20 text-center font-display font-bold uppercase">
      © 2026 AceThatPaper. All rights reserved.
    </div>
  </div>
</footer>
```

---

#### Leaderboard Table

```tsx
<div className="border-4 border-black shadow-brutal-lg overflow-hidden">
  {/* Table Header */}
  <div className="bg-secondary text-white grid grid-cols-12 gap-4 p-4 font-display font-bold uppercase">
    <div className="col-span-1">Rank</div>
    <div className="col-span-5">Name</div>
    <div className="col-span-3">Score</div>
    <div className="col-span-3">Accuracy</div>
  </div>
  
  {/* Row - Highlighted (Current User) */}
  <div className="bg-accent-yellow grid grid-cols-12 gap-4 p-4 border-t-4 border-black">
    <div className="col-span-1 font-display font-bold text-xl">#1</div>
    <div className="col-span-5 font-body font-bold">You 🏆</div>
    <div className="col-span-3 font-display font-bold">2,450</div>
    <div className="col-span-3 font-display font-bold text-accent-green">94%</div>
  </div>
  
  {/* Row - Normal */}
  <div className="bg-white grid grid-cols-12 gap-4 p-4 border-t-3 border-black">
    <div className="col-span-1 font-display font-bold text-xl">#2</div>
    <div className="col-span-5 font-body">John Doe</div>
    <div className="col-span-3 font-display font-bold">2,380</div>
    <div className="col-span-3 font-display">91%</div>
  </div>
</div>
```

---

#### Alerts & Notifications

```tsx
// Success Alert
<div className="
  bg-accent-green text-black
  border-4 border-black
  shadow-brutal
  p-4
  flex items-center gap-3
">
  <CheckCircle className="w-6 h-6" />
  <span className="font-body font-bold">Quiz submitted successfully!</span>
</div>

// Error Alert
<div className="
  bg-accent-red text-white
  border-4 border-black
  shadow-brutal
  p-4
  flex items-center gap-3
">
  <XCircle className="w-6 h-6" />
  <span className="font-body font-bold">Invalid email or password.</span>
</div>

// Warning Alert
<div className="
  bg-accent-yellow text-black
  border-4 border-black
  shadow-brutal
  p-4
  flex items-center gap-3
">
  <AlertTriangle className="w-6 h-6" />
  <span className="font-body font-bold">Your session will expire in 5 minutes.</span>
</div>
```

---

#### Loading States

```tsx
// Skeleton Card
<div className="
  bg-gray-light
  border-4 border-black
  shadow-brutal
  p-6
  animate-pulse
">
  <div className="h-6 bg-gray w-3/4 mb-4"></div>
  <div className="h-4 bg-gray w-1/2"></div>
</div>

// Loading Spinner
<div className="flex items-center justify-center p-8">
  <div className="
    w-12 h-12
    border-4 border-black border-t-primary
    animate-spin
  "></div>
</div>

// Loading Button
<button className="
  px-6 py-3
  bg-gray-light text-gray
  font-display font-bold text-lg uppercase
  border-4 border-black
  cursor-not-allowed
" disabled>
  <span className="animate-pulse">Loading...</span>
</button>
```

---

### Responsive Breakpoints

```javascript
// TailwindCSS Default Breakpoints
// sm: 640px   - Mobile landscape
// md: 768px   - Tablets
// lg: 1024px  - Small laptops
// xl: 1280px  - Desktops
// 2xl: 1536px - Large screens
```

**Mobile-First Approach:**
```tsx
// Example: Responsive Grid
<div className="
  grid
  grid-cols-1        /* Mobile: 1 column */
  sm:grid-cols-2     /* Tablet: 2 columns */
  lg:grid-cols-3     /* Desktop: 3 columns */
  xl:grid-cols-4     /* Large: 4 columns */
  gap-4 sm:gap-6
">
```

**Responsive Typography:**
```tsx
<h1 className="
  text-2xl          /* Mobile: 24px */
  sm:text-3xl       /* Tablet: 30px */
  md:text-4xl       /* Desktop: 36px */
  lg:text-5xl       /* Large: 48px */
  font-display font-bold
">
```

**Responsive Spacing:**
```tsx
<section className="
  px-4 py-8          /* Mobile */
  sm:px-6 sm:py-12   /* Tablet */
  lg:px-8 lg:py-16   /* Desktop */
">
```

---

## Project Structure

```
frontend/
├── public/
│   ├── favicon.ico
│   └── og-image.png
├── src/
│   ├── assets/                 # Static assets (images, fonts)
│   ├── components/             # Reusable UI components
│   │   ├── ui/                 # Base UI components
│   │   │   ├── Button.tsx
│   │   │   ├── Card.tsx
│   │   │   ├── Input.tsx
│   │   │   ├── Alert.tsx
│   │   │   └── ...
│   │   ├── layout/             # Layout components
│   │   │   ├── Header.tsx
│   │   │   ├── Footer.tsx
│   │   │   ├── Sidebar.tsx
│   │   │   └── Layout.tsx
│   │   ├── quiz/               # Quiz-specific components
│   │   │   ├── QuestionCard.tsx
│   │   │   ├── OptionButton.tsx
│   │   │   ├── QuizProgress.tsx
│   │   │   └── QuizResult.tsx
│   │   └── leaderboard/        # Leaderboard components
│   │       ├── LeaderboardTable.tsx
│   │       └── RankBadge.tsx
│   ├── pages/                  # Page components (routes)
│   │   ├── Home.tsx
│   │   ├── Login.tsx
│   │   ├── Register.tsx
│   │   ├── Dashboard.tsx
│   │   ├── Quiz.tsx
│   │   ├── QuizResult.tsx
│   │   ├── Leaderboard.tsx
│   │   ├── Profile.tsx
│   │   └── NotFound.tsx
│   ├── hooks/                  # Custom React hooks
│   │   ├── useAuth.ts
│   │   ├── useQuiz.ts
│   │   └── useLeaderboard.ts
│   ├── services/               # API service layer
│   │   ├── api.ts              # Axios instance
│   │   ├── auth.service.ts
│   │   ├── quiz.service.ts
│   │   ├── user.service.ts
│   │   └── leaderboard.service.ts
│   ├── context/                # React Context providers
│   │   ├── AuthContext.tsx
│   │   └── QuizContext.tsx
│   ├── types/                  # TypeScript type definitions
│   │   ├── auth.types.ts
│   │   ├── quiz.types.ts
│   │   ├── user.types.ts
│   │   └── api.types.ts
│   ├── utils/                  # Utility functions
│   │   ├── formatters.ts
│   │   ├── validators.ts
│   │   └── storage.ts
│   ├── styles/                 # Global styles
│   │   └── index.css
│   ├── App.tsx
│   ├── main.tsx
│   └── vite-env.d.ts
├── .env
├── .env.example
├── index.html
├── package.json
├── tailwind.config.js
├── tsconfig.json
└── vite.config.ts
```

---

## API Integration

### Base API Configuration

```typescript
// src/services/api.ts
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor - Add auth token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Response interceptor - Handle token refresh
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      
      try {
        const refreshToken = localStorage.getItem('refresh_token');
        const { data } = await axios.post(`${API_URL}/auth/refresh`, {
          refresh_token: refreshToken,
        });
        
        localStorage.setItem('access_token', data.data.access_token);
        localStorage.setItem('refresh_token', data.data.refresh_token);
        
        originalRequest.headers.Authorization = `Bearer ${data.data.access_token}`;
        return api(originalRequest);
      } catch (refreshError) {
        localStorage.removeItem('access_token');
        localStorage.removeItem('refresh_token');
        window.location.href = '/login';
        return Promise.reject(refreshError);
      }
    }
    
    return Promise.reject(error);
  }
);
```

### API Endpoints Reference

| Endpoint | Method | Auth Required | Description |
|----------|--------|---------------|-------------|
| `/user/register` | POST | No | Register new user |
| `/user/login` | POST | No | User login |
| `/admin/login` | POST | No | Admin login |
| `/auth/refresh` | POST | No | Refresh tokens |
| `/api/v1/dashboard` | GET | Yes | Get user dashboard |
| `/api/v1/quiz/create` | POST | Yes | Generate quiz |
| `/api/v1/quiz/submit` | POST | Yes | Submit quiz |
| `/api/v1/leaderboard` | GET | Yes | Get leaderboard |
| `/api/v1/leaderboard/me` | GET | Yes | Get user's rank |
| `/api/v1/admin/subject` | GET | Yes | Get all subjects |

---

## Pages & Routes

```tsx
// src/App.tsx
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import ProtectedRoute from './components/ProtectedRoute';

// Pages
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Quiz from './pages/Quiz';
import QuizResult from './pages/QuizResult';
import Leaderboard from './pages/Leaderboard';
import Profile from './pages/Profile';
import NotFound from './pages/NotFound';

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          {/* Public Routes */}
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          
          {/* Protected Routes */}
          <Route element={<ProtectedRoute />}>
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/quiz" element={<Quiz />} />
            <Route path="/quiz/result" element={<QuizResult />} />
            <Route path="/leaderboard" element={<Leaderboard />} />
            <Route path="/profile" element={<Profile />} />
          </Route>
          
          {/* 404 */}
          <Route path="*" element={<NotFound />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  );
}
```

---

## Key User Flows

### 1. Authentication Flow
```
Landing Page → Login/Register → Dashboard
                     ↓
              JWT Token Storage
                     ↓
              Auto Token Refresh
```

### 2. Quiz Flow
```
Dashboard → Select Subject → Configure Quiz (# questions)
                                    ↓
                            Generate Quiz API
                                    ↓
                            Quiz Page (Question by Question)
                                    ↓
                            Submit Answers
                                    ↓
                            Results Page (Score + Review)
```

### 3. Leaderboard Flow
```
Dashboard/Nav → Leaderboard Page
                      ↓
              Filter by Subject/Period
                      ↓
              View Rankings + Own Position
```

---

## Accessibility Requirements

1. **Keyboard Navigation**: All interactive elements must be keyboard accessible
2. **Focus States**: Visible focus indicators on all focusable elements
3. **ARIA Labels**: Proper labels for screen readers
4. **Color Contrast**: Minimum 4.5:1 for text, 3:1 for large text
5. **Error Announcements**: Form errors announced to screen readers
6. **Skip Links**: Skip to main content link for keyboard users

```tsx
// Focus ring for brutal design
.focus-brutal {
  @apply focus:outline-none focus:ring-4 focus:ring-black focus:ring-offset-2;
}
```

---

## Environment Variables

```env
# .env.example
VITE_API_URL=http://localhost:8080
VITE_APP_NAME=AceThatPaper
```

---

## Performance Guidelines

1. **Lazy Loading**: Use `React.lazy()` for route-based code splitting
2. **Image Optimization**: Use WebP format, proper sizing
3. **Memoization**: Use `useMemo` and `useCallback` appropriately
4. **Bundle Analysis**: Keep initial bundle under 200KB
5. **API Caching**: Cache static data (subjects, etc.)

---

## Quick Reference for AI Agents

When working on this codebase:
- **Entry point**: `src/main.tsx`
- **Routes defined**: `src/App.tsx`
- **Add new components**: `src/components/`
- **Add new pages**: `src/pages/`
- **API services**: `src/services/`
- **Type definitions**: `src/types/`
- **Global styles**: `src/styles/index.css`
- **Config files**: `tailwind.config.js`, `vite.config.ts`

### Design Checklist for Every Component:
- [ ] Uses thick black borders (3-4px)
- [ ] Has brutal drop shadow
- [ ] Uses design system colors
- [ ] Uses Space Grotesk font
- [ ] Has hover/active states with hard transitions
- [ ] Is responsive (mobile-first)
- [ ] Is keyboard accessible
- [ ] Has proper TypeScript types

