import React, { PureComponent } from "react";
import Drawer from "material-ui/Drawer";
import Divider from "material-ui/Divider";
import { List, ListItem } from "material-ui/List";
import Subheader from "material-ui/Subheader";
import { push } from "react-router-redux";

export default class extends PureComponent {
  state = { open: true };

  handleTouchTap = (event, route) => {
    event.preventDefault();
    const { dispatch } = this.props;
    if (!route) {
      // nothing for mobile menu
    } else if (route.indexOf("http") > -1 || route.indexOf("mailto:") > -1) {
      this.props.onRequestChange();
      window.location.href = route;
    } else {
      this.props.onRequestChange();
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
        this.props.onRequestChange();
        this.props.dispatch(push(route));
      }
    );
  };

  render = () => {
    return (
      <Drawer
        docked={false}
        onRequestChange={this.props.onRequestChange}
        open={this.props.open}
      >
        <List>
          <Subheader>Evan Jones / {this.props.title}</Subheader>
          <Divider />
          {this.props.links.map((link, index) => {
            return (
              <ListItem
                key={index}
                onClick={event => {
                  this.handleTouchTap(event, link.route);
                }}
                primaryTogglesNestedList={true}
                primaryText={link.name}
                nestedItems={link.children.map((child, index) => {
                  return (
                    <ListItem
                      key={index}
                      primaryText={child.name}
                      onTouchTap={event => {
                        this.handleTouchTapChild(event, child.route);
                      }}
                    />
                  );
                })}
              />
            );
          })}
        </List>
      </Drawer>
    );
  };
}
