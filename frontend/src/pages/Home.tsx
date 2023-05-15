import React from "react";
import Header from "../components/Header";
import Footer from "../components/Footer";
import MessageCard from "../components/MessageCard";
import styled from "styled-components";

function Home() {
  return (
    <>
      <Header></Header>
      <Main>
        <Swiper>
          <Navigation>
            <p>Входящие</p>
            <p>Отправленные</p>
          </Navigation>
        </Swiper>
        <MessageCard
          userName="Данька"
          topic={"Пупупу"}
          date={"20.10.2001"}
        ></MessageCard>
      </Main>
      <Footer></Footer>
    </>
  );
}

const Swiper = styled.section`
  width: 14vw;
`;

const Navigation = styled.nav`
  display: flex;
  justify-content: space-between;
  padding-bottom: 20px;
`;

const Main = styled.main`
  margin: 0 auto;
`;

export default Home;
