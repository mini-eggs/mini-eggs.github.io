import React from "react";
import Styled from "styled-components";

// const image = "https://static.pexels.com/photos/195399/pexels-photo-195399.jpeg";

const Container = Styled.div`
  height: 25vh;
  background-attachment: fixed;
  background-size: cover;
  background-position: center center;
  display: flex;
`;
const Content = Styled.div`
  display: flex;
  flex: 1;
  justify-content: flex-start;
  align-items: center;
  background-color: rgba(241,241,241,1);
`;
const Title = Styled.h1`
  padding-left: 25px;
  font-size: 32px;
  color: black;
`;

export default function() {
  return (
    <Container>
      <Content>
        <div className="max-width">
          <Title>Evan Jones,<br />web/app developer</Title>
        </div>
      </Content>
    </Container>
  );
}
