import "./style.css";
import ListLanguage from "./../../../components/Language/ListLanguage";

function Main({ listLanguage, pageQuestion }) {
  return (
    <section id="language-page">
      <div className="container">
        <h1>
          Bahasa <span> Pemrograman </span>{" "}
        </h1>
        <div className="card">
          <div className="card-content">
            {listLanguage.map((language) => (
              <ListLanguage
                key={language.id}
                image={language.image_url}
                language={language.name}
                id={language.id}
                pageQuestion={pageQuestion}
              />
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}

export default Main;
