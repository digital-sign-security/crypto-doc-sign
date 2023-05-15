import React from "react";

import styled from "styled-components";

function Footer() {
  return (
    <>
      <FooterBar>
        <FooterContent>
          <div>
            <p>Сделано с любовью</p>
          </div>
        </FooterContent>
      </FooterBar>
    </>
  );
}

const FooterContent = styled.footer`
  font-weight: 700;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 86%;
  margin: 0 auto;
  height: 15vh;
  margin-top: -14vh;
`;

const FooterBar = styled.div`
  background-color: #99eb90;
`;

export default Footer;
