let component = ReasonReact.statelessComponent "REHeader";

let styles = {
  "container": ReactDOMRe.Style.make (),
  "text": ReactDOMRe.Style.make ()
};

let make ::items _children => {
  let handleItems index item => {
    let key = string_of_int index;
    let children = ReasonReact.stringToElement item;
    <h1 key style=styles##text >
      children  
    </h1>
  };
  {
    ...component,
    render: fun _self => {
      let renderedItems = items
        |> List.mapi handleItems
        |> Array.of_list
        |> ReasonReact.arrayToElement;
      <header style=styles##container >
        renderedItems
      </header>
    }
  }
};