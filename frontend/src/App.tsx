import { useState, useEffect } from "react";
import HomePage from "./pages/HomePage";
import QuizzesPage from "./pages/QuizzesPage";

function App() {
  const [currentPath, setCurrentPath] = useState(window.location.pathname);

  // Handle back/forward navigation
  useEffect(() => {
    const handlePopState = () => {
      setCurrentPath(window.location.pathname);
    };

    window.addEventListener("popstate", handlePopState);
    return () => window.removeEventListener("popstate", handlePopState);
  }, []);

  // Navigate and update state
  const navigateTo = (path: string) => {
    if (path !== window.location.pathname) {
      window.history.pushState({}, "", path);
      setCurrentPath(path);
    }
  };

  let PageComponent;
  switch (currentPath) {
    case "/quizzes":
      PageComponent = <QuizzesPage navigateTo={navigateTo} />;
      break;
    case "/":
    default:
      PageComponent = <HomePage navigateTo={navigateTo} />;
  }

  return <>{PageComponent}</>;
}

export default App;
