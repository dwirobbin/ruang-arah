import React from "react";
import Main from "./Main/Main";
import Navbar from "../../components/Navbar/Navbar";
import LearningPath from "./LearningPath/LearningPath";

function Home() {
  return (
    <>
      <Navbar />
      <Main />
      <LearningPath />
    </>
  );
}

export default Home;
