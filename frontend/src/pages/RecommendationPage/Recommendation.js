import React, { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import "./style.css";
import Navbar from "../../components/Navbar/Navbar";
import Footer from "../../components/Footer/Footer";
import axios from "axios";
import API from "../../api/Api";

import { useNavigate } from "react-router-dom";
import { Container, Button } from "react-bootstrap";

const Recommendation = () => {
  let query = new URLSearchParams(useLocation().search);
  let levelId = parseInt(query.get("level_id"));

  const navigate = useNavigate();

  const [recommendation, setRecommendation] = useState({});

  const getRecommendation = async () => {
    let auth = localStorage.getItem("token");

    try {
      let { data: res } = await axios.get(
        `${API.API_URL}/api/home/recommendation?level_id=${levelId}`,
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
            Authorization: "Bearer " + auth,
          },
        }
      );

      setRecommendation(res.data.image_url);
    } catch (err) {}
  };

  useEffect(() => {
    getRecommendation();
  }, []);

  const handleDone = () => {
    alert("Congratulations! You have finished the test");
    navigate("/home");
  };
  return (
    <section id="recomendation-page">
      <Navbar />
      <h1 className="mb-5 text-center title">Recomendation Page</h1>
      <Container className="container-luar">
        <img src={recommendation} alt="recomendation" className="images" />
        <br />

        <Button variant="success" className="btn-done" onClick={handleDone}>
          SELESAI
        </Button>
      </Container>
      <Footer />
    </section>
  );
};

export default Recommendation;
