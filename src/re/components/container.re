let component = ReasonReact.statelessComponent "REConatainer";

let styles = {
  "container": ReactDOMRe.Style.make maxWidth::"600px" margin::"0 auto" ()
};

let make _children => {
  ...component,
  render: fun _self => {
    let greetings = ReasonReact.stringToElement "Hello World!";
    <div style=(styles##container)> 
      <Header items=Data.headerItems />
      <Body items=Data.bodyItems />
      <Footer items=Data.footerItems />
    </div>
  }
};

let jsComponent = ReasonReact.wrapReasonForJs ::component (fun jsProps => make [||])
