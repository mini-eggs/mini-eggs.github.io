import React, { PureComponent } from "react";
import FlatButton from "material-ui/FlatButton";
import Popover from "material-ui/Popover";
import Menu from "material-ui/Menu";
import MenuItem from "material-ui/MenuItem";
import { push } from "react-router-redux";
import { NavContainer, LinkContainer } from "./styles.js";

class DesktopButton extends PureComponent {
  state = { open: false };

  handleTouchTap = event => {
    event.preventDefault();
    const { dispatch, route } = this.props;
    if (!route) {
      this.setState({
        open: true,
        anchorEl: event.currentTarget
      });
    } else if (route.indexOf("http") > -1 || route.indexOf("mailto:") > -1) {
      window.location.href = route;
    } else {
      dispatch(push(route));
    }
  };

  handleTouchTapChild = (event, route) => {
    event.preventDefault();
    this.setState(
      () => {
        return { open: false };
      },
      () => {
        this.props.dispatch(push(route));
      }
    );
  };

  handleRequestClose = () => {
    this.setState(() => {
      return { open: false };
    });
  };

  render() {
    return (
      <LinkContainer>
        <FlatButton onTouchTap={this.handleTouchTap}>
          {this.props.name}
        </FlatButton>
        {(this.props.children || []).length > 0
          ? <Popover
              open={this.state.open}
              anchorEl={this.state.anchorEl}
              anchorOrigin={{ horizontal: "right", vertical: "bottom" }}
              targetOrigin={{ horizontal: "right", vertical: "top" }}
              onRequestClose={this.handleRequestClose}
            >
              <Menu>
                {this.props.children.map(({ name, route }, yindex) => {
                  return (
                    <MenuItem
                      key={yindex}
                      onTouchTap={event => {
                        this.handleTouchTapChild(event, route);
                      }}
                      primaryText={name}
                    />
                  );
                })}
              </Menu>
            </Popover>
          : null}
      </LinkContainer>
    );
  }
}

export default function(props) {
  return (
    <NavContainer>
      {props.links.map((link, index) => {
        return <DesktopButton key={index} {...link} {...props} />;
      })}
    </NavContainer>
  );
}
