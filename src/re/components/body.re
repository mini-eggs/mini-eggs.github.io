let component = ReasonReact.statelessComponent "REBody";

let styles = {
  "container": ReactDOMRe.Style.make (),
  "section": ReactDOMRe.Style.make (),
  "span": ReactDOMRe.Style.make (),
  "link": ReactDOMRe.Style.make (),
};

let make ::items _children => {
  let handleItems index item => {
    let key = string_of_int index;
    <section key style=styles##section >
      <span style=styles##span >  
        ( ReasonReact.stringToElement ( item##pre ^ " " ) )
      </span>
      <span className=( "link link-" ^ key ) style=styles##span >  
        <a style=styles##link href=item##link >
          ( ReasonReact.stringToElement item##text )
        </a>
      </span>
    </section>
  };
  {
    ...component,
    render: fun _self => {
      let renderedItems = items
        |> List.mapi handleItems
        |> Array.of_list
        |> ReasonReact.arrayToElement;
      <div style=styles##container>
        renderedItems 
      </div>
    }
  }
};