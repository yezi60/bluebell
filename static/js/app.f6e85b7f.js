(function(t){function s(s){for(var n,o,c=s[0],r=s[1],l=s[2],m=0,p=[];m<c.length;m++)o=c[m],Object.prototype.hasOwnProperty.call(a,o)&&a[o]&&p.push(a[o][0]),a[o]=0;for(n in r)Object.prototype.hasOwnProperty.call(r,n)&&(t[n]=r[n]);u&&u(s);while(p.length)p.shift()();return i.push.apply(i,l||[]),e()}function e(){for(var t,s=0;s<i.length;s++){for(var e=i[s],n=!0,c=1;c<e.length;c++){var r=e[c];0!==a[r]&&(n=!1)}n&&(i.splice(s--,1),t=o(o.s=e[0]))}return t}var n={},a={app:0},i=[];function o(s){if(n[s])return n[s].exports;var e=n[s]={i:s,l:!1,exports:{}};return t[s].call(e.exports,e,e.exports,o),e.l=!0,e.exports}o.m=t,o.c=n,o.d=function(t,s,e){o.o(t,s)||Object.defineProperty(t,s,{enumerable:!0,get:e})},o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},o.t=function(t,s){if(1&s&&(t=o(t)),8&s)return t;if(4&s&&"object"===typeof t&&t&&t.__esModule)return t;var e=Object.create(null);if(o.r(e),Object.defineProperty(e,"default",{enumerable:!0,value:t}),2&s&&"string"!=typeof t)for(var n in t)o.d(e,n,function(s){return t[s]}.bind(null,n));return e},o.n=function(t){var s=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(s,"a",s),s},o.o=function(t,s){return Object.prototype.hasOwnProperty.call(t,s)},o.p="/";var c=window["webpackJsonp"]=window["webpackJsonp"]||[],r=c.push.bind(c);c.push=s,c=c.slice();for(var l=0;l<c.length;l++)s(c[l]);var u=r;i.push([0,"chunk-vendors"]),e()})({0:function(t,s,e){t.exports=e("56d7")},"12be":function(t,s,e){"use strict";var n=e("e4c4"),a=e.n(n);a.a},"13ca":function(t,s,e){},2178:function(t,s,e){},"22ad":function(t,s,e){},2395:function(t,s,e){},3137:function(t,s,e){},"51c7":function(t,s,e){"use strict";var n=e("2178"),a=e.n(n);a.a},"56d7":function(t,s,e){"use strict";e.r(s);e("e260"),e("e6cf"),e("cca6"),e("a79d");var n=e("2b0e"),a=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{attrs:{id:"app"}},[e("div",{staticClass:"page"},[e("HeadBar"),e("router-view")],1)])},i=[],o=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("header",{staticClass:"header"},[e("span",{staticClass:"logo",on:{click:t.goIndex}},[t._v("bluebell")]),t._m(0),e("div",{staticClass:"btns"},[e("div",{directives:[{name:"show",rawName:"v-show",value:!t.isLogin,expression:"!isLogin"}]},[e("a",{staticClass:"login-btn",on:{click:t.goLogin}},[t._v("登录")]),e("a",{staticClass:"login-btn",on:{click:t.goSignUp}},[t._v("注册")])]),e("div",{directives:[{name:"show",rawName:"v-show",value:t.isLogin,expression:"isLogin"}],staticClass:"user-box"},[e("span",{staticClass:"user"},[t._v(t._s(t.currUsername))]),e("div",{staticClass:"dropdown-content"},[e("a",{on:{click:t.goLogout}},[t._v("登出")])])])])])},c=[function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"search"},[e("label",{staticClass:"s-logo"}),e("input",{staticClass:"s-input",attrs:{type:"text",placeholder:"搜索"}})])}],r={name:"HeadBar",created:function(){this.$store.commit("init")},computed:{isLogin:function(){return this.$store.getters.isLogin},currUsername:function(){return console.log(this.$store.getters.username),this.$store.getters.username}},methods:{goIndex:function(){this.$router.push({name:"Home"})},goLogin:function(){this.$router.push({name:"Login"})},goSignUp:function(){this.$router.push({name:"SignUp"})},goLogout:function(){this.$store.commit("logout")}}},l=r,u=(e("51c7"),e("2877")),m=Object(u["a"])(l,o,c,!1,null,"28d06594",null),p=m.exports,d={components:{HeadBar:p}},v=d,f=(e("7c55"),Object(u["a"])(v,a,i,!1,null,null,null)),g=f.exports,h=e("8c4f"),C=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"content"},[e("div",{staticClass:"left"},[e("div",{staticClass:"c-l-header"},[e("div",{staticClass:"new btn-iconfont",class:{active:t.timeOrder},on:{click:function(s){return t.selectOrder("time")}}},[e("i",{staticClass:"iconfont icon-polygonred"}),t._v("New ")]),e("div",{staticClass:"top btn-iconfont",class:{active:t.scoreOrder},on:{click:function(s){return t.selectOrder("score")}}},[e("i",{staticClass:"iconfont icon-top"}),t._v("Score ")]),e("button",{staticClass:"btn-publish",on:{click:t.goPublish}},[t._v("发表")])]),e("ul",{staticClass:"c-l-list"},t._l(t.postList,(function(s){return e("li",{key:s.post.id,staticClass:"c-l-item"},[e("div",{staticClass:"post"},[e("a",{staticClass:"vote"},[e("span",{staticClass:"iconfont icon-up",on:{click:function(e){return t.vote(s.post.id,"1")}}})]),e("span",{staticClass:"text"},[t._v(t._s(s.vote_num))]),e("a",{staticClass:"vote"},[e("span",{staticClass:"iconfont icon-down",on:{click:function(e){return t.vote(s.post.id,"-1")}}})])]),e("div",{staticClass:"l-container",on:{click:function(e){return t.goDetail(s.post.id)}}},[e("h4",{staticClass:"con-title"},[t._v(t._s(s.post.title))]),e("div",{staticClass:"con-memo"},[e("p",[t._v(t._s(s.post.content))])])])])})),0)]),t._m(0)])},_=[function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"right"},[e("div",{staticClass:"communities"},[e("h2",{staticClass:"r-c-title"},[t._v("今日火热频道排行榜")]),e("ul",{staticClass:"r-c-content"},[e("li",{staticClass:"r-c-item"},[e("span",{staticClass:"index"},[t._v("1")]),e("i",{staticClass:"icon"}),t._v(" b/coding ")]),e("li",{staticClass:"r-c-item"},[e("span",{staticClass:"index"},[t._v("2")]),e("i",{staticClass:"icon"}),t._v(" b/tree_hole ")]),e("li",{staticClass:"r-c-item"},[e("span",{staticClass:"index"},[t._v("3")]),e("i",{staticClass:"icon"}),t._v(" b/job ")])]),e("button",{staticClass:"view-all"},[t._v("查看所有")])]),e("div",{staticClass:"r-trending"},[e("h2",{staticClass:"r-t-title"},[t._v("持续热门频道")]),e("ul",{staticClass:"rank"},[e("li",{staticClass:"r-t-cell"},[e("div",{staticClass:"r-t-cell-info"},[e("div",{staticClass:"avatar"}),e("div",{staticClass:"info"},[e("span",{staticClass:"info-title"},[t._v("b/Book")]),e("p",{staticClass:"info-num"},[t._v("7.1k members")])])]),e("button",{staticClass:"join-btn"},[t._v("JOIN")])]),e("li",{staticClass:"r-t-cell"},[e("div",{staticClass:"r-t-cell-info"},[e("div",{staticClass:"avatar"}),e("div",{staticClass:"info"},[e("span",{staticClass:"info-title"},[t._v("b/coding")]),e("p",{staticClass:"info-num"},[t._v("3.2k members")])])]),e("button",{staticClass:"join-btn"},[t._v("JOIN")])]),e("li",{staticClass:"r-t-cell"},[e("div",{staticClass:"r-t-cell-info"},[e("div",{staticClass:"avatar"}),e("div",{staticClass:"info"},[e("span",{staticClass:"info-title"},[t._v("b/job")]),e("p",{staticClass:"info-num"},[t._v("2.5k members")])])]),e("button",{staticClass:"join-btn"},[t._v("JOIN")])])])])])}],b={name:"Home",components:{},data:function(){return{order:"time",page:1,postList:[]}},methods:{selectOrder:function(t){this.order=t,this.getPostList()},goPublish:function(){this.$router.push({name:"Publish"})},goDetail:function(t){this.$router.push({name:"Content",params:{id:t}})},getPostList:function(){var t=this;this.$axios({method:"get",url:"/posts2",params:{page:this.page,order:this.order}}).then((function(s){console.log(s.data,222),1e3==s.code?t.postList=s.data:console.log(s.msg)})).catch((function(t){console.log(t)}))},vote:function(t,s){this.$axios({method:"post",url:"/vote",data:JSON.stringify({post_id:JSON.stringify(t),direction:s})}).then((function(t){1e3==t.code?console.log("vote success"):console.log(t.msg)})).catch((function(t){console.log(t)}))}},mounted:function(){this.getPostList()},computed:{timeOrder:function(){return"time"==this.order},scoreOrder:function(){return"score"==this.order}}},w=b,y=(e("cd30"),Object(u["a"])(w,C,_,!1,null,"4cabe882",null)),x=y.exports,$=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"content"},[e("div",{staticClass:"left"},[e("div",{staticClass:"container"},[t._m(0),e("div",{staticClass:"l-container"},[e("h4",{staticClass:"con-title"},[t._v(t._s(t.object.post.title))]),e("div",{staticClass:"con-info"},[t._v(t._s(t.object.post.content))]),t._m(1)])])]),e("div",{staticClass:"right"},[e("div",{staticClass:"topic-info"},[e("h5",{staticClass:"t-header"}),e("div",{staticClass:"t-info"},[e("a",{staticClass:"avatar"}),e("span",{staticClass:"topic-name"},[t._v("分类："+t._s(t.object.community.name))])]),e("p",{staticClass:"t-desc"},[t._v("树洞 树洞 无限树洞的树洞")]),t._m(2),e("div",{staticClass:"date"},[t._v("Created Apr 10, 2008")]),e("button",{staticClass:"topic-btn"},[t._v("JOIN")])])])])},O=[function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"post"},[e("a",{staticClass:"vote"},[e("span",{staticClass:"iconfont icon-up"})]),e("span",{staticClass:"text"},[t._v("50.2k")]),e("a",{staticClass:"vote"},[e("span",{staticClass:"iconfont icon-down"})])])},function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"user-btn"},[e("span",{staticClass:"btn-item"},[e("i",{staticClass:"iconfont icon-comment"}),t._v("comment ")])])},function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("ul",{staticClass:"t-num"},[e("li",{staticClass:"t-num-item"},[e("p",{staticClass:"number"},[t._v("5.2m")]),e("span",{staticClass:"unit"},[t._v("Members")])]),e("li",{staticClass:"t-num-item"},[e("p",{staticClass:"number"},[t._v("5.2m")]),e("span",{staticClass:"unit"},[t._v("Members")])])])}],L={name:"Content",data:function(){return{object:{}}},methods:{getPostDetail:function(){var t=this;this.$axios({method:"get",url:"/post/"+this.$route.params.id}).then((function(s){console.log(1,s.data),1e3==s.code?t.object=s.data:console.log(s.msg)})).catch((function(t){console.log(t)}))}},mounted:function(){this.getPostDetail()}},k=L,P=(e("c579"),Object(u["a"])(k,$,O,!1,null,"d29c9b40",null)),j=P.exports,S=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"content"},[e("div",{staticClass:"left"},[e("div",{staticClass:"post-name"},[t._v("我好想写点什么")]),e("div",{staticClass:"post-type"},[e("input",{directives:[{name:"model",rawName:"v-model",value:t.selectCommunity.name,expression:"selectCommunity.name"}],staticClass:"post-type-value",attrs:{type:"text",placeholder:"选择一个频道"},domProps:{value:t.selectCommunity.name},on:{click:function(s){return t.showCommunity()},input:function(s){s.target.composing||t.$set(t.selectCommunity,"name",s.target.value)}}}),e("ul",{directives:[{name:"show",rawName:"v-show",value:t.showCommunityList,expression:"showCommunityList"}],staticClass:"post-type-options"},t._l(t.communityList,(function(s,n){return e("li",{key:s.id,staticClass:"post-type-cell",on:{click:function(s){return t.selected(n)}}},[t._v(" "+t._s(s.name)+" ")])})),0),e("i",{staticClass:"p-icon"})]),e("div",{staticClass:"post-content"},[t._m(0),e("div",{staticClass:"post-sub-container"},[e("div",{staticClass:"post-sub-header"},[e("textarea",{directives:[{name:"model",rawName:"v-model",value:t.title,expression:"title"}],staticClass:"post-title",attrs:{id:"",cols:"30",rows:"10",placeholder:"标题"},domProps:{value:t.title},on:{input:function(s){s.target.composing||(t.title=s.target.value)}}}),e("span",{staticClass:"textarea-num"},[t._v("0/300")])]),e("div",{staticClass:"post-text-con"},[e("textarea",{directives:[{name:"model",rawName:"v-model",value:t.content,expression:"content"}],staticClass:"post-content-t",attrs:{id:"",cols:"30",rows:"10",placeholder:"内容"},domProps:{value:t.content},on:{input:function(s){s.target.composing||(t.content=s.target.value)}}})])]),e("div",{staticClass:"post-footer"},[e("div",{staticClass:"btns"},[e("button",{staticClass:"btn"},[t._v("取消")]),e("button",{staticClass:"btn",on:{click:function(s){return t.submit()}}},[t._v("发表")])])])])]),t._m(1)])},N=[function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("ul",{staticClass:"cat"},[e("li",{staticClass:"cat-item active"},[e("i",{staticClass:"iconfont icon-edit"}),t._v("post ")]),e("li",{staticClass:"cat-item"},[e("i",{staticClass:"iconfont icon-image"}),t._v("image/video ")])])},function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"right"},[e("div",{staticClass:"post-rank"},[e("h5",{staticClass:"p-r-title"},[e("i",{staticClass:"p-r-icon"}),t._v("发帖规范 ")]),e("ul",{staticClass:"p-r-content"},[e("li",{staticClass:"p-r-item"},[t._v("1.网络不是法外之地")]),e("li",{staticClass:"p-r-item"},[t._v("2.网络不是法外之地")]),e("li",{staticClass:"p-r-item"},[t._v("3.网络不是法外之地")])])])])}],R={name:"Publish",data:function(){return{title:"",content:"",showCommunityList:!1,selectCommunity:{},communityList:[]}},methods:{submit:function(){var t=this;this.$axios({method:"post",url:"/post",data:JSON.stringify({title:this.title,content:this.content,community_id:this.selectCommunity.id})}).then((function(s){console.log(s.data),1e3==s.code?t.$router.push({path:t.redirect||"/"}):console.log(s.msg)})).catch((function(t){console.log(t)}))},getCommunityList:function(){var t=this;this.$axios({method:"get",url:"/community"}).then((function(s){console.log(s.data),1e3==s.code?t.communityList=s.data:console.log(s.msg)})).catch((function(t){console.log(t)}))},showCommunity:function(){this.showCommunityList=!this.showCommunityList},selected:function(t){this.selectCommunity=this.communityList[t],this.showCommunityList=!1,console.log(this.selectCommunity)}},mounted:function(){this.getCommunityList()}},E=R,I=(e("12be"),Object(u["a"])(E,S,N,!1,null,"6620057e",null)),J=I.exports,U=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"main"},[e("div",{staticClass:"container"},[e("h2",{staticClass:"form-title"},[t._v("登录")]),e("div",{staticClass:"form-group"},[e("label",{attrs:{for:"name"}},[t._v("用户名")]),e("input",{directives:[{name:"model",rawName:"v-model",value:t.username,expression:"username"}],staticClass:"form-control",attrs:{type:"text",name:"name",id:"name",placeholder:"用户名"},domProps:{value:t.username},on:{input:function(s){s.target.composing||(t.username=s.target.value)}}})]),e("div",{staticClass:"form-group"},[e("label",{attrs:{for:"pass"}},[t._v("密码")]),e("input",{directives:[{name:"model",rawName:"v-model",value:t.password,expression:"password"}],staticClass:"form-control",attrs:{type:"password",name:"pass",id:"pass",placeholder:"密码"},domProps:{value:t.password},on:{input:function(s){s.target.composing||(t.password=s.target.value)}}})]),e("div",{staticClass:"form-btn"},[e("button",{staticClass:"btn btn-info",attrs:{type:"button"},on:{click:t.submit}},[t._v("提交")])])])])},H=[],M={name:"Login",data:function(){return{username:"",password:"",submitted:!1}},computed:{},created:function(){},methods:{submit:function(){var t=this;this.$axios({method:"post",url:"/login",data:JSON.stringify({username:this.username,password:this.password})}).then((function(s){console.log(s.data),1e3==s.code?(localStorage.setItem("loginResult",JSON.stringify(s.data)),t.$store.commit("login",s.data),t.$router.push({path:t.redirect||"/"})):console.log(s.msg)})).catch((function(t){console.log(t)}))}}},B=M,D=(e("b103"),Object(u["a"])(B,U,H,!1,null,"18c22d5f",null)),A=D.exports,T=function(){var t=this,s=t.$createElement,e=t._self._c||s;return e("div",{staticClass:"main"},[e("div",{staticClass:"container"},[e("h2",{staticClass:"form-title"},[t._v("注册")]),e("div",{staticClass:"form-group"},[e("label",{attrs:{for:"name"}},[t._v("用户名")]),e("input",{directives:[{name:"model",rawName:"v-model",value:t.username,expression:"username"}],staticClass:"form-control",attrs:{type:"text",name:"name",id:"name",placeholder:"用户名"},domProps:{value:t.username},on:{input:function(s){s.target.composing||(t.username=s.target.value)}}})]),e("div",{staticClass:"form-group"},[e("label",{attrs:{for:"pass"}},[t._v("密码")]),e("input",{directives:[{name:"model",rawName:"v-model",value:t.password,expression:"password"}],staticClass:"form-control",attrs:{type:"password",name:"pass",id:"pass",placeholder:"密码"},domProps:{value:t.password},on:{input:function(s){s.target.composing||(t.password=s.target.value)}}})]),e("div",{staticClass:"form-group"},[e("label",{attrs:{for:"re_pass"}},[t._v("确认密码")]),e("input",{directives:[{name:"model",rawName:"v-model",value:t.re_password,expression:"re_password"}],staticClass:"form-control",attrs:{type:"password",name:"re_pass",id:"re_pass",placeholder:"确认密码"},domProps:{value:t.re_password},on:{input:function(s){s.target.composing||(t.re_password=s.target.value)}}})]),e("div",{staticClass:"form-btn"},[e("button",{staticClass:"btn btn-info",attrs:{type:"button"},on:{click:t.submit}},[t._v("提交")])])])])},q=[],z={name:"SignUp",data:function(){return{username:"",password:"",re_password:"",submitted:!1}},computed:{},created:function(){},methods:{submit:function(){var t=this;this.$axios({method:"post",url:"/signup",data:JSON.stringify({username:this.username,password:this.password,re_password:this.re_password})}).then((function(s){console.log(s.data),1e3==s.code?(console.log("signup success"),t.$router.push({name:"Login"})):console.log(s.msg)})).catch((function(t){console.log(t)}))}}},F=z,G=(e("99b3"),Object(u["a"])(F,T,q,!1,null,"43662fcc",null)),K=G.exports,Q=h["a"].prototype.push;h["a"].prototype.push=function(t){return Q.call(this,t).catch((function(t){return t}))},n["a"].use(h["a"]);var V=[{path:"/",name:"Home",component:x},{path:"/post/:id",name:"Content",component:j},{path:"/publish",name:"Publish",component:J,meta:{requireAuth:!0}},{path:"/login",name:"Login",component:A},{path:"/signup",name:"SignUp",component:K}],W=new h["a"]({mode:"history",base:"/",routes:V}),X=W,Y=e("2f62");n["a"].use(Y["a"]);var Z={token:null,user_id:null,user_name:null},tt=new Y["a"].Store({state:{isLogin:!1,loginResult:Z},mutations:{init:function(t){var s=JSON.parse(localStorage.getItem("loginResult"));console.log(localStorage.getItem("loginResult")),null!=s&&(t.loginResult=s)},login:function(t,s){t.loginResult=s},logout:function(t){localStorage.removeItem("loginResult"),t.loginResult=Z}},actions:{},getters:{isLogin:function(t){return null!==t.loginResult.user_id},userID:function(t){return t.loginResult.user_id},username:function(t){return t.loginResult.user_name},accessToken:function(t){return t.loginResult.token}}}),st=(e("d3b7"),e("bc3a")),et=e.n(st);et.a.defaults.baseURL="/api/v1/",et.a.interceptors.request.use((function(t){var s=JSON.parse(localStorage.getItem("loginResult"));if(s){var e=s.token;t.headers.Authorization="Bearer ".concat(e)}return t}),(function(t){return Promise.reject(t)})),et.a.interceptors.response.use((function(t){return 200===t.status?Promise.resolve(t.data):Promise.reject(t.data)}),(function(t){console.log("error",t)}));var nt=et.a;n["a"].prototype.$axios=nt,n["a"].config.productionTip=!1,X.beforeEach((function(t,s,e){console.log(t),console.log(s),t.meta.requireAuth?localStorage.getItem("loginResult")||"/login"===t.path?e():e({path:"/login"}):e(),"/login"==t.fullPath&&(localStorage.getItem("loginResult")?e({path:s.fullPath}):e())})),new n["a"]({router:X,store:tt,render:function(t){return t(g)}}).$mount("#app")},6070:function(t,s,e){},"7c55":function(t,s,e){"use strict";var n=e("2395"),a=e.n(n);a.a},"99b3":function(t,s,e){"use strict";var n=e("3137"),a=e.n(n);a.a},b103:function(t,s,e){"use strict";var n=e("22ad"),a=e.n(n);a.a},c579:function(t,s,e){"use strict";var n=e("13ca"),a=e.n(n);a.a},cd30:function(t,s,e){"use strict";var n=e("6070"),a=e.n(n);a.a},e4c4:function(t,s,e){}});
//# sourceMappingURL=app.f6e85b7f.js.map