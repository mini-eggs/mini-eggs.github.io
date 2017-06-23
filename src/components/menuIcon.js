import React from "react";
import "./menuIcon.css";

export default class extends React.Component {
  render() {
    const aClass = this.props.open ? "navigation-icon open" : "navigation-icon";
    return (
      <div id="nav-icon3" className={aClass} onClick={this.props.onClick}>
        <span />
        <span />
        <span />
        <span />
      </div>
    );
  }
}
