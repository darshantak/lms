// import React, { Component } from "react";
import { Navigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

const ProtectedRoute = ({ element: Component, ...rest }) => {
  const token = localStorage.token;
  console.log("token",token)
  const isTokenValid = () => {
    if (!token) return false;
    try {
      const decodeToken = jwtDecode(token);
      const currentTime = Date.now() / 1000;
      console.log(decodeToken.exp)
      return decodeToken.exp > currentTime;
    } catch (error) {
      console.error("Token validation error:", error,);
      return false;
    }
  };
  return isTokenValid() ? Component : <Navigate to="/" />;
};

export default ProtectedRoute;
