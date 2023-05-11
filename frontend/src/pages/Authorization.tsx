import React, { useState } from "react";

import styled from "styled-components";

import AuthForm from "../components/AuthForm";
import RegitrstionForm from "../components/RegistrationForm";

function Authorization() {
  const [isRegistered, setRegistered] = useState(false);

  return (
    <FormMain>
      <AuthorizationContainer>
        {(!isRegistered && (
          <AuthForm onClick={() => setRegistered(true)} />
        )) || <RegitrstionForm onClick={() => setRegistered(false)} />}
      </AuthorizationContainer>
    </FormMain>
  );
}

const FormMain = styled.main`
  height: 100vh;
`;

const AuthorizationContainer = styled.div`
  width: 100%;
  height: 100%;
  position: fixed;
  left: 0;
  top: 0;
  display: flex;
  justify-content: center;
  align-items: center;
`;

export default Authorization;
