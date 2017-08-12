import React from "react";
import { render } from "react-snapshot";

import { jsComponent as REContainer } from "./lib/js/re/components/container.js";
import registerServiceWorker from "./registerServiceWorker";
import "./styles/main.css";
import "./styles/header.css";
import "./styles/body.css";
import "./styles/footer.css";
import "./styles/mobile.css";

render(<REContainer />, document.getElementById("root"));
registerServiceWorker();
