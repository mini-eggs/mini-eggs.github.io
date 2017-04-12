import { combineReducers } from "redux";
import { routerReducer, routerMiddleware } from "react-router-redux";
import PagesReducer from "./pages";
import LinksReducer from "./links";

const history = createHistory();
const middleware = routerMiddleware(history);

export default combineReducers({
  PagesReducer,
  LinksReducer,
  route: routerReducer
});
