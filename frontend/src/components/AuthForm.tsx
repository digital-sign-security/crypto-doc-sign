import React, { useState } from "react";
import { Field, Form, Formik } from "formik";

import "../styles/Form.css";
import styled from "styled-components";
import { FormHeading, SubmitButton } from "./RegistrationForm";
import { validateRequiredFields } from "../utils/validate";
import { postAuth } from "../utils/api";
import { useNavigate } from "react-router";
import { trackPromise } from "react-promise-tracker";

function AuthForm({ onClick }: any) {
  const [errMessage, setErrMessage] = useState(false);

  const navigate = useNavigate();

  return (
    <Formik
      initialValues={{
        login: "",
        password: "",
      }}
      onSubmit={async (initialValues) => {
        trackPromise(
          postAuth(initialValues.login, initialValues.password).then(() => {
            if (localStorage.getItem("isCorrect") === "false") {
              setErrMessage(true);
            } else navigate("/courses");
          })
        );
      }}
    >
      {({ errors, touched }) => (
        <Form className="form">
          <FormHeading>Войти в аккаунт</FormHeading>
          <div>
            <Field
              className={`form__field  ${
                errors.login && touched.login
                  ? "form__invalid-field"
                  : touched.login
                  ? "form__valid-field"
                  : ""
              }`}
              type="text"
              name="login"
              id="login"
              placeholder="Введите логин"
              validate={validateRequiredFields}
            ></Field>

            <div>
              <Field
                className={`form__field  ${
                  errors.password && touched.password
                    ? "form__invalid-field"
                    : touched.password
                    ? "form__valid-field"
                    : ""
                }`}
                type="text"
                name="password"
                id="password"
                placeholder="Введите пароль"
                validate={validateRequiredFields}
              ></Field>
            </div>
          </div>
          {errMessage && <ErrorMessage>Неверный логин или пароль</ErrorMessage>}
          <ButtonContainer>
            <LinkToOtherForm onClick={onClick}>Нет аккаунта?</LinkToOtherForm>
            <div>
              <SubmitButton
                type="submit"
                style={{ backgroundColor: "#e8e6e8" }}
                value="default"
              >
                Продолжить
              </SubmitButton>
            </div>
          </ButtonContainer>
        </Form>
      )}
    </Formik>
  );
}

export const ErrorMessage = styled.p`
  color: red;
  text-align: center;
`;

export const LinkToOtherForm = styled.p`
  cursor: pointer;
  color: blue;
`;

export const ButtonContainer = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

export default AuthForm;
