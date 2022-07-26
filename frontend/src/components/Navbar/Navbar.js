import React, { useContext } from "react";
import { useNavigate } from "react-router-dom";
import AuthContext from "../../store/AuthContext";

import "./navbar.css";
import Logo from "../../../src/assets/images/Logo.png";

// React Boostrap
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";

function HomeNavbar() {
  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();

  const logoutHandler = () => {
    authCtx.logout();
    navigate("/");
  };

  return (
    <section id="navbar-home">
      <Navbar className="nav" expand="lg" variant="light">
        <Container>
          <Row>
            <Col>
              <img src={Logo} className="img-fluid" alt="logo" />
            </Col>
          </Row>
          <Row>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav" className="nav-content mx-3">
              <Nav className="me-auto">
                {authCtx.role === "guest" && (
                  <>
                    <a href="#main-pages">
                      <Nav className="text-light pe-3 mt-2">Home</Nav>
                    </a>
                    <a href="#learning-path">
                      <Nav className="text-light pe-4 mt-2">Learning Path</Nav>
                    </a>
                  </>
                )}
                <Nav className="text-light">
                  <Button
                    className="button-logout"
                    variant="warning"
                    onClick={logoutHandler}
                  >
                    Logout
                  </Button>
                </Nav>
              </Nav>
            </Navbar.Collapse>
          </Row>
        </Container>
      </Navbar>
    </section>
  );
}

export default HomeNavbar;
