import React from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import { Container, Table, Button, Row, Col, Card } from "react-bootstrap";

import "./style.css";
import API from "../../../api/Api";

const Dashboard = ({ auth, fetchQuestions, questions }) => {
  const navigate = useNavigate();

  function deleteQuestion(questionId) {
    axios.delete(
      `${API.API_URL}/api/admin/questions/delete?questionId=${questionId}`,
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          Authorization: "Bearer " + auth,
        },
      }
    );
    fetchQuestions();
  }

  function handleUpdate(questionId) {
    navigate({
      pathname: "/questions/update",
      search: `?questionId=${questionId}`,
    });
  }

  return (
    <section id="dashboard-pages">
      <div className="profile-name pt-3">
        <h1 className="mb-5">Welcome, Admin!</h1>
      </div>
      <div>
        <Container>
          <Row className="mb-5">
            <Col>
              <Card className="card float-left text-center">
                <p className="my-auto mx-auto">
                  Jumlah Soal : {questions.length}
                </p>
              </Card>
            </Col>
            <Col className="float-sm-right">
              <Link to="/questions/create">
                <Button
                  variant="warning"
                  className="button-create float-sm-right"
                >
                  New Question
                </Button>
              </Link>
            </Col>
          </Row>
          <Table className="table">
            <thead>
              <tr>
                <th>Programming Language</th>
                <th>Question</th>
                <th>Update</th>
                <th>Delete</th>
              </tr>
            </thead>
            <tbody id="table-body" className="table-body">
              {questions.map((question) => (
                <tr className="tr-class" key={question.id}>
                  <td className="my-auto">{question.programming_language}</td>
                  <td className="questions">{question.question}</td>
                  <td>
                    <Button
                      variant="warning"
                      onClick={() => handleUpdate(question.id)}
                    >
                      Update
                    </Button>
                  </td>
                  <td>
                    <Button
                      variant="danger"
                      onClick={() => deleteQuestion(question.id)}
                    >
                      Delete
                    </Button>
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
        </Container>
      </div>
    </section>
  );
};

export default Dashboard;
