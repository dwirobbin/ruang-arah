import React from "react";
import { useLocation } from "react-router-dom";

import Footer from "../../components/Footer/Footer";
import Navbar from "../../components/Navbar/Navbar";
import UpdateQuestion from "../../components/Admin/UpdateQuestion/UpdateQuestion";

const Update = () => {
  const query = new URLSearchParams(useLocation().search);
  const questionId = parseInt(query.get("questionId"));

  return (
    <>
      <Navbar />
      <UpdateQuestion questionId={questionId} />
      <Footer />
    </>
  );
};

export default Update;
