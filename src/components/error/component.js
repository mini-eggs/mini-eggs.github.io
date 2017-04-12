import React from "react";
import { Link } from "react-router-dom";

export default function() {
  return (
    <Link
      to="/"
      style={{
        color: "white",
        textAlign: "center",
        textDecoration: "none",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "50vh"
      }}
    >
      Page not found,<br />go home?
    </Link>
  );
}
