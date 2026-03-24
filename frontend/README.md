# AceThatPaper (Frontend) 🚀
AceThatPaper is a modern, high-performance web application designed to help Nigerian students excel in their JAMB exams. It features a rapidly scalable exam simulation engine combined with gamified, viral social elements.

## 🌟 Key Features Implemented

### 1. The Core Exam Engine
- **Subject & Quiz Selection**: Users can choose any available subject and dynamically select the number of questions (up to 60) per session.
- **Real-Time Scoring System**: Accuracy is dynamically calculated. Rather than showing a raw correct count, Total Scores are intelligently scaled out of 100 on completion to match familiar grading structures.
- **Explanations:** Students don't just see what they missed; they get detailed explanations for every single answer immediately after testing.

### 2. Gamification & Leaderboard
- **Global & Periodic Rankings**: The highly competitive leaderboard supports `All Time`, `This Week`, and `This Month` filters, allowing students to compete on short-term milestones.
- **Intelligent User Rank Tracking**: The backend `userRank` is smoothly integrated on the frontend so the user can see their exact rank contextualized securely within the currently active time-filter.

### 3. Progressive Web App (PWA) 📱
- **Zero-Friction Installs**: Using `vite-plugin-pwa`, the site is fully installable on Android and iOS devices directly from the browser (bypassing the App Store entirely).
- **Offline Reliability**: The service worker caches static assets to ensure fast load times across unpredictable internet connections.

### 4. Viral Loops & Monetization 💸
- **WhatsApp Share Gateway**: Upon scoring >= 50% in a quiz session, students are prompted to automatically broadcast a brag/challenge to their WhatsApp statuses, artificially injecting viral growth.
- **Facebook Community Integration**: Strategic "Join the Group" Call-to-Actions funnel isolated users into the massive 1.2M member Facebook study group.
- **AdSense Ready**: Base integration for Google Auto Ads is implemented in `<head>` to begin monetizing the spike in Facebook organic traffic.

### 5. UI/UX: Neobrutalism
- Designed using TailwindCSS to implement a **Neobrutalist** aesthetic (thick black borders, harsh shadows, stark high-contrast colors). It feels deliberately modern, gamified, and distinct from traditional "boring" academic apps.

## 🛠 Tech Stack
- **Framework**: React.js 18 + Vite
- **Routing**: React Router
- **Global State**: React Context API (`AuthContext`)
- **Styling**: Tailwind CSS v3
- **Icons**: `lucide-react` & Custom SVGs
- **Deployment**: Vercel (Configured via `vercel.json` for SPA rewrites)
- **Backend API**: Echo (Golang)

## 📦 Getting Started Locally

```bash
# 1. Install dependencies
npm install

# 2. Setup your Environment Variables
cp .env.example .env
# Edit .env and enter your VITE_API_URL

# 3. Start the dev server
npm run dev
```
