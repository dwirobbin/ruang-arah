import React, { useEffect, useContext } from "react";
import { useNavigate, Outlet } from "react-router-dom";
import AuthContext from "../store/AuthContext";

function PrivateRoute() {
  const user = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!user.isLoggedIn && !token) {
      navigate("/login");
    }
    if (user.role === "admin") {
      navigate("/dashboard");
    }
  }, []);

  return (
    <div>
      <Outlet />
    </div>
  );
}

export default PrivateRoute;
