import React from "react";
import { useNavigate } from "react-router-dom";
import "./style.css";
import Banner from "../../assets/images/Level-image.png";

import Card from "react-bootstrap/Card";
import Image from "react-bootstrap/Image";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Footer from "../../components/Footer/Footer";
import Navbar from "../../components/Navbar/Navbar";

function Result() {
  const navigate = useNavigate();
  const level = JSON.parse(localStorage.getItem("Level"));

  const handleRecommendation = () => {
    navigate({
      pathname: "/quiz/result/recommendation",
      search: `?level_id=${level.level_id}`,
    });
  };

  return (
    <>
      <Navbar />
      <section id="level-pages">
        <Container>
          <Row>
            <Col sm={3}>
              <div className="level-image pt-5">
                <Image src={Banner} fluid />
              </div>
            </Col>
            <Col sm={8}>
              <div className="level-body pt-2">
                <h1>
                  Congratulation! <span> Your Level Is</span>{" "}
                </h1>
              </div>
              <div className="card-rank">
                <Card>
                  <Card.Body>
                    <Card.Text>{level.level}</Card.Text>
                  </Card.Body>
                </Card>
              </div>
              <div className="recommendation">
                <button onClick={handleRecommendation}>
                  SEE RECOMMENDATION
                </button>
              </div>
            </Col>
          </Row>
        </Container>
      </section>
      <Footer />
    </>
  );
}

export default Result;
