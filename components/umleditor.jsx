import "@uiw/react-md-editor/markdown-editor.css";
import "@uiw/react-markdown-preview/markdown.css";
import dynamic from "next/dynamic";
import mermaid from "mermaid";
import { useRef, useEffect } from "react";

const PREFIX = "```mermaid\n"
const SUFFIX = "\n```"

const MarkdownPreview = dynamic(
  () => import("@uiw/react-markdown-preview"),
  { ssr: false }
);

const getCode = (arr) =>
    arr.map((dt) => {
        if (typeof dt === "string") {
            return dt;
        }
        if (dt.props && dt.props.children) {
            return getCode(dt.props.children);
        }
        return false;
        })
        .filter(Boolean)
        .join("");

const randomid = () => parseInt(String(Math.random() * 1e15), 10).toString(36);

const Code = ({ inline, children = [], className, ...props }) => {
    const demoid = useRef(`demo${randomid()}`);
    const code = getCode(children);
    const demo = useRef(null);
    useEffect(() => {
      if (demo.current) {
        try {
          const str = mermaid.render(
            demoid.current,
            code,
            () => null,
            demo.current
          );
          // @ts-ignore
          demo.current.innerHTML = str;
        } catch (error) {
          // @ts-ignore
          demo.current.innerHTML = error;
        }
      }
    }, [children, code, demo]);
  
    if (
      typeof code === "string" &&
      typeof className === "string" &&
      /^language-mermaid/.test(className.toLocaleLowerCase())
    ) {
      return (
        <code ref={demo}>
          <code id={demoid.current} style={{ display: "none" }} />
        </code>
      );
    }
    return <code className={String(className)}>{children}</code>;
  };

export default function UMLEditor({value}) {
  return (
    <div>
        <div data-color-mode="light">
          <MarkdownPreview
            source={PREFIX+value+SUFFIX}
            // source={value}
            height={600}
            components={{
              code: Code
            }}
          />
        </div>
    </div>
  );
};
