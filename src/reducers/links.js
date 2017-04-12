const initialState = {
  links: []
};

export default function(state = initialState, action) {
  switch (action.type) {
    case "SET_LINKS": {
      return Object.assign({}, state, action.payload);
    }
    default: {
      return state;
    }
  }
}
