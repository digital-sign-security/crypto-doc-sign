import React from "react";
import { usePromiseTracker } from "react-promise-tracker";
import * as Loader from "react-loader-spinner";

function LoadingIndicator() {
  const { promiseInProgress } = usePromiseTracker();
  return (
    <>
      {promiseInProgress && (
        <div
          className="spinner"
          style={{
            position: "fixed",
            left: "42vw",
            top: "36vh",
            width: "10vw",
            height: "10vh",
            zIndex: "9999",
          }}
        >
          <Loader.TailSpin color="#445bad" height="100%" width="100%" />
        </div>
      )}
    </>
  );
}

export default LoadingIndicator;
