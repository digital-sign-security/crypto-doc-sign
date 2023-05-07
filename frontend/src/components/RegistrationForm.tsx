import { Field, Form, Formik } from "formik";
import React, { useState } from "react";
import styled from "styled-components";
import { postAuth, postRegistrationData } from "../utils/api";
import { useNavigate } from "react-router-dom";
import { validateEmail, validateRequiredFields } from "../utils/validate";
import { ButtonContainer, ErrorMessage, LinkToOtherForm } from "./AuthForm";
import { trackPromise } from "react-promise-tracker";

function RegitrstionForm({ onClick }: any) {
  const [isCorrect, setCorrect] = useState<boolean>(true);
  const navigate = useNavigate();

  return (
    <Formik
      initialValues={{
        username: "",
        email: "",
        password: "",
      }}
      onSubmit={(initialValues) => {
        trackPromise(
          postRegistrationData(
            initialValues.username,
            initialValues.email,
            initialValues.password
          ).then(() => {
            postAuth(initialValues.email, initialValues.password).then(() => {
              if (localStorage.getItem("isCorrectRegField") === "false") {
                setCorrect(false);
              } else {
                navigate("/courses");
              }
            });
          })
        );
      }}
    >
      {({ errors, touched }) => (
        <Form className="form">
          <FormHeading>Регистрация</FormHeading>
          <div>
            <div>
              <Field
                className={`form__field  ${
                  errors.username && touched.username
                    ? "form__invalid-field"
                    : touched.username
                    ? "form__valid-field"
                    : ""
                }`}
                type="text"
                name="username"
                id="username"
                placeholder="Введите позывной"
                validate={validateRequiredFields}
              ></Field>
            </div>

            <div>
              <Field
                className={`form__field  ${
                  errors.email && touched.email
                    ? "form__invalid-field"
                    : touched.email
                    ? "form__valid-field"
                    : ""
                }`}
                type="text"
                name="email"
                id="email"
                placeholder="Введите почту"
                validate={validateEmail}
              ></Field>
            </div>

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
          {!isCorrect && (
            <ErrorMessage>Введены некорректные данные</ErrorMessage>
          )}

          <ButtonContainer>
            <LinkToOtherForm onClick={onClick}>Войти в аккаунт</LinkToOtherForm>
            <div>
              <SubmitButton type="submit" value="default">
                Продолжить
              </SubmitButton>
            </div>
          </ButtonContainer>
        </Form>
      )}
    </Formik>
  );
}

export const SubmitButton = styled.button`
  border: none;
  width: 100px;
  background-color: №f7c6c6;
  cursor: pointer;
  border-radius: 20px;
  height: 30px;
`;

export const FormHeading = styled.h2`
  text-align: center;
  margin-top: 0;
`;

export default RegitrstionForm;
