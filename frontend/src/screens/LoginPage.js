import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "react-toastify/dist/ReactToastify.css";
import Toast from "../components/Toast";
import BounceLoader from "react-spinners/BounceLoader";
import jwt_decode, { jwtDecode } from "jwt-decode";
const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);
  const [role, setRole] = useState("");
  const navigate = useNavigate();
  const handleLogin = async (event) => {
    // Handle login logic here
    console.log("Email:", email);
    console.log("Password:", password);
    event.preventDefault();
    try {
      const response = await fetch("/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        throw new Error("Network response not ok");
      }

      const data = await response.json();
      console.log("Success:", data);
      localStorage.setItem("token", data.token);
      const decoded = jwtDecode(data.token);
      const userRole = decoded.role;

      if (userRole=="ADMIN"){
      navigate("/admin-dashboard");}
      else{
        navigate("/student-dashboard")
      }
    } catch (error) {
      console.error("Error:", error);
      setError("Login failed. Please check your email or password.");
    }
  };

  return (
    <div
      className="min-h-screen flex items-center justify-center"
      style={{
        backgroundImage: `url(${process.env.PUBLIC_URL + "/background.jpg"})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      <div className="bg-gray-300 shadow-md rounded-lg px-8 py-6 w-full max-w-xl mx-4">
        <h1 className="text-2xl font-bold text-center mb-4 dark:text-gray-800">
          Welcome Back!
        </h1>
        <form onSubmit={handleLogin}>
          <div className="mb-4">
            <label
              htmlFor="emailLabel"
              className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2"
            >
              Email Address
            </label>
            <input
              type="email"
              id="email"
              className="shadow-sm rounded-md w-full px-3 py-2 border border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              placeholder="your@email.com"
              onChange={(event) => setEmail(event.target.value)}
              required
            />
          </div>
          <div className="mb-4">
            <label
              htmlFor="passwordLabel"
              className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2"
            >
              Password
            </label>
            <input
              type="password"
              id="password"
              className="shadow-sm rounded-md w-full px-3 py-2 border border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              placeholder="Enter your password"
              onChange={(event) => setPassword(event.target.value)}
              required
            />
          </div>
          <div className="flex items-center justify-between mb-4">
            <a
              href="#"
              className="text-xs text-indigo-500 hover:text-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Create Account
            </a>
            <a
              href="#"
              className="text-xs text-indigo-500 hover:text-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Forgot Password
            </a>
          </div>
          <button
            type="submit"
            className="w-full flex justify-center py-2 px-4 btn btn-neutral"
          >
            Login
          </button>

          {error && <div className="text-red-500 text-sm mt-4">{error}</div>}
        </form>
        {/* <div className="spinner-container">
          <BounceLoader color="#36D7B7" loading={true} size={150} />
        </div> */}
      </div>
    </div>
  );
};

export default LoginPage;
