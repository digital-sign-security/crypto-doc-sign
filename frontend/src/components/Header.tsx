import React from "react";

import styled from "styled-components";

function Header() {
  return (
    <HeadBar>
      <p>Крипто-месенджер</p>
      <Navigation>
        <Button>Отправить сообщение</Button>
        <Button>Выйди из аккаунта</Button>
      </Navigation>
    </HeadBar>
  );
}

const HeadBar = styled.header`
  font-weight: 700;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 86%;
  margin: 0 auto;
  height: 64px;
`;

const Navigation = styled.nav`
  display: flex;
  justify-content: space-between;
  gap: 10px;
`;

const Button = styled.button`
  border: none;
  background-color: #99eb90;
  padding: 8px;
  border-radius: 10px;
  cursor: pointer;
`;

export default Header;
