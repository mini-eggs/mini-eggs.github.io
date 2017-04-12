import { connect } from "react-redux";
import Nav from "./component";

export default connect(({ Links }) => {
  return { links: Links };
})(Nav);
