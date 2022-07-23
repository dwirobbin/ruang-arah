import React from "react";
import API from "../../api/Api";
import { useState, useEffect } from "react";
import axios from "axios";
import Main from "./Main/Main";
import Footer from "../../components/Footer/Footer";

function Language() {
  const [listLanguage, setListLanguage] = useState([]);
  const [pageQuestion] = useState(0);

  const fetchListPost = async () => {
    try {
      let auth = localStorage.getItem("token");
      let { data: resp } = await axios.get(
        `${API.API_URL}/api/home/languages`,
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
            Authorization: "Bearer " + auth,
          },
        }
      );

      setListLanguage(resp.data);
    } catch (err) {}
  };

  useEffect(() => {
    fetchListPost();
  }, []);

  return (
    <>
      <Main listLanguage={listLanguage} pageQuestion={pageQuestion} />
      <Footer />
    </>
  );
}

export default Language;
