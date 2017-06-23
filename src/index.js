import React from "react";
import ReactDOM from "react-dom";
import Animated from "animated/lib/targets/react-dom";
import Styled, { injectGlobal } from "styled-components";
import MenuIcon from "./components/menuIcon";

injectGlobal`
  body { margin: 0; }
`;

/**
 * Menu component.
 */
const MenuContainer = Styled.div`
  display: flex;
  flex: 1;
  padding: 50px;
  background-color: blue;
  flex-direction: column;
`;
const MenuOption = Styled.span`
  font-size: 32px;
`;
function Menu() {
  return (
    <MenuContainer>
      <MenuOption>hi</MenuOption>
      <br />
      <MenuOption>hi</MenuOption>
      <br />
      <MenuOption>hi</MenuOption>
      <br />
      <MenuOption>hi</MenuOption>
    </MenuContainer>
  );
}

/**
 * Animated component.
 */
const AnimatedContainer = {
  position: "absolute",
  top: "0",
  left: "0",
  width: "100%",
  height: "100%"
};
class AnimatedModal extends React.Component {
  animation = new Animated.Value(0);
  timing = 250;
  state = { show: false, display: false, animating: false };

  toggle = () => {
    if (this.state.animating) return;

    if (this.state.show) {
      this.hide();
    } else {
      this.show();
    }
  };

  show() {
    this.setState(
      () => ({ show: true, display: true, animating: true }),
      () => {
        this.animate(true);
        setTimeout(() => {
          this.setState(() => ({ animating: false }));
        }, this.timing);
      }
    );
  }

  hide() {
    this.setState(
      () => ({ show: false, animating: true }),
      () => {
        this.animate(false);
        setTimeout(() => {
          this.setState(() => ({ display: false, animating: false }));
        }, this.timing);
      }
    );
  }

  animate(status) {
    Animated.timing(this.animation, {
      toValue: status ? 1 : 0,
      duration: this.timing
    }).start();
  }

  render() {
    const opacity = this.animation;
    const display = this.state.display ? "flex" : "none";
    const style = { ...AnimatedContainer, opacity, display };
    return (
      <div>
        <MenuIcon open={this.state.show} onClick={this.toggle} />
        <Animated.div style={style} children={this.props.children} />
      </div>
    );
  }
}

/**
 * Main App.
 */
const Container = Styled.div`
  display: flex;
  flex: 1;
  min-height: 100vh;
  justify-content: center;
  align-items: center;
`;
const Title = Styled.span`
  font-size: 48px;
`;

function App() {
  return (
    <Container>
      <AnimatedModal>
        <Menu />
      </AnimatedModal>
      <Title>Hello</Title>
    </Container>
  );
}

/**
 * Render.
 */
ReactDOM.render(<App />, document.getElementById("root"));
