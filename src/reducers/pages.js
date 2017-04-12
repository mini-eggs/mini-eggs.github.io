const initialState = {
  pages: []
};

export default function(state = initialState, action) {
  switch (action.type) {
    case "SET_PAGES": {
      return Object.assign({}, state, action.payload);
    }
    default: {
      return state;
    }
  }
}
