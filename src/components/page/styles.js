import Styled from "styled-components";

export const Container = Styled.article`
  display: flex;
  flex-direction: row-reverse;
  flex: 1;
  min-height: calc(75vh - 64px);
`;
export const Content = Styled.div`
  padding: 50px 50px 500px;
  line-height: 160%;
  background-color: white;
  @media (max-width: 748px) {
    padding: 50px 15px 150px;
  }
`;
export const Image = Styled.div`
  background-attachment: fixed;
  background-size: cover;
  background-position: center center;
  display: flex;
  flex: 1;
`;
export const Left = Styled.div`
  transition-duration: 250ms;
  transform-origin: right;
  flex-direction: column;
  display: flex;
  flex: 1;
`;
export const Right = Styled.div`
  transition-duration: 250ms;
  transform-origin: left;
  flex-direction: column;
  display: flex;
  flex: 1;
  @media (max-width: 748px) {
    display: none;
  }
`;
export const Title = Styled.h2`
  margin: 0;
`;
export const Break = Styled.hr`
  margin: 40px 0 50px;
  border-width: 0;
  background-color: black;
  height: 2px;
  border-radius: 1px;
`;
