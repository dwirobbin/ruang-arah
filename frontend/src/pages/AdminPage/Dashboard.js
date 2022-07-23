import React, { useState, useEffect } from "react";
import axios from "axios";

import API from "../../api/Api";
import Navbar from "../../components/Navbar/Navbar";
import Footer from "../../components/Footer/Footer";
import Dashboard from "../../components/Admin/Dashboard/Dashboard";

export default function Admin() {
  let auth = localStorage.getItem("token");

  const [questions, setQuestions] = useState([]);

  const fetchQuestions = async () => {
    const { data: res } = await axios.get(
      `${API.API_URL}/api/admin/questions`,
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          Authorization: "Bearer " + auth,
        },
      }
    );
    setQuestions(res.data);
  };

  useEffect(() => {
    fetchQuestions();
  }, []);

  return (
    <>
      <Navbar />
      <Dashboard
        auth={auth}
        questions={questions}
        fetchQuestions={fetchQuestions}
      />
      <Footer />
    </>
  );
}
