# BrainGrid Backend

BrainGrid is a SaaS learning platform backend built with Go, Gin, GORM, PostgreSQL, and JWT authentication. This repository currently contains the backend foundation for user authentication and protected API access, and it is structured so the product can grow into a full study and productivity platform.

## What This Repository Does Today

The current backend supports:

- User registration with hashed passwords
- User login with JWT token generation
- Protected API routes guarded by JWT middleware
- PostgreSQL connection and automatic `User` model migration

### Current API Routes

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/user/profile`

## Technology Stack

- Go
- Gin web framework
- GORM ORM
- PostgreSQL
- JWT authentication
- bcrypt password hashing
- dotenv-based environment loading

## Project Structure

- `cmd/` - application entry point
- `config/` - database connection setup
- `controllers/` - request handlers
- `dto/` - request payload models
- `middleware/` - auth middleware
- `models/` - database entities
- `routes/` - route registration
- `services/` - business logic and data access
- `utils/` - JWT and password helpers

## Environment Variables

Create a `.env` file in the project root with these values:

- `PORT`
- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `JWT_SECRET`

Example:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=braingrid
JWT_SECRET=super_secret_key
```

## Local Setup

1. Install PostgreSQL and create a database for BrainGrid.
2. Set up the `.env` file with the values above.
3. Install dependencies with `go mod tidy`.
4. Run the backend with `go run ./cmd`.

The server starts on the port defined in `PORT`, or `8080` if the variable is not set.

## BrainGrid Product Vision

The codebase is currently a backend starter, but the target product is a modern AI-powered learning SaaS. The following feature set is the intended direction for BrainGrid.

### AI-Powered Study Tools

- Smart lesson summarization
- Topic explanation and Q&A assistant
- Flashcard generation from notes and documents
- Quiz and exam question generation
- Personalized study recommendations

### Smart Notes System

- Rich note creation and editing
- Tagging, folders, and smart search
- Auto-linking related concepts
- Note-to-summary and note-to-quiz workflows
- AI-assisted note cleanup and organization

### Productivity Management

- Task and goal tracking
- Study session scheduling
- Pomodoro and focus timers
- Deadline reminders and daily planning
- Progress dashboards

### Study Analytics

- Study streaks
- Session duration tracking
- Weak-topic detection
- Knowledge progress over time
- Goal completion analytics

### Collaboration Features

- Shared study spaces
- Group notes and study boards
- Team task assignment
- Comments and mentions
- Peer review of notes and answers

### Gamification

- XP and leveling
- Badges and achievements
- Streak rewards
- Challenge-based learning
- Leaderboards for groups and cohorts

### Admin Panel

- User management
- Content moderation
- Feature flags and access control
- Subscription and plan management
- Analytics and usage reporting

### Mobile App Features

- On-the-go note access
- Voice-to-note capture
- Quick revision cards
- Offline study mode
- Push notifications for study reminders

## Backend Modules To Grow Into

### Authentication and Identity

- Sign up, login, logout
- Password reset and email verification
- Role-based access control
- Session and token management

### Study Content

- Notes
- Flashcards
- Quizzes
- Study plans
- AI-generated summaries

### Productivity

- Tasks
- Goals
- Events
- Reminders
- Time tracking

### Collaboration

- Teams
- Shared boards
- Comments
- Invites
- Notifications

### Analytics

- Study sessions
- Progress metrics
- Performance trends
- Engagement tracking

### Administration

- Users
- Plans
- Reports
- Moderation tools

## Frontend Modules To Pair With This Backend

- Authentication screens
- Student dashboard
- Notes workspace
- Study planner
- Analytics dashboard
- Collaboration space
- Admin console
- Mobile companion app

## Database Table Ideas

- `users`
- `roles`
- `sessions`
- `notes`
- `note_tags`
- `flashcards`
- `quizzes`
- `tasks`
- `goals`
- `study_sessions`
- `study_metrics`
- `groups`
- `group_members`
- `comments`
- `notifications`
- `achievements`
- `subscriptions`

## MVP Roadmap

### Phase 1: Core Platform

- Authentication
- User profile
- Notes CRUD
- Basic study sessions
- Protected API structure

### Phase 2: Learning Tools

- Flashcards
- Quiz generation
- Smart summaries
- Search and tagging

### Phase 3: Productivity Layer

- Task tracking
- Goals
- Timers
- Reminders
- Basic analytics

### Phase 4: Collaboration and Growth

- Shared spaces
- Comments
- Admin tools
- Subscription support
- Mobile app support

## Future Advanced AI Features

- Personalized tutoring chat
- Adaptive learning paths
- AI study coach
- Context-aware recommendations
- Knowledge graph suggestions
- Automatic weak-area detection
- Smart revision planning

## Notes

This repository is currently a backend foundation. The README includes the intended BrainGrid product direction so the codebase documentation matches the broader feature vision and can evolve alongside future frontend and AI work.
