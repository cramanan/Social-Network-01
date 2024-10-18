// src/app.tsx
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./page";
import Profile from "./profile/page";

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={
            <HomeProfileLayout>
              <Home />
              <Profile />
            </HomeProfileLayout>
          }
        />
        <Route
          path="/profile"
          element={
            <HomeProfileLayout>
              <Profile />
            </HomeProfileLayout>
          }
        />
      </Routes>
      <Route
        path="/"
        element={
          <HomeProfileLayout>
            <Home />
          </HomeProfileLayout>
        }
      />
      <Route
        path="/profile"
        element={
          <HomeProfileLayout>
            <Profile />
          </HomeProfileLayout>
        }
      />
    </Router>
  );
};

export default App;
