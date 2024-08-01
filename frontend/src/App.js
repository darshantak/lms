import LoginPage from "./components/LoginPage";
import Dashboard from "./components/Dashboard";
import ProtectedRoute from "./components/ProtectedRoute";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
function App() {
  return (
    // <div className="App">
      <Router>
        <Routes> 
          <Route path="/" element={<LoginPage/>} />
          <Route path="/dashboard" element={<ProtectedRoute element={<Dashboard />}/>} />
       </Routes>
        {/* <LoginPage /> */}
      </Router>
    // </div>
  );
}

export default App;
