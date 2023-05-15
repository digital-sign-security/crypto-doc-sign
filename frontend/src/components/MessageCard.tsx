import React from "react";

import styled from "styled-components";
import messageLogo from "../assets/images/messageIcon.svg";

interface MessageData {
  userName: string;
  topic: string;
  date: string;
}

function MessageCard(data: MessageData) {
  return (
    <Card>
      <img src={messageLogo} alt="Message icon" />
      <p>Отправитель: {data.userName}</p>
      <p>Тема: {data.topic}</p>
      <p>Дата: {data.date}</p>
      <Button>Прочитать</Button>
    </Card>
  );
}

const Card = styled.div`
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100%;
  background-color: #67e19f;
  height: 10vh;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 13px;
`;

export const Button = styled.button`
  border: none;
  background: white;
  padding: 8px;
  border-radius: 10px;
  cursor: pointer;
`;

export default MessageCard;
