import axios from "axios";

const config = {
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
    "Content-Type": "application/json",
  },
};

export const postAuth = async (login: string, password: string) => {
  return await axios
    .post(
      "https://ed-platform-6106.nh2023.codenrock.com/api/login",
      {
        email: login,
        password: password,
      },
      config
    )
    .then((data) => {
      localStorage.removeItem("isCorrect");
      localStorage.setItem("user", JSON.stringify(data.data));
    })
    .catch((data) => {
      localStorage.setItem("isCorrect", "false");
      console.log(data.response);
    });
};

export const postRegistrationData = async (
  username: string,
  email: string,
  password: string
) => {
  return axios
    .post(
      "https://ed-platform-6106.nh2023.codenrock.com/api/register",
      {
        username: username,
        email: email,
        password: password,
      },
      config
    )
    .then(() => {
      localStorage.removeItem("isCorrectRegField");
    })
    .catch((data) => {
      console.log(data.response);
      localStorage.setItem("isCorrectRegField", "false");
    });
};
