import { replaceAll } from "../../utilities/";
import Home from "./home";
import WhatIsTheDom from "./blog/whatIsTheDom";
import Learning from "./blog/learning";
import Emoji from "./blog/emoji";
import Flippour from "./projects/flippour/";
import Triangly from "./projects/triangly/";
import CRASSR from "./projects/crassr/";

function fixBody(page) {
  page.body = replaceAll(page.body, "  ", "");
  page.body = replaceAll(page.body, "\n", "");
  return page;
}

const allPages = [
  Home,
  Flippour,
  WhatIsTheDom,
  Learning,
  Emoji,
  Flippour,
  Triangly,
  CRASSR
];

export default allPages.map(page => fixBody(page));
