import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Container, Form, Button, Row, Col } from "react-bootstrap";
import axios from "axios";

import "./style.css";
import API from "./../../../api/Api";

const CreateQuestion = () => {
  const navigate = useNavigate();
  const [listLanguages, setListLanguages] = useState([]);

  const [programming_language, setProgramming_language] = useState("");
  const [question, setQuestion] = useState("");
  const [correctAnswer, setCorrectAnswer] = useState("");
  const [incorrectOne, setIncorrectOne] = useState("");
  const [incorrectTwo, setIncorrectTwo] = useState("");

  const fetchListLanguages = async () => {
    const { data: resp } = await axios.get(
      `${API.API_URL}/api/home/languages`,
      {
        headers: {
          Accept: "/",
          "Content-Type": "application/json",
          Authorization: "Bearer " + localStorage.getItem("token"),
        },
      }
    );

    setListLanguages(resp.data);
  };

  const handleCreate = async (event) => {
    event.preventDefault();
    try {
      let { data: resp } = await axios.post(
        `${API.API_URL}/api/admin/questions/create`,
        {
          programming_language: programming_language,
          question: question,
          correct_answer: correctAnswer,
          incorrect_one: incorrectOne,
          incorrect_two: incorrectTwo,
        },
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
            Authorization: "Bearer " + localStorage.getItem("token"),
          },
        }
      );

      if (resp.code === 200) {
        navigate("/dashboard");
      }
    } catch (error) {}
  };

  useEffect(() => {
    fetchListLanguages();
  }, []);

  return (
    <section id="create-question">
      <Container className="container-outside">
        <Form onSubmit={handleCreate}>
          <h1 className="mb-5">Create Question</h1>
          <Row>
            <Col lg={6} md={6} sm={12} xs={12}>
              <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label className="card-title">
                  Programming Language
                </Form.Label>
                <Form.Select
                  size="lg"
                  onChange={(e) => setProgramming_language(e.target.value)}
                  value={programming_language}
                  name="programming_language"
                >
                  <option label="Select Programming Language" value=""></option>
                  {listLanguages.map((language) => (
                    <option key={language.id} value={language.name}>
                      {language.name}
                    </option>
                  ))}
                </Form.Select>
              </Form.Group>

              <Form.Group className="mb-3" controlId="formBasicPassword">
                <Form.Label className="card-title">Question</Form.Label>
                <Form.Control
                  as="textarea"
                  className="mx-auto form-area"
                  onChange={(event) => setQuestion(event.target.value)}
                />
              </Form.Group>
            </Col>

            <Col
              lg={6}
              md={6}
              sm={12}
              xs={12}
              style={{
                justifyContent: "space-between",
                display: "flex",
                flexDirection: "column",
              }}
            >
              <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label className="card-title">Correct Answer</Form.Label>
                <Form.Control
                  size="lg"
                  type="text"
                  className="mx-auto form-input"
                  onChange={(event) => setCorrectAnswer(event.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3 " controlId="formBasicEmail">
                <Form.Label className="card-title">Incorrect One</Form.Label>
                <Form.Control
                  size="lg"
                  type="text"
                  className="mx-auto form-input"
                  onChange={(event) => setIncorrectOne(event.target.value)}
                />
              </Form.Group>

              <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label className="card-title">Incorrect Two</Form.Label>
                <Form.Control
                  size="lg"
                  type="text"
                  className="mx-auto form-input"
                  onChange={(event) => setIncorrectTwo(event.target.value)}
                />
              </Form.Group>
            </Col>
            <Button variant="warning button" className="mt-5" type="submit">
              Create
            </Button>
          </Row>
        </Form>
      </Container>
    </section>
  );
};

export default CreateQuestion;
