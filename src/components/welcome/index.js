import { createElement } from "react";
import Loadable from "react-loadable";

export default function(props) {
  return createElement(
    Loadable({
      loader: () => import("./component"),
      LoadingComponent: () => null
    }),
    props
  );
}
