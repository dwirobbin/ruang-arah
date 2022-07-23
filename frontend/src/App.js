import { Routes, Route, Navigate } from "react-router-dom";
import PrivateRoute from "./routes/PrivateRoutes";
import {
  LandingPage,
  LoginPage,
  RegisterPage,
  HomePage,
  LanguagePage,
  QuizPage,
  ResultPage,
  RecommendationPage,
  DashboardPage,
  CreateQuestionPage,
  UpdateQuestionPage,
} from "./pages/switch.js";

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />}></Route>
      <Route path="/login" element={<LoginPage />}></Route>
      <Route path="/register" element={<RegisterPage />}></Route>

      <Route path="" element={<PrivateRoute />}>
        <Route path="/home" element={<HomePage />}></Route>
        <Route path="/languages" element={<LanguagePage />}></Route>
        <Route path="/quiz" element={<QuizPage />}></Route>
        <Route path="/quiz/result" element={<ResultPage />}></Route>
        <Route
          path="/quiz/result/recommendation"
          element={<RecommendationPage />}
        ></Route>
        <Route path="/dashboard" element={<DashboardPage />}></Route>
        <Route
          path="/questions/create"
          element={<CreateQuestionPage />}
        ></Route>
        <Route
          path="/questions/update"
          element={<UpdateQuestionPage />}
        ></Route>
      </Route>

      <Route path="*" element={<Navigate to="/" />}></Route>
    </Routes>
  );
}

export default App;
