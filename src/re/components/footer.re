let component = ReasonReact.statelessComponent "REFooter";

let styles = {
  "footer": ReactDOMRe.Style.make (),
  "link": ReactDOMRe.Style.make (),
};

let make ::items _children => {
  let handleItems index item => {
    let key = string_of_int index;
    let stringItem = if ( index == 0 ) {
      item##text;
    } else {
      [%bs.raw {| " • " |}] ^ item##text;
    };
    <a key style=styles##link href=item##link >
      ( ReasonReact.stringToElement stringItem )
    </a>
  };
  {
    ...component,
    render: fun _self => {
      let renderedItems = items
        |> List.mapi handleItems
        |> Array.of_list
        |> ReasonReact.arrayToElement;
      <footer style=styles##footer >
        renderedItems 
      </footer>
    }
  };
};