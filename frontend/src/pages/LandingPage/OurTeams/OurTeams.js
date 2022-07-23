import React from "react";
import "./OurTeams.css";
import Footer from "../../../components/Footer/Footer";

// React BS
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Card from "react-bootstrap/Card";
import Container from "react-bootstrap/Container";

// Assets
import Team1 from "../../../assets/images/Team1.png";
import Team2 from "../../../assets/images/Team2.png";
import Team3 from "../../../assets/images/Team3.png";
import Team4 from "../../../assets/images/Team4.png";
import Team5 from "../../../assets/images/Team5.png";
import Team6 from "../../../assets/images/Team6.png";

export default function OurTeams() {
  return (
    <>
      <section id="OurTeams-page">
        <Container>
          <div className="text-heading">
            <h1>
              Our <span> Teams </span>{" "}
            </h1>
          </div>
          <Row className="main-content mt-5">
            <Col>
              <Card className="profile-detail">
                <Row>
                  <Col className="col-4 p-3 ps-3">
                    <img
                      src={Team1}
                      className="avatar img-fluid rounded-circle ms-3"
                      alt="logo"
                    />
                  </Col>
                  <Col className="col-8 pt-5 ps-3">
                    <p className="name pt-2">
                      Mochammad Andi Rambana
                      <br />
                      <span>
                        <span className="color">As</span> Frontend Engineer
                      </span>
                    </p>
                  </Col>
                </Row>
              </Card>
            </Col>
            <Col>
              <Card className="profile-detail">
                <Row>
                  <Col className="col-4 p-3 ps-3">
                    <img
                      src={Team2}
                      className="avatar img-fluid rounded-circle ms-3"
                      alt="logo"
                    />
                  </Col>
                  <Col className="col-8 pt-5 ps-3">
                    <p className="name pt-2">
                      Muhammad Rifqi Setiawan
                      <br />
                      <span>
                        <span className="color">As</span> Frontend Engineer
                      </span>
                    </p>
                  </Col>
                </Row>
              </Card>
            </Col>
          </Row>
          {/* 2 */}
          {/* <span></span> */}
          <Row className="main-content">
            <Col>
              <Card className="profile-detail">
                <Row>
                  <Col className="col-4 p-3 ps-3">
                    <img
                      src={Team3}
                      className="avatar img-fluid rounded-circle ms-3"
                      alt="logo"
                    />
                  </Col>
                  <Col className="col-8 pt-5 ps-3">
                    <p className="name pt-2">
                      Dwi Robbi Prasetyo
                      <br />
                      <span>
                        <span className="color">As</span> Backend Engineer
                      </span>
                    </p>
                  </Col>
                </Row>
              </Card>
            </Col>
            <Col>
              <Card className="profile-detail">
                <Row>
                  <Col className="col-4 p-3 ps-3">
                    <img
                      src={Team4}
                      className="avatar img-fluid rounded-circle ms-3"
                      alt="logo"
                    />
                  </Col>
                  <Col className="col-8 pt-5 ps-3">
                    <p className="name pt-2">
                      Navis Abdullah Farhan
                      <br />
                      <span>
                        <span className="color">As</span> Backend Engineer
                      </span>
                    </p>
                  </Col>
                </Row>
              </Card>
            </Col>
          </Row>
          {/* 3 */}
          <Row className="main-content">
            <Col>
              <Card className="profile-detail">
                <Row>
                  <Col className="col-4 p-3 ps-3">
                    <img
                      src={Team5}
                      className="avatar img-fluid rounded-circle ms-3"
                      alt="logo"
                    />
                  </Col>
                  <Col className="col-8 pt-5 ps-3">
                    <p className="name pt-2">
                      Paskahl herbert Simarmata
                      <br />
                      <span>
                        <span className="color">As</span> Backend Engineer
                      </span>
                    </p>
                  </Col>
                </Row>
              </Card>
            </Col>
            <Col>
              <Card className="profile-detail">
                <Row>
                  <Col className="col-4 p-3 ps-3">
                    <img
                      src={Team6}
                      className="avatar img-fluid rounded-circle ms-3"
                      alt="logo"
                    />
                  </Col>
                  <Col className="col-8 pt-5 ps-3">
                    <p className="name pt-2">
                      Maulana Dwi Wahyudi
                      <br />
                      <span>
                        <span className="color">As</span> Backend Engineer
                      </span>
                    </p>
                  </Col>
                </Row>
              </Card>
            </Col>
          </Row>
        </Container>
      </section>
      <Footer />
    </>
  );
}
