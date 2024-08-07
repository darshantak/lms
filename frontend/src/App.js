import LoginPage from "./screens/LoginPage";
import StudentDashboard from "./screens/StudentDashboard";
import AdminDashboard from "./screens/AdminDashboard";
import ProtectedRoute from "./screens/ProtectedRoute";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { ErrorBoundary } from "react-error-boundary";
function App() {
  return (
    <ErrorBoundary fallback={<h1>Something went wrong</h1>}>
      <Router>
        <Routes>
          <Route path="/" element={<LoginPage />} />
          <Route
            path="/admin-dashboard"
            element={<ProtectedRoute element={<AdminDashboard />} />}
          />
          <Route
            path="/student-dashboard"
            element={<ProtectedRoute element={<StudentDashboard />} />}
          />
        </Routes>
      </Router>
    </ErrorBoundary>
  );
}

export default App;
