import React from "react";
import Paper from "material-ui/Paper";
import Nav from "../nav/";
import { Container, Image, Left, Right, Content, Dark } from "./styles";

export default function Page(props) {
  return (
    <div className="max-width">
      <Nav {...props} />
      <Paper
        zDepth={3}
        style={{
          marginTop: "25px",
          marginBottom: "50px",
          marginLeft: "5px",
          marginRight: "5px",
          borderRadius: 3,
          overflow: "hidden"
        }}
      >
        <Container>
          <Right>
            <Image style={{ backgroundImage: `url("${props.image}")` }}>
              {/*<Dark />*/}
            </Image>
          </Right>
          <Left>
            <Content dangerouslySetInnerHTML={{ __html: props.body }} />
          </Left>
        </Container>
      </Paper>
    </div>
  );
}
