import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import ErrorMessage from "../ErrorMessage/ErrorMessage";
import "./style.css";
import API from "../../api/Api";
import axios from "axios";

const Question = ({
  questions,
  setQuestions,
  options,
  pageQuestion,
  setPageQuestion,
  score,
  setScore,
  programmingLanguageId,
  correct,
}) => {
  const [selected, setSelected] = useState();
  const [error, setError] = useState(false);

  const navigate = useNavigate();

  const handleSelect = (index) => {
    if (selected === index && selected === correct) return "select";
    else if (selected === index && selected !== correct) return "wrong";
    else if (index === correct) return "select";
  };

  const handleCheck = (index) => {
    setSelected(index);
    let values = {
      question_id: questions[0].id,
      answer: index,
    };

    if (pageQuestion === 1) {
      let data = { answers: [values] };
      localStorage.setItem("Answers", JSON.stringify(data));
    } else if (pageQuestion > 1) {
      let data = JSON.parse(localStorage.getItem("Answers"));
      data.answers.push(values);
      localStorage.setItem("Answers", JSON.stringify(data));
    }

    if (index === correct) {
      setScore(score + 1);
    }

    setError(false);
  };

  const handleNext = async () => {
    console.log(localStorage.getItem("Answers"));
    if (pageQuestion > 9) {
      let auth = localStorage.getItem("token");
      let answers = localStorage.getItem("Answers");

      try {
        let { data: res } = await axios.post(
          `${API.API_URL}/api/home/process-and-result`,
          answers,
          {
            headers: {
              Accept: "/",
              "Content-Type": "application/json",
              Authorization: "Bearer " + auth,
            },
          }
        );

        localStorage.setItem("Level", JSON.stringify(res.data));
        navigate("/quiz/result");
      } catch (err) {}
    } else if (selected) {
      navigate({
        pathname: "/quiz",
        search: `?programmingLanguageId=${programmingLanguageId}&page=${
          pageQuestion + 1
        }`,
      });

      setPageQuestion(pageQuestion + 1);
      setSelected();
    } else setError("Please select an answer");
  };

  const handleQuit = () => {
    navigate({ pathname: "/home" });
    setPageQuestion(0);
    setQuestions();
  };

  return (
    <section id="question-page">
      <div className="container">
        <h1>Question</h1>
        <div className="card">
          <div className="card-body">
            <div className="row card-title mb-5 mt-3">
              <div className="quizInfo">
                <span> Soal ke - {pageQuestion} / 10</span>
                <span>Score: {score}</span>
              </div>
            </div>

            <div className="card-text">
              {questions.map((item, index) => (
                <h5 key={index} className="mb-4">
                  {item.question}
                </h5>
              ))}
              <div className="options">
                {error && <ErrorMessage>{error}</ErrorMessage>}
                {options.map((index) => (
                  <button
                    className={`singleOption  ${
                      selected && handleSelect(index)
                    }`}
                    key={index}
                    onClick={() => handleCheck(index)}
                    disabled={selected}
                    value={index}
                  >
                    {index}
                  </button>
                ))}
              </div>
            </div>
          </div>
          <div className="row card-title controls">
            <button
              type="button"
              className="btn btn-danger mr-3"
              style={{ width: 170 }}
              onClick={() => handleQuit()}
            >
              Quit
            </button>
            <button
              type="button"
              className="btn btn-warning mr-3"
              style={{ width: 170 }}
              onClick={() => handleNext()}
            >
              {pageQuestion > 9 ? "Submit" : "Next Question"}
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Question;
