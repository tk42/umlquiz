(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[405],{8312:function(a,b,c){(window.__NEXT_P=window.__NEXT_P||[]).push(["/",function(){return c(1328)}])},3465:function(b,f,a){"use strict";var c=a(6992),d=a(7273),e=a(1258);b.exports={translations:{en:c,ja:d,fr:e},defaultLang:"en",useBrowserDefault:!0}},1328:function(h,b,a){"use strict";a.r(b),a.d(b,{"default":function(){return K}});var i=a(8598),j=a(2498),c=a(1887),k=a.n(c),d=a(2691),l=a.n(d),m=a(9495),n=a(7745),o=a(1111);a(8316),a(5393);var e=a(1774),f=a.n(e),p=a(6892),q=a(2684),r=f()(function(){return a.e(682).then(a.bind(a,3682))},{loadableGenerated:{webpack:function(){return[3682]}},ssr:!1}),s=function(a){return a.map(function(a){return"string"==typeof a?a:!!a.props&&!!a.props.children&&s(a.props.children)}).filter(Boolean).join("")},t=function(a){a.inline;var d=a.children,b=void 0===d?[]:d,c=a.className;!function(a,d){if(null==a)return{};var b,c,e=function(c,f){if(null==c)return{};var a,b,d={},e=Object.keys(c);for(b=0;b<e.length;b++)a=e[b],f.indexOf(a)>=0||(d[a]=c[a]);return d}(a,d);if(Object.getOwnPropertySymbols){var f=Object.getOwnPropertySymbols(a);for(c=0;c<f.length;c++)b=f[c],!(d.indexOf(b)>=0)&&Object.prototype.propertyIsEnumerable.call(a,b)&&(e[b]=a[b])}return e}(a,["inline","children","className"]);var g=(0,q.useRef)("demo".concat(parseInt(String(1e15*Math.random()),10).toString(36))),e=s(b),f=(0,q.useRef)(null);return((0,q.useEffect)(function(){if(f.current)try{var a=p.Z.render(g.current,e,function(){return null},f.current);f.current.innerHTML=a}catch(b){f.current.innerHTML=b}},[b,e,f]),"string"==typeof e&&"string"==typeof c&&/^language-mermaid/.test(c.toLocaleLowerCase()))?(0,i.jsx)("code",{ref:f,children:(0,i.jsx)("code",{id:g.current,style:{display:"none"}})}):(0,i.jsx)("code",{className:String(c),children:b})};function u(a){var b=a.value,c=a.prefix;return(0,i.jsx)("div",{children:(0,i.jsx)("div",{"data-color-mode":"light",children:(0,i.jsx)(r,{source:"```mermaid\n"+c+"\n"+b+"\n```",height:600,components:{code:t}})})})}var v="class Interface\n<<interface>> Interface\nclass Abstract\n<<abstract>> Abstract\nclass Enum\n<<enum>> Enum",w="class Visibility\nVisibility : +int public_member\nVisibility : #int protected_member\nVisibility : -int private_member\nVisibility : ~int package_member",x="classA0 --|> classB0 : Inheritance\nclassA1 --* classB1 : Composition\nclassA2 --o classB2 : Aggregation\nclassA3 --> classB3 : Association\nclassA4 ..> classB4 : Dependency\nclassA5 ..|> classB5 : Realization",y="Alice->>+Bob: Hello Bob, how are you?\nAlice-->>+Bob: Bob, can you hear me?\nBob-)-Alice: Hi Alice, I can hear you! (async)\nBob--)-Alice: I feel great! (async)";function z(){var a=(0,j.$G)().t;return(0,i.jsxs)("div",{className:"glass",children:[(0,i.jsx)(m.Z,{variant:"h2",children:a("notation")}),(0,i.jsx)(m.Z,{variant:"h3",children:a("caption_classdiagram")}),(0,i.jsxs)(n.ZP,{container:!0,spacing:2,children:[(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(o.Z,{label:"Annotations",multiline:!0,rows:6,fullWidth:!0,defaultValue:v})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(u,{value:v,prefix:"classDiagram"})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(o.Z,{label:"Visibility",multiline:!0,rows:5,fullWidth:!0,defaultValue:w})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(u,{value:w,prefix:"classDiagram"})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(o.Z,{label:"Relationships",multiline:!0,rows:6,fullWidth:!0,defaultValue:x})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(u,{value:x,prefix:"classDiagram"})})]}),(0,i.jsx)(m.Z,{variant:"h3",children:a("caption_sequencediagram")}),(0,i.jsxs)(n.ZP,{container:!0,spacing:2,children:[(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(o.Z,{label:"Sequence",multiline:!0,rows:4,fullWidth:!0,defaultValue:y})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(u,{value:y,prefix:"sequenceDiagram"})})]})]})}var A=a(9986),B=a(2677),C="Animal <|-- Duck\nAnimal <|-- Fish\nAnimal <|-- Zebra\nAnimal : +int age\nAnimal : +String gender\nAnimal: +isMammal()\nAnimal: +mate()\nclass Duck{\n    +String beakColor\n    +swim()\n    +quack()\n}\nclass Fish{\n    -int sizeInFeet\n    -canEat()\n}\nclass Zebra{\n    +bool is_wild\n    +run()\n}",D=a(7255),E=a(6626),F=a(3017),g=a(4407),G=a.n(g),H={position:"absolute",top:"50%",left:"50%",transform:"translate(-50%, -50%)",width:"80%",bgcolor:"background.paper",border:"2px solid #000",boxShadow:24,p:4};function I(a){var g=a.caption,h=a.diagram_type,k=a.text,d=a.diagram,l=a.linkto,c=(0,j.$G)().t,e=(0,q.useState)(!1),p=e[0],r=e[1],f=(0,q.useState)(C),b=f[0],s=f[1];return(0,i.jsxs)("div",{className:"glass",children:[(0,i.jsx)(m.Z,{variant:"h2",children:g}),(0,i.jsx)("p",{children:k}),(0,i.jsxs)(n.ZP,{container:!0,spacing:2,children:[(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(o.Z,{value:b,multiline:!0,fullWidth:!0,onChange:function(a){s(a.target.value)}})}),(0,i.jsx)(n.ZP,{item:!0,xs:6,children:(0,i.jsx)(u,{value:b,prefix:h})})]}),(0,i.jsx)(D.Z,{onClick:function(){return r(!0)},startIcon:"\u2713",children:"CHECK"}),(0,i.jsx)(E.Z,{open:p,onClose:function(a){return r(!1)},children:(0,i.jsxs)(F.Z,{sx:H,children:[b===d?(0,i.jsx)(m.Z,{sx:{mt:2},children:c("correct_message")}):(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(m.Z,{variant:"h6",component:"h2",children:c("wrong_message")}),(0,i.jsx)(G(),{oldValue:b,newValue:d,splitView:!0})]}),(0,i.jsx)(D.Z,{href:l,children:c("proceed_quiz")})]})})]})}function J(){var a=(0,j.$G)().t;return(0,i.jsx)(i.Fragment,{children:(0,i.jsx)(I,{caption:"Tutorial",diagram_type:"classDiagram",text:a("tutorial"),diagram:C,linkto:"/signin"})})}var K=function(){var a=(0,j.$G)().t;return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(k(),{children:(0,i.jsx)("title",{children:"UML Quiz"})}),(0,i.jsxs)("div",{className:"bg",children:[(0,i.jsx)(m.Z,{variant:"h1",children:"UML Quiz"}),(0,i.jsxs)(B.Z,{exclusive:!0,children:[(0,i.jsx)(l(),{href:"/?lang=en",children:(0,i.jsx)(A.Z,{value:"en",children:"English"})}),(0,i.jsx)(l(),{href:"/?lang=fr",children:(0,i.jsx)(A.Z,{value:"fr",children:"Fran\xe7ais"})}),(0,i.jsx)(l(),{href:"/?lang=ja",children:(0,i.jsx)(A.Z,{value:"ja",children:"\u65E5\u672C\u8A9E"})})]}),(0,i.jsx)("div",{className:"glass",children:a("intro")}),(0,i.jsx)(z,{}),(0,i.jsx)(J,{})]})]})}},6992:function(a){"use strict";a.exports=JSON.parse('{"intro":"The Unified Modeling Language (UML) is a general-purpose, developmental, modeling language in the field of software engineering that is intended to provide a standard way to visualize the design of a system. umlquiz.com helps to improve your UML writing skill.","notation":"Notation","caption_classdiagram":"Class Diagram","caption_sequencediagram":"Sequence Diagram","tutorial":"Animal class has age (+int) and gender (+string), with operations to get isMammal (mammal or not:+), mate (mate or not:+). The child classe of Animal includes Duck, Fish and Zebra. Ducks have beakColour (beak colour:+string) as an attribute and can swim (swim:+) and quack (quack:+). Fish have an attribute sizeInFeet (length:-int) that can be referenced internally and can be checked if it canEat (edible:-). A zebra has an attribute is_wild (wild or not:+bool) and can run (run:+). Draw an appropriate class diagram.","correct_message":"Congratulations\u{1F389} There\'s no mistake.","wrong_message":"Too bad\u{1F62D} The correct answers as follows","proceed_quiz":"More quiz"}')},1258:function(a){"use strict";a.exports=JSON.parse('{"intro":"Le langage de mod\xe9lisation unifi\xe9 (UML) est un langage de mod\xe9lisation polyvalent et \xe9volutif dans le domaine de l\'ing\xe9nierie logicielle, destin\xe9 \xe0 fournir un moyen standard de visualiser la conception d\'un syst\xe8me. umlquiz.com vous aide \xe0 am\xe9liorer vos comp\xe9tences en mati\xe8re de r\xe9daction UML.","notation":"Notation","caption_classdiagram":"Diagramme de classe","caption_sequencediagram":"Diagramme de s\xe9quence","tutorial":"La classe Animal a age (l\'\xe2ge:+int) et gender (le sexe:+string), avec des op\xe9rations pour obtenir isMammal (mammif\xe8re ou pas:+), mate (mate ou pas:+). La classe enfant de Animal comprend Canard, Poisson et Z\xe8bre. Les canards ont pour attribut beakColour (couleur du bec:+cha\xeene) et swim (peuvent nager:+) et quack (faire coin-coin:+). Les poissons ont un attribut sizeInFeet (longueur:-int) qui peut \xeatre r\xe9f\xe9renc\xe9 en interne et dont on peut v\xe9rifier s\'il canEat (comestible:-). Un z\xe8bre a un attribut is_wild (sauvage ou non:+bool) et run (peut courir:+). Dessinez un diagramme de classes appropri\xe9.","correct_message":"F\xe9licitations\u{1F389} Il n\'y a pas d\'erreur.","wrong_message":"Too bad\u{1F62D} Les bonnes r\xe9ponses sont les suivantes","proceed_quiz":"Plus de quiz"}')},7273:function(a){"use strict";a.exports=JSON.parse('{"intro":"\u7D71\u4E00\u30E2\u30C7\u30EA\u30F3\u30B0\u8A00\u8A9E\uFF08UML\uFF09\u306F\u3001\u30BD\u30D5\u30C8\u30A6\u30A7\u30A2\u30A8\u30F3\u30B8\u30CB\u30A2\u30EA\u30F3\u30B0\u306E\u5206\u91CE\u3067\u3001\u30B7\u30B9\u30C6\u30E0\u306E\u8A2D\u8A08\u3092\u8996\u899A\u5316\u3059\u308B\u305F\u3081\u306E\u6A19\u6E96\u7684\u306A\u65B9\u6CD5\u3092\u63D0\u4F9B\u3059\u308B\u3053\u3068\u3092\u76EE\u7684\u3068\u3057\u305F\u6C4E\u7528\u306A\u3044\u3057\u958B\u767A\u7528\u306E\u30E2\u30C7\u30EA\u30F3\u30B0\u8A00\u8A9E\u3067\u3059\u3002umlquiz.com \u306FUML\u30C7\u30B6\u30A4\u30F3\u30B9\u30AD\u30EB\u5411\u4E0A\u306E\u304A\u624B\u4F1D\u3044\u3092\u3057\u307E\u3059\u3002","notation":"\u8868\u8A18","caption_classdiagram":"\u30AF\u30E9\u30B9\u56F3","caption_sequencediagram":"\u30B7\u30FC\u30B1\u30F3\u30B9\u56F3","tutorial":"Animal(\u52D5\u7269)\u30AF\u30E9\u30B9\u306B\u306Fage(\u5E74\u9F62:+int)\u3068gender(\u6027\u5225:+string)\u304C\u3042\u308A\uFF0CisMammal(\u54FA\u4E73\u985E\u304B\u3069\u3046\u304B:+), mate(\u3064\u304C\u3044\u304B\u3069\u3046\u304B:+)\u3092\u53D6\u5F97\u3059\u308B\u64CD\u4F5C\u304C\u3042\u308A\u307E\u3059\uFF0EAnimal\u306E\u5B50\u30AF\u30E9\u30B9\u306B\u306FDuck(\u3042\u3072\u308B)\u3068Fish(\u9B5A)\u3068Zebra(\u3057\u307E\u3046\u307E)\u304C\u3044\u307E\u3059\uFF0E\u3042\u3072\u308B\u306FbeakColor(\u30AF\u30C1\u30D0\u30B7\u306E\u8272:+string)\u3092\u5C5E\u6027\u306B\u3082\u3061\uFF0Cswim(\u6CF3\u3050:+)\u3057\u305F\u308Aquack(\u30AC\u30FC\u30AC\u30FC\u9CF4\u304F:+)\u3057\u305F\u308A\u3059\u308B\u3053\u3068\u304C\u3067\u304D\u307E\u3059\uFF0E\u9B5A\u306FsizeInFeet(\u4F53\u9577:-int)\u3068\u3044\u3046\u5916\u90E8\u304B\u3089\u53C2\u7167\u3067\u304D\u306A\u3044\u5C5E\u6027\u3092\u6301\u3061\uFF0CcanEat(\u98DF\u7528\u53EF:-)\u304B\u3069\u3046\u304B\u8ABF\u3079\u308B\u3053\u3068\u304C\u3067\u304D\u307E\u3059\uFF0E\u3057\u307E\u3046\u307E\u306Fis_wild(\u91CE\u751F\u304B\u3069\u3046\u304B:+bool)\u3092\u5C5E\u6027\u306B\u6301\u3061\uFF0Crun(\u8D70\u308B:+)\u3053\u3068\u304C\u3067\u304D\u307E\u3059\uFF0E\u9069\u5207\u306A\u30AF\u30E9\u30B9\u56F3\u3092\u63CF\u304D\u307E\u3057\u3087\u3046\uFF0E","correct_message":"\u304A\u3081\u3067\u3068\u3046\u3054\u3056\u3044\u307E\u3059\u{1F389} \u9593\u9055\u3044\u306F\u4E00\u3064\u3082\u898B\u3064\u304B\u308A\u307E\u305B\u3093\u3067\u3057\u305F","wrong_message":"\u6B8B\u5FF5\u{1F62D} \u6B63\u7B54\u4F8B\u306F\u4E0B\u8A18\u306E\u901A\u308A","proceed_quiz":"\u304A\u304B\u308F\u308A\uFF01"}')}},function(a){a.O(0,[578,672,961,774,888,179],function(){return a(a.s=8312)}),_N_E=a.O()}])