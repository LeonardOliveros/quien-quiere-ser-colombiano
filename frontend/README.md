# Quiz Application - Vue.js 3 Frontend

This is the Vue.js 3 frontend for the Colombian Quiz Application, migrated from vanilla JavaScript.

## Technology Stack

- **Vue.js 3** - Progressive JavaScript framework
- **TypeScript** - Type-safe JavaScript
- **Vite** - Fast build tool and dev server
- **Vue Router** - Official router for Vue.js
- **Pinia** - State management library
- **Axios** - HTTP client
- **Bootstrap 5** - CSS framework

## Project Structure

```
frontend/
├── src/
│   ├── components/      # Reusable Vue components (modals)
│   ├── views/          # Page components (Login, Menu, Game, Results)
│   ├── stores/         # Pinia stores (auth, game)
│   ├── services/       # API service layer
│   ├── types/          # TypeScript type definitions
│   ├── router/         # Vue Router configuration
│   ├── style.css       # Global styles
│   ├── App.vue         # Root component
│   └── main.ts         # Application entry point
├── index.html          # HTML template
├── vite.config.ts      # Vite configuration
├── tsconfig.json       # TypeScript configuration
└── package.json        # Dependencies and scripts
```

## Development

### Install Dependencies

```bash
cd frontend
npm install
```

### Run Development Server

```bash
npm run dev
```

This will start the Vite dev server on http://localhost:5173

The dev server includes:
- Hot Module Replacement (HMR)
- Proxy to backend API on http://localhost:8080

### Build for Production

```bash
npm run build
```

This will:
1. Type-check the code with `vue-tsc`
2. Build the application for production
3. Output to `../dist` directory

### Preview Production Build

```bash
npm run preview
```

## Features

### Authentication
- Login and registration forms
- JWT token-based authentication
- Automatic session management with localStorage

### Game Modes
- **Practice Mode**: Unlimited time to answer questions
- **Timed Mode**: 3-hour countdown for all questions
- **Weak Areas**: Focus on categories with lower performance
- **Category Mode**: Select specific category to practice

### Game Features
- Question display with 4 multiple-choice answers
- Three lifelines:
  - 50:50 (remove two incorrect answers)
  - Hint (show question hint)
  - Skip (skip current question, 3 uses)
- Flag questions for review
- Progress tracking
- Real-time timer for timed mode

### Results
- Overall score with pass/fail indication
- Category-wise performance breakdown
- Review of incorrect answers with explanations
- Flagged questions list
- Personalized study recommendations

### Statistics
- Overall game statistics
- Category performance tracking
- Weak and strong areas identification
- Game history
- Reset statistics option

## API Integration

The frontend communicates with the Go backend via REST API:

- **Authentication**: `/api/register`, `/api/login`
- **Game Management**: `/api/game/*`
- **Statistics**: `/api/user/:userId/stats`
- **Recommendations**: `/api/recommendations/:userId`

All protected routes require JWT authentication via `Authorization` header.

## State Management

### Auth Store
- User authentication state
- Login/logout functionality
- Token management

### Game Store
- Current game session state
- Question management
- Answer submission
- Lifeline usage
- Results loading

## Routing

- `/login` - Authentication page
- `/` - Main menu (protected)
- `/game` - Game screen (protected)
- `/results` - Results screen (protected)

## Styling

The application uses a "Millionaire Game Show" theme with:
- Gold and dark blue color scheme
- Animated transitions and effects
- Responsive design for mobile devices
- Custom styled components with scoped CSS

## Environment Variables

The Vite dev server proxies API requests to `http://localhost:8080` by default.

To change this, modify `vite.config.ts`:

```typescript
server: {
  proxy: {
    '/api': {
      target: 'http://your-backend-url',
      changeOrigin: true
    }
  }
}
```

## Browser Support

- Modern browsers with ES6+ support
- Chrome, Firefox, Safari, Edge (latest versions)
