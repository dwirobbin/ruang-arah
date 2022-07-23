import React, { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import API from "../../api/Api";
import axios from "axios";
import Footer from "../../components/Footer/Footer";
import Question from "../../components/Question/Question";
import Navbar from "../../components/Navbar/Navbar";

const Quiz = () => {
  let query = new URLSearchParams(useLocation().search);
  let programmingLanguageId = parseInt(query.get("programmingLanguageId"));
  let pageNo = parseInt(query.get("page"));

  const [questions, setQuestions] = useState([]);
  const [options, setOptions] = useState([]);
  const [pageQuestion, setPageQuestion] = useState(pageNo);
  const [score, setScore] = useState(0);

  useEffect(() => {
    const fetchQuestions = async () => {
      try {
        let auth = localStorage.getItem("token");
        const { data: res } = await axios.get(
          `${API.API_URL}/api/home/questions?programmingLanguageId=${programmingLanguageId}&page=${pageQuestion}&limit=1`,
          {
            headers: {
              Accept: "/",
              "Content-Type": "application/json",
              Authorization: "Bearer " + auth,
            },
          }
        );

        setQuestions(res.data);

        setOptions(
          [res.data[0].correct_answer, ...res.data[0].incorrect_answers].sort(
            () => Math.random() - 0.5
          )
        );
      } catch (error) {}
    };

    fetchQuestions();
  }, [pageQuestion, programmingLanguageId]);

  return (
    <>
      <Navbar />
      <Question
        questions={questions}
        setQuestions={setQuestions}
        options={options}
        score={score}
        setScore={setScore}
        pageQuestion={pageQuestion}
        setPageQuestion={setPageQuestion}
        programmingLanguageId={programmingLanguageId}
        correct={questions[0]?.correct_answer}
      />
      <Footer />
    </>
  );
};

export default Quiz;
