import React from "react";

import Navbar from "../../components/Navbar/Navbar";
import Footer from "../../components/Footer/Footer";
import CreateQuestion from "../../components/Admin/CreateQuestion/CreateQuestion";

export default function Create() {
  return (
    <>
      <Navbar />
      <CreateQuestion />
      <Footer />
    </>
  );
}
