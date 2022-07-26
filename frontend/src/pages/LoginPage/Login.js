import React, { useRef, useContext } from "react";
import { useNavigate } from "react-router-dom";
import AuthContext from "../../store/AuthContext";
import API from "../../api/Api";

import "./style.css";
import Banner from "../../assets/images/Login-Images.png";

export const Login = () => {
  const navigate = useNavigate();
  const emailInputRef = useRef();
  const passwordInputRef = useRef();
  const authCtx = useContext(AuthContext);

  if (authCtx.role === "admin") {
    navigate("/dashboard");
  }
  if (authCtx.isLoggedIn) {
    navigate("/home");
  }

  const handleSubmit = (event) => {
    event.preventDefault();

    const enteredEmail = emailInputRef.current.value;
    const enteredPassword = passwordInputRef.current.value;
    let auth = localStorage.getItem("token");

    let url;
    url = `${API.API_URL}/api/auth/login`;
    fetch(url, {
      method: "POST",
      body: JSON.stringify({
        email: enteredEmail,
        password: enteredPassword,
        returnSecureToken: true,
      }),
      headers: {
        Accept: "/",
        "Content-Type": "application/json",
        Authorization: "Bearer " + auth,
      },
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else {
          return res.json().then((data) => {
            let errorMassage = "Authentication failed!";
            throw new Error(errorMassage);
          });
        }
      })
      .then((data) => {
        authCtx.login(data.data.token, data.data.role);
      })
      .catch((err) => {
        alert("Data yang anda masukkan salah, silahkan di cek!");
      });
  };

  return (
    <section id="login-pages">
      <div className="row">
        <div className="login-left col-lg-6">
          <img className="w-75 img-fluid mx-auto" src={Banner} alt="banner" />
        </div>
        <div className="login-right col-lg-6 my-auto">
          <p>Welcome Back</p>
          <h2 className="mb-5">Login to your Account</h2>
          <form onSubmit={handleSubmit}>
            <div className="mb-3">
              <label htmlFor="emailInput" className="form-label">
                Email
              </label>
              <input
                type="email"
                id="email"
                className="form-control"
                ref={emailInputRef}
              />
            </div>
            <div className="mb-3">
              <label htmlFor="exampleInputPassword1" className="form-label">
                Password
              </label>
              <input
                type="password"
                id="password"
                className="form-control"
                ref={passwordInputRef}
              />
            </div>
            <button type="submit" className="btn mt-3">
              Login
            </button>
            <p className="make-account mt-4">
              Doesn't have an account ?&nbsp;
              <a href="/register" className="text-decoration-none">
                SignUp
              </a>
            </p>
          </form>
        </div>
      </div>
    </section>
  );
};

export default Login;
