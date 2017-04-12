import React, { PureComponent } from "react";
import AppBar from "material-ui/AppBar";
import Desktop from "./desktop";
import Mobile from "./mobile";

export default class extends PureComponent {
  state = { open: false };

  drawer = () => {
    this.setState(({ open }) => {
      return { open: !open };
    });
  };

  render = () => {
    return (
      <AppBar width={300} onLeftIconButtonTouchTap={this.drawer} zDepth={0}>
        <Desktop {...this.props} />
        <Mobile
          {...this.props}
          onRequestChange={this.drawer}
          open={this.state.open}
        />
      </AppBar>
    );
  };
}
