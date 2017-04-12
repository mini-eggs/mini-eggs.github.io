import React from "react";
import ReactDOM from "react-dom";
import "./styles/main.css";
import injectTapEventPlugin from "react-tap-event-plugin";
injectTapEventPlugin();
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider";
import getMuiTheme from "material-ui/styles/getMuiTheme";
import createHistory from "history/createHashHistory";
import { Route, Switch, Redirect } from "react-router";
import { HashRouter } from "react-router-dom";
import { createStore, combineReducers, applyMiddleware } from "redux";
import { Provider } from "react-redux";
import { routerReducer, routerMiddleware } from "react-router-redux";
import Welcome from "./components/welcome/";
import Page from "./components/page/";
import Pages from "./reducers/pages";
import Links from "./reducers/links";
import PagesContent from "./content/pages/";
import LinksContent from "./content/links/";

const theme = getMuiTheme({
  palette: {
    primary1Color: "transparent",
    alternateTextColor: "black"
  }
});

const preloaded = {
  Pages: PagesContent,
  Links: LinksContent
};

const history = createHistory();

const middleware = routerMiddleware(history);

const reducers = combineReducers({
  Pages,
  Links,
  route: routerReducer
});

const store = createStore(reducers, preloaded, applyMiddleware(middleware));

ReactDOM.render(
  <MuiThemeProvider muiTheme={theme}>
    <Provider store={store}>
      <HashRouter>
        <div>
          <Welcome />
          <Switch>
            <Route path="/page/:page" component={Page} />
            <Redirect from="/" to="/page/home" />
          </Switch>
        </div>
      </HashRouter>
    </Provider>
  </MuiThemeProvider>,
  document.getElementById("root")
);
