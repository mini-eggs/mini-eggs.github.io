import React from "react";
import { connect } from "react-redux";
import Error from "../error/";
import Page from "./component";

function findPage(pages, pathname) {
  const find = pathname.replace("/page/", "");
  let status = false;
  pages.forEach(page => {
    if (page.link === find) {
      status = page;
    }
  });
  return status;
}

function pageContainer({ history, pages }) {
  const page = findPage(pages, history.location.pathname);
  return page ? <Page {...page} /> : <Error />;
}

export default connect(({ Pages }) => {
  return { pages: Pages };
})(pageContainer);
